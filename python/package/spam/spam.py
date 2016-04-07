# -*- coding: utf-8 -*-
from ..lib import create_greeting_func


def hello(name):
    print('hello, {0} from spam!'.format(name))


def greeting(greeting, name):
    f = create_greeting_func(greeting)
    f(name)
