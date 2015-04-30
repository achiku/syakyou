# -*- coding: utf-8 -*-


def make_named_set(name, s):
    return {'name': name, 'set': s}


if __name__ == '__main__':
    sets = [
        make_named_set('set_a', set([1, 2, 3, 4])),
        make_named_set('set_b', set([1, 2, 3])),
        make_named_set('set_c', set([1, 2, 3, 4])),
        make_named_set('set_d', set([1, 2, 3, 4, 5])),
        make_named_set('set_e', set([5, 6])),
    ]

    # pick one main set and sets excluding the main set and put them in to list of tuple
    test_data_set = [(sets[idx], [s for i, s in enumerate(sets) if i != idx]) for idx in xrange(len(sets) - 1)]
    for i in test_data_set:
        print "## {} is the main set".format(i[0]['name'])
        for j in i[1]:
            print "({}){} - ({}){} = {}".format(
                i[0]['name'], i[0]['set'], j['name'], j['set'], i[0]['set'] - j['set']
            )
