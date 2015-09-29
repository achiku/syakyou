# -*- coding: utf-8 -*-
import math
import random

if __name__ == '__main__':
    li = range(1, 11)
    rate = 0.3
    n = int(math.ceil(rate * len(li)))

    for _ in range(5):
        random.shuffle(li)
        print(li[:n])
