# -*- coding: utf-8 -*-
# http://lucumr.pocoo.org/2007/5/21/getting-started-with-wsgi/
import re
from cgi import escape


def index(environ, start_response):
    start_response('200 OK', [('Content-Type', 'text/plain')])
    return ['''
    Hello World Application
    This is the Hello World Application:
    '''.encode('utf-8')]


def hello(environ, start_response):
    args = environ['myapp.url_args']
    if args:
        subject = escape(args[0])
    else:
        subject = 'World'
    start_response('200 OK', [('Content-Type', 'text/plain')])
    return ['''
    Hello {subject}
    Hello {subject}!!
    '''.format(subject=subject).encode('utf-8')]


def not_found(environ, start_response):
    start_response('404 NOT FOUND', [('Content-Type', 'text/plain')])
    return ['NOT FOUND'.encode('utf-8')]


urls = [
    ('^$', index),
    ('^hello/?$', hello),
    ('^hello/(.+)$', hello),
]


def application(environ, start_response):
    path = environ.get('PATH_INFO', '').lstrip('/')
    for regex, callback in urls:
        match = re.search(regex, path)
        if match is not None:
            environ['myapp.url_args'] = match.groups()
            return callback(environ, start_response)
    return not_found(environ, start_response)


if __name__ == '__main__':
    from wsgiref.simple_server import make_server
    srv = make_server('localhost', 8080, application)
    srv.serve_forever()
