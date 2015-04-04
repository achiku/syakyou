# -*- coding: utf-8 -*-
# http://www-he.scphys.kyoto-u.ac.jp/member/n.kamo/wiki/doku.php?id=study:language:python:scope

if __name__ == '__main__':
    var = 'var is defined locally'
    print var

    def local_var():
        var = 'var in func is defined'
        # var is defined locally in this function
        print locals()
        return var

    print local_var()

    def global_var():
        global var
        var = 'make local var global'
        # since var is set to be global, locals() doesn't return anything
        print locals()
        return var

    print global_var()

    globals()['var'] = 'var can be changed using globals()'
    print var

    def wrapper_func():
        var = 'var is defined in wrapper func'

        # closure function
        def inner_func():
            return var

        return inner_func

    f = wrapper_func()
    print f()
