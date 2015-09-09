# -*- coding: utf-8 -*-
from multiprocessing import Process
import time


def worker(name):
    print('hello, {}!'.format(name))
    time.sleep(2)
    return


if __name__ == '__main__':
    print('start')
    p = Process(target=worker, args=('8maki',))
    p.start()
    print('done')
