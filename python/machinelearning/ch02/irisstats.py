# -*- coding: utf-8 -*-
from matplotlib import pyplot as plt
from sklearn.datasets import load_iris


def basic_stats():
    data = load_iris()
    features = data.data
    feature_names = data.feature_names
    target = data.target
    # target_names = data.target_names

    fig, axes = plt.subplots(2, 3)
    pairs = [(0, 1), (0, 2), (0, 3), (1, 2), (1, 3), (2, 3)]
    for i, (p0, p1) in enumerate(pairs):
        ax = axes.flat[i]
        for t, marker, c in zip(range(3), ">ox", "rgb"):
            ax.scatter(
                features[target == t, p0],
                features[target == t, p1],
                marker=marker,
                c=c,
                s=40)
        ax.set_xlabel(feature_names[p0])
        ax.set_ylabel(feature_names[p1])
        ax.set_xticks([])
        ax.set_yticks([])
    fig.tight_layout()
