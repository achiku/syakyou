# -*- coding: utf-8 -*-

class Person(object):
    num_instances = 0

    def __init__(self, name):
        Person.num_instances += 1
        self.name = name

    def introduce_yourself(self):
        print 'Hi. My name is {}. Nice to meet you.'.format(self.name)

    @property
    def reverse_name(self):
        return self.name[::-1]

    @classmethod
    def get_num_instances(cls):
        return cls.num_instances

    @staticmethod
    def say_hello_to(name):
        print 'Hello, {}!'.format(name)


if __name__ == '__main__':
    achiku = Person('achiku')
    achiku.introduce_yourself()
    print Person.get_num_instances()
    print Person.say_hello_to('world')
    print achiku.reverse_name

    print '=' * 10

    moqada = Person('moqada')
    moqada.introduce_yourself()
    print Person.get_num_instances()
    print Person.say_hello_to(achiku.name)
    print moqada.reverse_name
