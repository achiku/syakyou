# -*- coding: utf-8 -*-


class Person(object):

    def __init__(self, name):
        self.name = name

    def hello(self):
        return "Hey, I'm {}.".format(self.name)


if __name__ == '__main__':
    p1 = Person('moqada')
    print p1.hello()
    print dir(p1)

    m = p1.hello
    p1.hello_backup = m
    p1.hello = lambda: "Python rocks!!!!"
    print p1.hello()
    print p1.hello_backup()
    print dir(p1)
