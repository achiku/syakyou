# -*- coding: utf-8 -*-
import copy


class Person(object):

    default_name = 'world'
    default_names = ['moqada', 'ide', '8maki']

    def __init__(self, name=None):
        self.name = name

    def greeting(self):
        return "Hey, {}, {}".format(', '.join(self.default_names), ', '.join(Person.default_names))

    def hello(self):
        if self.name is None:
            return "Hello, Instance {} and Class {}!".format(self.default_name, Person.default_name)
        return "Hello, I'm {}".format(self.name)


if __name__ == '__main__':
    p1 = Person('moqada')
    print p1.hello()

    p2 = Person()
    print p2.hello()
    p2.default_name = '世界'
    print p2.hello()

    OriginalPerson = copy.deepcopy(Person)
    print Person
    print OriginalPerson
    p3 = Person()
    Person.default_name = 'мир'
    print p3.hello()

    Person.default_name = 'wereld'
    p4 = Person()
    print p4.hello()

    p5 = Person()
    print p5.hello()

    p6 = OriginalPerson()
    print p6.hello()
