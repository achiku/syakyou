# -*- coding: utf-8 -*-


def puts(arg1, arg2):
    """just print"""
    print(arg1)
    print(arg2)


if __name__ == '__main__':
    colors = ['red', 'green', 'blue']
    vals = [11, 21, 93, 112]
    for col, val in zip(colors, vals):
        print(col, val)

    args = ('spam', 'egg')
    puts(*args)

    dots = [(1, 2), (3, 4), (5, 6)]
    x, y = zip(*dots)
    print(x, y)

    # transpose matrix
    mtx = [
        (1, 2),
        (3, 4),
        (5, 6),
    ]
    print(zip(*mtx))

    # rotate matrix
    print(zip(*mtx[::-1]))

    seq = range(1, 10)
    print([iter(seq)]*2)
    print([1, 2]*2)
    print(zip(*[iter(seq)]*3))
    print(zip(seq, seq, seq))

    keys = ['spam', 'egg']
    vals = [42, 119]
    d = dict(zip(keys, vals))
    print(d)

    inv_d = dict(zip(d.values(), d.keys()))
    print(inv_d)

    my_scores = [1, 2, 3, 4, 5]
    your_scores = [2, 4, 2, 4, 0]
    winning_scores = map(
        lambda pair: max(pair),
        zip(my_scores, your_scores)
    )
    for i, score in enumerate(winning_scores):
        print("{}: {}".format(i, score))
