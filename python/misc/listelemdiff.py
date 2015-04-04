# -*- coding: utf-8 -*-


def make_named_list(name, l):
    return {'name': name, 'list': l}


if __name__ == '__main__':
    lists = [
        make_named_list('same_elems1', [1, 1, 1]),
        make_named_list('same_elems2', [1]),
        make_named_list('same_elems3', ['test', 'test', 'test']),
        make_named_list('diff_elems1', [2, 1, 1]),
        make_named_list('diff_elems2', [1, 1, 2]),
        make_named_list('diff_elems3', [1, 2, 1]),
        make_named_list('diff_elems4', ['test', 100, 2]),
    ]

    print "## using normal function"
    for l in lists:
        num_distinct_elems = len(set(l['list']))
        print "{}: {} -> NumDistinctElems: {}".format(
            l['name'], l['list'], num_distinct_elems)
