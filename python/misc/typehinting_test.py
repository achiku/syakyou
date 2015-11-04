# -*- coding: utf-8 -*-


def greeting(name: str) -> str:
    """greeting function"""
    return 'Hello ' + name


if __name__ == '__main__':
    print(greeting('moqada'))
    greeting(11)
    greeting('achiku') + 1
