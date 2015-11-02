# -*- coding: utf-8 -*-
from datetime import date
from dateutil.relativedelta import relativedelta

if __name__ == '__main__':
    d = date.today()
    for i in range(12):
        print(d + relativedelta(months=+i))
