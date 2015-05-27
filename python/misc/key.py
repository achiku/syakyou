# -*- coding: utf-8 -*-
import random


if __name__ == '__main__':
    length = 50
    lower = 'abcdefghijklmnopqnstuvwxyz'
    upper = 'ABCDEFGHIJKLMNOPQNSTUVWXYZ'
    digit = '0123456789'
    symbol = '!@#$%^&*_+'

    source = list(lower + upper + digit + symbol)
    print ''.join([source[random.randint(0, len(source)-1)] for _ in xrange(length)])
