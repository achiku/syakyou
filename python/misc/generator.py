#!/usr/bin/env python
# -*- coding: utf-8 -*-

if __name__ == '__main__':
    l = range(10)

    # generator
    for i in (i for i in l if i % 2 == 0):
        print(i)

    # list complihension
    for i in [i for i in l if i % 2 == 0]:
        print(i)

    # simple if statement
    for i in l:
        if i % 2 == 0:
            print(i)
