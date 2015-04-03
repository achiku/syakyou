# -*- coding: utf-8 -*-
from functools import reduce


def comp_elem(x, y):
    return x == y


def make_named_list(name, l):
    return {'name': name, 'list': l}


if __name__ == '__main__':
    lists = [
        make_named_list('same_elems1', [1, 1, 1]),
        make_named_list('same_elems2', [1]),
        make_named_list('diff_elems1', [2, 1, 1]),
        make_named_list('diff_elems2', [1, 1, 2]),
        make_named_list('diff_elems3', [1, 2, 1]),
    ]

    for l in lists:
        print "{}: {} -> {}".format(
            l['name'], l['list'], reduce(comp_elem, l['list']))
