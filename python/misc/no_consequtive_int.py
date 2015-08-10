# -*- coding: utf-8 -*-


def is_consequtive(l):
    """check if list has consequtive integers"""
    a = sorted([i + 1 for i in l])
    b = sorted(l)
    for x in a:
        if x in b:
            return True
    return False


if __name__ == '__main__':
    test_data = [
        ('empty list', []),
        ('only one elemten', [1]),
        ('no consequtive int', [1, 10, 14, 16]),
        ('one consequtive int', [2, 3, 14, 13]),
        ('two consequtive int', [1, 2, 3, 14, 13]),
    ]

    for t in test_data:
        print("{}: {}".format(t[0], is_consequtive(t[1])))
