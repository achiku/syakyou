import pandas as pd
import numpy as np


if __name__ == '__main__':
    # pd.Series
    s = pd.Series(np.random.randn(5), index=['a', 'b', 'c', 'd', 'e'])
    print(s)
    print(s.index)

    d = {'b': 1, 'a': 0, 'c': 2}
    print(pd.Series(d))

    s = pd.Series(5.0, index={'a', 'b', 'c'})
    print(s)
    print(s[:2])
    print(s[0])
    print(s.dtype)
