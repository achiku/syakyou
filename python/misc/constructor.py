# -*- coding: utf-8 -*-


class Person(object):
    def __init__(self, name, age):
        self.name = name
        self.age = age

    def __call__(self):
        return "inside __call__"

    def greet(self):
        return "Hello, I'm {}".format(self.name)


if __name__ == '__main__':
    p = Person('moqada', 29)
    print(p())
    print(p.greet())

    p.__call__ = lambda: "overwritten __call__"
    p.greet = lambda: "overwritten greet"
    print(p.greet())
    print(p())
    print(p.__call__())
    print(type(p).__call__(p))

    print(type(p))
    type(p).__call__ = lambda p: "new __call__"
    print(type(p).__call__(p))
    print(p.__call__())
    print(p())
