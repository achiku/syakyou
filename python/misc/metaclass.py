# -*- coding: utf-8 -*-
# thanks a lot to
# http://www.yunabe.jp/docs/python_metaclass.html
# https://fuhm.net/super-harmful/


def line():
    print "=" * 10


class Person(object):
    def __init__(self, name, age):
        self.name = name
        self.age = age

    def hello(self):
        print "Hello! I'm {}.".format(self.name)


def type_person_init(self, name, age):
    self.name = name
    self.age = age


def type_person_hello(self):
    print "Hello! I'm {}.".format(self.name)


# class is a just syntax sugar of type(classname, tpl, dict)
TypePerson = type(
    'TypePerson',
    (object,),
    {
        '__init__': type_person_init,
        'hello': type_person_hello,
    }
)


class TestNewClass1(object):
    def __new__(cls, arg):
        print '__new__ with {} and {}'.format(cls, arg)
        # return just a str, not TestNewClass instance
        # this will not trigger __init__ function and just create string representation of arg
        # This class is just for illustration, and not practical at any fucking sense
        return str(arg)

    def __init__(self, arg):
        self.var = arg
        print '__init__ with {}'.format(arg)


class TestNewClass2(object):
    def __new__(cls, arg):
        print '__new__ with {} and {}'.format(cls, arg)
        ins = super(TestNewClass2, cls).__new__(cls)
        # return TestNewClass instance, this will trigger __init__
        # this will actually generate instance of TestNewClass2 initialized by __init__
        # note, this is written in Python 2.7, so we can't eliminate arguments in super()
        return ins

    def __init__(self, arg):
        self.var = arg
        print '__init__ with {}'.format(arg)


def to_lowercase_classname(name, bases, d):
    return type(name.lower(), bases, d)


class TestMetaClass1(object):
    __metaclass__ = to_lowercase_classname

    def __init__(self, arg):
        self.val = arg
        print '__init__ with {} and {}'.format(self, arg)


if __name__ == '__main__':
    achiku = Person('achiku', 29)
    print achiku.name, achiku.age
    achiku.hello()
    print type(achiku)

    moqada = TypePerson('moqada', 29)
    print moqada.name, moqada.age
    moqada.hello()
    print type(moqada)

    line()

    t1 = TestNewClass1(43)
    print t1
    print type(t1)

    line()

    t2 = TestNewClass2(43)
    print t2
    print type(t2)
    print t2.var

    line()

    print TestMetaClass1.__name__
    m = TestMetaClass1('test')
    print m.val
