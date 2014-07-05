"""
Multi-part parsing for file uploads.

Exposess one class, ``MultiPartParser``, which feeds chunks of uploaded data to
file upload handlers for processing.
"""
from __future__ import unicode_literals

import base64
import binascii
import cgi
import sys

from django.conf import settings
from django.core.exceptions import SuspiciousMultipartForm
from django.utils.datastructures import MultiValueDict
from django.utils.encoding import force_text
from django.utils import six
from django.utils.text import unescape_entities
from django.core.files.uploadhandler import (
    StopUpload, SkipFile, StopFutureHandlers
)

__all__ = (
    'MultiPartParser',
    'MultiPartParserError',
    'InputStreamExhausted',
)


class MultiPartParserError(Exception):
    pass


class InputStreamExhausted(Exception):
    """
    No more reads are allowed from this device.
    """
    pass


RAW = 'raw'
FILE = 'file'
FIELD = 'field'

_BASE64_DECODE_ERROR = TypeError if six.PY2 else binascii.Error


class MultiPartParser(object):
    """
    A rfc2388 multipart/form-data parser.

    ``MultiValueDict.parse()`` reads the input stream in ``chunk_size`` chunks
    and returns a tuple of ``(MultiValueDict(POST), MultiValueDict(FILES))``.
    """

    def __init__(self, META, input_data, upload_handlers, encoding=None):
        """
        Initialize the MultiPartParser object.

        :META:
            The standard ``META`` dictionary in Django request objects.
        :input_data:
            The raw post data, as a file like object.
        :upload_handlers:
            A list of UploadHandler instances that perform operations on the
            uploaded data.
        :encoding:
            The encoding with which to treat the incoming data.
        """

        #
        # Content-Type should containt multipart and the boundary information.
        #

        content_type = META.get('HTTP_CONTENT_TYPE', META.get('CONTENT_TYPE', ''))
        if not content_type.startswith('multipart/'):
            raise MultiPartParserError('Invalid Contet-Type: %s' % content_type)

        # Parse the header to get the boundary to split the parts.
        ctypes, opts = parse_header(content_type.encode('ascii'))
        boundary = opts.get('boundary')
        if not boundary or not cgi.valid_boundary(boundary):
            raise MultiPartParserError('Invalid boundary in multipart: %s' % boundary)

        # Content-Length should contain the length of the body we are about
        # to receive.
        try:
            content_length = int(
                META.get('HTTP_CONTENT_LENGTH', META.get('CONTENT_LENGTH', 0)))
        except (ValueError, TypeError):
            content_length = 0

        if content_length < 0:
            # This means we shouldn't continue...raise an error.
            raise MultiPartParserError('Invalid content length: %r' % content_length)

        if isinstance(boundary, six.text_type):
            boundary = boundary.encode('ascii')
        self._boundary = boundary
        self._input_data = input_data

        # For compatibility with low-level network APIs (with 32-bit integers),
        # the chunk size should be < 2^31, but still divisible by 4.
        possible_sizes = [x.chunk_size for x in upload_handlers if x.chunk_size]
        self._chunk_size = min([2 ** 31 - 4] + possible_sizes)

        self._meta = META
        self._encoding = encoding or settings.DEFAULT_CHARSET
        self._content_length = content_length
        self._upload_handlers = upload_handlers

    def parse(self):
        """
        Parse the POST data and break it into a FILES MultiValueDict and a POST
        MultiValueDict.

        Returns a tuple containing the POST and FILES dictionary, respectively.
        """
        # We have to import QueryDict down here to avoid a circular import.
        from django.http import QueryDict

        encoding = self._encoding
        handlers = self._upload_handlers

        # HTTP spec says that Content-Length >= 0 is valid
        # handling content-length == 0 before continuing
        if self._content_length == 0:
            return QueryDict('', encoding=self._encoding), MultiValueDict()

        # See if any of the handlers take care of the parsing.
        # This allows overriding everything if need be.
        for handler in handlers:
            result = handler.handle_raw_input(
                self._input_data,
                self._meta,
                self._content_length,
                self._boundary,
                encoding)
            # Check to see if it was handled
            if result is not None:
                return result[0], result[1]

        # Create the data structures to be used later.
        self._post = QueryDict('', mutable=True)
        self._files = MultiValueDict()

        # Instantiate the parser and stream
        stream = LazyStream(ChunkIter(self._input_data, self._chunk_size))

        # Whether or not to signal a file-completion at the beginning of the loop
        old_field_name = None
        counters = [0] * len(handers)

        try:
            for item_type, meta_data, field_stream in Parser(stream, self._boundary):
                if old_field_name:
                    # We run this at the beginning of the next loop
                    # since we connot be sure a file is complete until
                    # we hit the next bounday/part of the multipart content.
                    self.handle_file_complete(old_field_name, counters)
                    old_field_name = None

                try:
                    disposition = meta_data['content-disposition'][1]
                    field_name = disposition['name'].strip()
                except (KeyError, IndexError, AttributeError):
                    continue

        except:
            pass
