# -*- coding: utf-8 -*-
import numpy as np
from sklearn.datasets import load_boston
from sklearn.linear_model import LinearRegression
from sklearn.linear_model import ElasticNet
from sklearn.cross_validation import KFold
from matplotlib import pyplot as plt


def base_stats():
    boston = load_boston()
    # print boston.feature_names
    # print boston.DESCR

    x = boston.data
    y = boston.target
    lr = LinearRegression()
    lr.fit(x, y)
    rmse = np.sqrt(lr.residues_/len(x))
    print 'RMSE: {}'.format(rmse)

    # plt.subplot(行数, 列数, 何番目のプロットか)
    plt.subplot(2, 1, 1)
    plt.scatter(lr.predict(x), boston.target)
    plt.plot([0, 50], [0, 50], '-', color=(.9, .3, .3), lw=4)
    plt.xlabel('predicted')
    plt.ylabel('real')

    x = np.array([np.concatenate((v, [1])) for v in boston.data])
    y = boston.target

    s, total_error, _, _ = np.linalg.lstsq(x, y)
    rmse = np.sqrt(total_error[0] / len(x))
    print 'Residual: {}'.format(rmse)

    plt.subplot(2, 1, 2)
    plt.plot(np.dot(x, s), boston.target, 'ro')
    plt.plot([0, 50], [0, 50], 'g-')
    plt.ylabel('real')


def num_rooms():
    boston = load_boston()
    plt.xlabel('# of rooms')
    plt.ylabel('Rent')
    plt.scatter(boston.data[:, 5], boston.target, color='b')

    x = boston.data[:, 5]  # get room size
    x = np.array([[v] for v in x])
    y = boston.target
    slope, _, _, _ = np.linalg.lstsq(x, y)

    def f(x):
        return slope * x

    fx = np.linspace(0, 10)
    plt.plot(fx, f(fx), linewidth=4, color='r')


def overview():
    boston = load_boston()
    features = [
        [0, 'CRIM', "per capita crime rate by town"],
        [1, 'ZN', "proportion of residential land zoned for lots over 25,000 sq.ft."],
        [2, 'INDUS', "proportion of non-retail business acres per town"],
        [3, 'CHAS', "Charles River dummy variable (= 1 if tract bounds river; 0 otherwise)"],
        [4, 'NOX', "nitric oxides concentration (parts per 10 million)"],
        [5, 'RM', "average number of rooms per dwelling"],
        [6, 'AGE', "proportion of owner-occupied units built prior to 1940"],
        [6, 'DIS', "weighted distances to five Boston employment centres"],
        [7, 'RAD', "index of accessibility to radial highways"],
        [8, 'TAX', "full-value property-tax rate per $10,000"],
        [9, 'PTRATIO', "pupil-teacher ratio by town"],
        [10, 'B', "1000(Bk - 0.63)^2 where Bk is the proportion of blacks by town"],
        [11, 'LSTAT', "% lower status of the population"],
        [12, 'MEDV', "Median value of owner-occupied homes in $1000's"],
    ]

    plot_row = 4
    plot_col = 4
    plt.figure(figsize=(10, 10))
    for f in features:
        print '{}:\t{}'.format(f[1], f[2])

    for feature in features:
        # plt.subplot(行数, 列数, 何番目のプロットか)
        plt.subplot(plot_row, plot_col, feature[0] + 1)
        plt.scatter(boston.data[:, feature[0]], boston.target)
        plt.xlabel(feature[1])
    plt.tight_layout()


def sklean_linear_model():
    lr = LinearRegression(fit_intercept=True)
    boston = load_boston()
    x = boston.data
    y = boston.target

    lr.fit(x, y)
    p = map(lr.predict, x)
    e = p - y

    total_error = np.sum(e * e)
    rmse_train = np.sqrt(total_error / len(p))
    print "RMSE on training: {}".format(rmse_train)


def sklean_linear_model_cross_validation():
    lr = LinearRegression(fit_intercept=True)
    boston = load_boston()
    x = boston.data
    y = boston.target

    kf = KFold(len(x), n_folds=10)
    err = 0
    for train, test in kf:
        lr.fit(x[train], y[train])
        p = map(lr.predict, x[test])
        e = p - y[test]
        err += np.sum(e * e)
    rmse_10cv = np.sqrt(err / len(x))
    print "RMSE on 10-fold CV: {}".format(rmse_10cv)


def sklean_linear_model_elastic_net():
    en = ElasticNet(fit_intercept=True, alpha=0.5)
    boston = load_boston()
    x = boston.data
    y = boston.target

    kf = KFold(len(x), n_folds=10)
    err = 0
    for train, test in kf:
        en.fit(x[train], y[train])
        p = map(en.predict, x[test])
        e = p - y[test]
        err += np.sum(e * e)
    rmse_10cv = np.sqrt(err / len(x))
    print "RMSE on 10-fold CV: {}".format(rmse_10cv)
