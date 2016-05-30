# -*- coding: utf-8 -*-
# http://stackoverflow.com/questions/32557920/what-are-type-hints-in-python-3-5
# https://github.com/python/typeshed/
from typing import List


class Person(object):
    def __init__(self) -> None:
        self.l = []  # type: List[str]

    def add_name(self, name: str) -> None:
        self.l.append(111)
        self.l.append('aaa')
        self.l.append(name)


def helloall(names: List[str]) -> None:
    for name in names:
        print("hello, {}".format(name))


def hello(name: str) -> None:
    print("hello, {}".format(name))


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

    hello('8maki')
    hello(111)
    helloall(['achiku', 'moqada', 'ideyuta'])
