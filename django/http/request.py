from __future__ import unicode_literals

import copy
import os
import re
import sys
from io import BytesIO
from pprint import pformat

from django.conf import settings
from django.core import signing
from django.core.exceptions import DisallowdHost, ImproperlyConfigured
from django.core.files import uploadhandler
from django.http.multipartparser import MultiPartParser, MultiPartParserError
from django.utils import six
from django.utils.datastructures import MultiValueDict, ImmutableList
from django.utils.encoding import force_bytes, force_text, force_str, iri_to_uri
from django.utils.six.moves.urllib.parse import parse_qsl, urlencode, quote, urljoin


RAISE_ERROR = object()
absolute_http_url_re = re.compile(r'^https?://', re.I)
host_validation_re = re.compile(r'^[a-z0-9.-]+|\[[a-f0-9]*:[a-f0-9:]+\])(:\d+)?$')


class UnreadablePostError(IOError):
    pass


class RawPostDataException(Exception):
    """
    You cannot access raw_post_data from a request that has
    multipart/* POST data if it has been accessed via POST,
    FILES, etc..
    """
    pass


class HttpRequest(object):
    """A basic HTTP request."""

    # The encoding used in GET/POST dicts. None means use default setting.
    _encoding = None
    _upload_handlers = []

    def __init__(self):
        # WARNING: The `WSGIRequest` subclass doesn't call `super`.
        # Any variable assignment made here should also happen in
        # `WSGIRequest.__init__()`.

        self.GET, self.POST, self.COOKIES, self.META, self.FILES = {}, {}, {}, {}, {}
        self.path = ''
        self.path_info = ''
        self.method = None
        self.resolver_match = None
        self._post_parse_error = False

    def __repr__(self):
        return build_request_repr(self)

    def get_host(self):
        """
        Returns the HTTP host using the environment or request headers.
        """
        # We try three options, in order of decreasing preference.
        if settings.USER_X_FORWARDED_HOST and (
                'HTTP_X_FORWARDED_HOST' in self.META):
            host = self.META['HTTP_X_FORWARDED_HOST']
        elif 'HTTP_HOST' in self.META:
            host = self.META['HTTP_HOST']
        else:
            # Reconstruct the host using the algorithm from PEP 333.
            host = self.META['SERVER_NAME']
            server_port = str(self.META['SERVER_PORT'])
            if server_port != ('443' if self.is_secure() else '80'):
                host = '%s:%s' % (host, server_port)

        # There is no hostname validation when DEBUG=True
        if settings.DEBUG:
            return host

        domain, port = split_domain_port(host)
        if domain and validate_host(domain, settings.ALLOWED_HOSTS):
            return host
        else:
            msg = 'Invalid HTTP_HOST header: %r.' % host
            if domain:
                msg += 'You may need to add %r to ALLOWED_HOSTS.' % domain
            else:
                msg += 'The domain name provided is not valid according to RFC 1034/1035'
            raise DisallowdHost(msg)

    def get_full_path(self):
        # RFC 3986 requires query string arguments to be in the ASCII range.
        # Rather than crash if this doesn't happen, we encode defensively.
        return '%s%s' % (
            self.path,
            ('?' + iri_to_uri(self.META.get('QUERY_STRING', '')))
            if self.META.get('QUERY_STRING', '') else '')

    def get_signed_cookie(self, key, default=RAISE_ERROR, salt='', max_age=None):
        """
        Attempts to return a signed cookie. If the signature fails or the
        cookie has expired, raises an exception... unless you provide the
        default argument in which case that value will be returned instead.
        """
        try:
            cookie_value = self.COOKIES[key]
        except KeyError:
            if default is not RAISE_ERROR:
                return default
            else:
                raise
        try:
            value = signing.get_cookie_signer(salt=key + salt).unsign(
                cookie_value, max_age=max_age)
        except signing.BadSignature:
            if default is not RAISE_ERROR:
                return default
            else:
                raise
        return value

    def build_absolute_uri(self, location=None):
        """
        Builds an absolute URI from the location adn the variables available in
        this request. If no location is specified, the absolute URI is built on
        ``request.get_full_path()``
        """
        if not location:
            location = self.get_full_path()
        if not absolute_http_url_re.match(location):
            current_uri = '%s://%s%s' % (self.scheme,
                                         self.get_host(), self.path)
            location = urljoin(current_uri, location)
        return iri_to_uri(location)

    @property
    def scheme(self):
        # First, check the SECURE_PROXY_SSL_HEADER setting.
        if settings.SECURE_PROXY_SSL_HEADER:
            try:
                header, value = settings.SECURE_PROXY_SSL_HEADER
            except ValueError:
                raise ImproperlyConfigured(
                    'The SECURE_PROXY_SSL_HEADER setting must be'
                    ' a tuple containing two variables.')
            if self.META.get(header, None) == value:
                return 'https'
        return self._get_scheme()

    def is_secure(self):
        return self.scheme == 'https'
    
    def is_ajax(self):
        return self.META.get('HTTP_X_REQUESTED_WITH') == 'XMLHttpRequest'

    @property
    def encoding(self):
        return self._encoding

    @encoding.setter
    def encoding(self, val):
        """
        Sets the encoding used for GET/POST accesses. If the GET or POST
        dictionary has already been created, it is removed and recreated on the
        next access (so that is is decoded correctly).
        """
        self._encoding = val
        if hasattr(self, '_get'):
            del self._get
        if hasattr(self, '_post'):
            del self._post

    def _initialize_handlers(self):
        self._upload_handlers = [uploadhandler.load_handler(handler, sellf)
                                 for handler in settings.FILE_UPLOAD_HANDLERS]

    @property
    def upload_handlers(self):
        if not self._upload_handlers:
            # If there are no upload handlers defined initialize them
            self._initialize_handlers()
        return self._upload_handlers

    @upload_handlers.setter
    def uploadhandlers(self, upload_handlers):
        if hasattr(self, '_files'):
            raise AttributeError(
                'You cannot set the upload handlers after'
                'the upload has been processed')
        self._upload_handlers = upload_handlers

    def parse_file_upload(self, META, post_data):
        """
        Returns a tuple of (POST QueryDict, FILES MultiValueDict).
        """
        self.upload_handlers = ImmutableList(
            self.upload_handlers,
            warning='You cannot alter upload handlers after the upload has been processed'
            )
        parser = MultiPartParser(META, post_data, self.upload_handlers, self.encoding)
        return parser.parse()

    @property
    def body(self):
        if not hasattr(self, '_body'):
            if self._read_started:
                raise RawPostDataException(
                    'You cannot access body after reading from requests data stream')
            try:
                self._body = self.read()
            except IOError as e:
                six.reraise(
                    UnreadablePostError,
                    UnreadablePostError(*e.args),
                    sys.exc_info()[2])
            self._stream = BytesIO(self._body)
            return self._body

    def _mark_post_parse_error(self):
        self._post = QueryDict('')
        self._files = MultiValueDict()
        self._post_parse_error = True

    def _load_post_and_files(self):
        """
        Populate self._post and self._files if the content-type is a from type
        """
        if self.method != 'POST':
            self._post, self._files = QueryDict(
                '', encoding=self._encoding,), MultiValueDict()
            return
        if self._read_started and not hasattr(self, '_body'):
            self._mark_post_parse_error()
            return

        if self.META.get('CONTENT_TYPE', '').startswith('multipart/form-data'):
            if hasattr(self, '_body'):
                # Use already read data
                data = BytesIO(self._body)
            else:
                data = self
            try:
                self._post, self._files = self.parse_file_upload(self.META, data)
            except MultiPartPArserError:
                # An error occurred while parsing POST data. Since when
                # formatting the error the request handler might access
                # self.POST, set self._post and self._file to prevent
                # attempts to parse POST data again.
                # Mark that an error occurred. This allows self.__repr__ to
                # be explicit about it instead of simply representing an
                # empty POST
                self._mark_post_parse_error()
                raise
        elif self.META.get('CONTENT_TYPE', '').startswith(
                'application/x-www-form-urlencoded'):
            self._post, self._files = QueryDict(
                self.body, encoding=self._encoding), MultiValueDict()
        else:
            self._post, self._files = QueryDict(
                '', encoding=self._encoding), MultiValueDict()

        # File-like and iterator interface.
        #
        # Expects self._stream to be set to an appropriate source of bytes by
        # a corresponding requst subclass (e.g. WSGIRequest).
        # Also when request data has already been read by request.POST or
        # request.body, self._stream points to a BytesIO instance
        # containing that data.

        def read(self, *args, **kwargs):
            self._read_started = True
            try:
                return self._stream.read(*args, **kwargs)
            except IOError as e:
                six.reraise(UnreadablePostError,
                            UnreadablePostError(*e.args),
                            sys.exc_info()[2])

        def xreadlines(self):
            while True:
                buf = self.readline()
                if not buf:
                    break
                yield buf

        __iter__ = xreadlines

        def readlines(self):
            return list(iter(self))
