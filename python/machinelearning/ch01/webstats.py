# -*- coding: utf-8 -*-
import os
import scipy as sp
import matplotlib.pyplot as plt
from IPython.display import display


DATA_DIR = os.path.join(
    os.path.dirname(os.path.realpath(__file__)), 'data')


class ListTable(list):
    def _repr_html_(self):
        html = ["<table>"]
        for row in self:
            html.append("<tr>")
            for col in row:
                html.append("<td>{0}</td>".format(col))
            html.append("</tr>")
        html.append("</table>")
        return ''.join(html)


def get_cleaned_data():
    data = sp.genfromtxt(os.path.join(DATA_DIR, 'web_traffic.tsv'), delimiter='\t')
    x = data[:, 0]
    y = data[:, 1]
    print "Number of invalid entries: {}".format(sp.sum(sp.isnan(y)))
    print "Removing invalid entries."

    x = x[~sp.isnan(y)]
    y = y[~sp.isnan(y)]
    print "Number of invalid entries: {}".format(sp.sum(sp.isnan(y)))
    return x, y


def display_raw_data(x, y):
    plt.figure(num=None, figsize=(8, 6))
    plt.clf()
    plt.scatter(x, y, s=10)
    plt.title('Web traffic over the last month')
    plt.xlabel('Time')
    plt.ylabel('Hits/hour')
    plt.xticks(
        [w * 7 * 24 for w in range(10)],
        ['week {}'.format(w) for w in range(10)]
    )
    plt.autoscale(tight=True)
    plt.grid()


def basic_stats():
    x, y = get_cleaned_data()
    display_raw_data(x, y)


def error(f, x, y):
    return sp.sum((f(x) - y) ** 2)


def display_nd_models():
    x, y = get_cleaned_data()
    display_raw_data(x, y)

    # n次元モデルデータを元データプロットに合わせて描画
    legends = []
    errors = []
    for d in [1, 2, 3, 10, 100]:
        fp, residuals, rank, sv, rcond = sp.polyfit(x, y, d, full=True)
        f = sp.poly1d(fp)
        errors.append("Error (d={}): {}".format(d, error(f, x, y)))
        fx = sp.linspace(0, x[-1], 1000)
        plt.plot(fx, f(fx), linewidth=4)
        legends.append("d={}".format(d))

    plt.legend(legends, loc="upper left")
    display(errors)


def get_inflection_data(x, y):
    inflection = 3.5 * 7 * 24
    xa = x[:inflection]
    ya = y[:inflection]
    xb = x[inflection:]
    yb = y[inflection:]
    return xa, ya, xb, yb


def display_model_with_inflection_point():
    x, y = get_cleaned_data()
    display_raw_data(x, y)
    xa, ya, xb, yb = get_inflection_data(x, y)

    # inflection pointを加味したモデルをプロット
    errors = []
    for x, y in [(xa, ya), (xb, yb)]:
        f = sp.poly1d(sp.polyfit(x, y, 1))
        fx = sp.linspace(0, x[-1], 1000)
        errors.append(error(f, x, y))
        plt.plot(fx, f(fx), linewidth=4)
        plt.ylim(ymin=0)
        plt.xlim(xmin=0)

    print "Error: {}".format(sum([e for e in errors]))


def display_model_with_first_half():
    print "Train n models with first half of data set"
    x, y = get_cleaned_data()
    display_raw_data(x, y)
    xa, ya, xb, yb = get_inflection_data(x, y)

    legends = []
    errors = []
    for d in [1, 2, 3, 10, 100]:
        f = sp.poly1d(sp.polyfit(xa, ya, d))
        fx = sp.linspace(0 * 7 * 24, 6 * 7 * 24, 100)
        plt.plot(fx, f(fx), linewidth=2)
        plt.ylim(ymin=0)
        plt.xlim(xmin=0)
        errors.append(error(f, xa, ya))
        legends.append("d={}".format(d))

    plt.legend(legends, loc="upper left")
    display(errors)


def display_model_with_second_half():
    print "Train n models with second half of data set"
    x, y = get_cleaned_data()
    display_raw_data(x, y)
    xa, ya, xb, yb = get_inflection_data(x, y)

    legends = []
    errors = []
    for d in [1, 2, 3, 10, 100]:
        f = sp.poly1d(sp.polyfit(xb, yb, d))
        fx = sp.linspace(0 * 7 * 24, 6 * 7 * 24, 100)
        plt.plot(fx, f(fx), linewidth=2)
        plt.ylim(ymin=0)
        plt.xlim(xmin=0)
        errors.append(error(f, xa, ya))
        legends.append("d={}".format(d))

    plt.legend(legends, loc="upper left")
    display(errors)


def display_cross_validateion_model():
    x, y = get_cleaned_data()
    display_raw_data(x, y)
    xa, ya, xb, yb = get_inflection_data(x, y)

    frac = 0.3
    split_idx = int(frac * len(xb))
    shuffled = sp.random.permutation(list(range(len(xb))))
    print shuffled[:5]

    test_dataset = sorted(shuffled[:split_idx])
    training_dataset = sorted(shuffled[split_idx:])

    funcs = []
    for d in [1, 2, 3, 10, 100]:
        name = 'fn{}'.format(d)
        f = sp.poly1d(sp.polyfit(xb[training_dataset], yb[training_dataset], d))
        funcs.append((name, f))

    for name, f in funcs:
        print "Error d=%i: %f" % (f.order, error(f, xb[test_dataset], yb[test_dataset]))
