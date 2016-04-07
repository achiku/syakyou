# -*- coding: utf-8 -*-
from .lib import hello as lib_hello
from egg import egg
from spam import spam

if __name__ == '__main__':
    print(__package__)
    print(__name__)
    egg.hello('achiku')
    spam.hello('achiku')
    spam.greeting('こんにちは', 'achiku')
    lib_hello('achiku')
