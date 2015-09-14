# -*- coding: utf-8 -*-
# http://lucumr.pocoo.org/2007/5/21/getting-started-with-wsgi/
from cgi import parse_qs, escape


def hello_world(environ, start_response):
    parameters = parse_qs(environ.get('QUERY_STRING', ''))
    if 'subject' in parameters:
        subject = escape(parameters['subject'][0])
    else:
        subject = 'world'
    start_response('200 OK', [('Content-Type', 'text/html')])
    return ['Hello {subject} Helllo {subject}!'.format(subject=subject).encode('utf-8')]


if __name__ == '__main__':
    from wsgiref.simple_server import make_server
    srv = make_server('localhost', 8080, hello_world)
    srv.serve_forever()
