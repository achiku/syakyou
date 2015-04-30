# -*- coding: utf-8 -*-


def load():
    '''Load ML-100k data

    Returns the review matrix as a numpy array'''
    import numpy as np
    from scipy import sparse

    # The input is in the form of a CSC sparse matrix, so it's a natural fit to
    # load the data, but we then convert to a more traditional array before
    # returning
    data = np.loadtxt('sampledata/ml-100k/u.data')
    ij = data[:, :2]
    ij -= 1  # original data is in 1-based system
    values = data[:, 2]
    reviews = sparse.csc_matrix((values, ij.T)).astype(float)
    return reviews.toarray()
