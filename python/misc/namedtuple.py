# -*- coding: utf-8 -*-
from typing import NamedTuple, List

if __name__ == '__main__':
    Item = NamedTuple(
        'Item',
        [
            ('id', int),
            ('name', str),
            ('amount', int),
        ]
    )
    User = NamedTuple(
        'User',
        [
            ('id', int),
            ('name', str),
            ('age', int),
            ('items', List[Item]),
            ('money', int),
        ]
    )

    users = [
        User(id=0, name='moqada', age=30, items=[Item(id=1, name='onigiri', amount=2)], money=100),
        User(id=1, name='8maki', age=30, items=[Item(id=1, name='onigiri', amount=2)], money=210),
    ]

    for u in users:
        print(u.id, u.name, u.age)

    u = User(id=3, name='achiku', age=31, items=[Item(id=1, name='onigiri', amount=2)], money=100)
    print(u.id, u.name, u.age)
