# -*- coding: utf-8 -*-
from functools import partial, wraps


def deco(f):
    @wraps(f)
    def wrapper(*args, **kwargs):
        print 'Calling decorated function'
        return f(*args, **kwargs)
    return wrapper


@deco
def add(x, n, y):
    ''' add two number '''
    return x + n + y


if __name__ == '__main__':
    for i in xrange(1, 10):
        add_ten_and = partial(add, 10)
        print add_ten_and.func.__doc__
        print add_ten_and.args, add_ten_and.keywords
        print 'Parameter->{}, Result->{}'.format(i, add_ten_and(i, 10))
