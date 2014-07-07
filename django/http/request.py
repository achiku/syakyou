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
