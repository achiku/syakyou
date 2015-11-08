# -*- coding: utf-8 -*-
from typing import NamedTuple

if __name__ == '__main__':
    User = NamedTuple(
        'User',
        [
            ('id', int),
            ('name', str),
            ('age', int),
        ]
    )

    users = [
        User(id=0, name='moqada', age=30),
        User(id=1, name='8maki', age=30),
        User(id=2, name='ide', age=27),
    ]

    for u in users:
        print(u.id, u.name, u.age)

    u = User(id=3, name='achiku', age=31)
    print(u.id, u.name, u.age)
