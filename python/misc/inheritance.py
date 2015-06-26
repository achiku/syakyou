# -*- coding: utf-8 -*-


class Parent(object):

    def implicit(self):
        print 'Parent implicit()'

    def override(self):
        print 'Parent override()'

    def altered(self):
        print 'Parent altered()'


class Child(Parent):

    def override(self):
        print 'Child override()'

    def altered(self):
        print 'Child, before Parent altered()'
        super(Child, self).altered()
        print 'Child, after Parent altered()'


if __name__ == '__main__':
    dad = Parent()
    son = Child()

    dad.implicit()
    son.implicit()

    dad.override()
    son.override()

    dad.altered()
    son.altered()
