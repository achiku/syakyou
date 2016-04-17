# -*- coding: utf-8 -*-
# http://stackoverflow.com/questions/32557920/what-are-type-hints-in-python-3-5
# https://github.com/python/typeshed/


def greeting(name: str) -> str:
    """greeting function"""
    return 'Hello ' + name


def add(x: int, y: int) -> int:
    """add two numbers"""
    return x + y


if __name__ == '__main__':
    print(greeting('moqada'))
    greeting(11)
    greeting('achiku') + 1

    add('a', 'b')
    add(1, 3)
