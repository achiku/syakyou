# -*- coding: utf-8 -*-
import itertools as itr


if __name__ == '__main__':
    xs = range(1, 5)
    ys = range(3, 6)

    # Return r length subsequences of elements from the input iterable.
    print 'itr.combinations(xs, 3)'
    for i in itr.combinations(xs, 3):
        print(i)

    # Return r length subsequences of elements from the input iterable
    # allowing individual elements to be repeated more than once.
    print 'itr.combinations_with_replacement(xs, 3)'
    for i in itr.combinations_with_replacement(xs, 3):
        print(i)

    # Cartesian product of input iterables.
    print 'itr.product(xs, ys)'
    for i in itr.product(xs, ys):
        print(i)

    # Return successive r length permutations of elements in the iterable.
    print 'itr.permutations(ys, 3)'
    for i in itr.permutations(ys, 3):
        print(i)
