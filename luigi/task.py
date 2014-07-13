# Copyright (c) 2012 Spotify AB
#
# Licensed under the Apache License, Version 2.0 (the "License"); you may not
# use this file except in compliance with the License. You may obtain a copy of
# the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
# WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
# License for the specific language governing permissions and limitations under
# the License.

import abc
import logging
import parameter
import warnings
import traceback


Parameter = parameter.Parameter
logger = logging.getLogger('luigi-interface')


def namespace(namespace=None):
    """ Call to set namespace of tasks declared after the call.

    If called without arguments or with ``None`` as the namespace, the namespace
    is reset, which is recommended to do at the end of any file where the
    namespace is set to avlid unintentionally setting namespace on tasks outside
    of the scope of the current file
    """
    Register._default_namespace = namespace


def id_to_name_and_params(task_id):
    """ Turn a task_id into a (task_family, {params}) tuple.

    e.g. calling with ``Foo(bar=bar, baz=baz)`` returns
    ``('Foo', {'bar': 'bar', 'baz': 'baz'})``
    """
    lparen = task_id.index('(')
    task_family = task_id[:lparen]
    params = task_id[lparen + 1:-1]

    def split_equal(x):
        equals = x.index('=')
        return x[:equals], x[equals + 1:]
    if params:
        # TODO: param values with ',' in them will break this
        param_list = map(split_equal, params.split(','))
    else:
        param_list = []
    return task_family, dict(param_list)


class Register(abc.ABCMeta):
    """
    The Metaclass of :py:class:`Task`. Acts as a global registry of Tasks with
    the following properties:

    1. Cache instances of objects so that e.g. ``X(1, 2, 3)`` always returns the
       same object.
    2. Keep track of all subclasses of :py:class:`Task` adn expose them.
    """
    __instance_cache = {}
    _default_namespace = None
    _reg = []
    AMBIGUOUS_CLASS = object()  # Placeholder denoting an error
    """ If this value is returned by :py:meth:`get_reg` then there is an
    ambiguous task name (two :py:class:`Task` have the same name). This denotes
    an error.
    """

    def __new__(metacls, classname, bases, classdict):
        """ Custom class creation for namespacing. Also register all subclasses

        Set the task namespace to whatever the currently declared namespace is.
        """
        if 'task_namespace' not in classdict:
            classdict['task_namespace'] = metacls._default_namespace
        cls = super(Register, metacls).__new__(metacls, classname, bases, classdict)
        metacls._reg.append(cls)

        return cls

    def __call__(cls, *args, **kwargs):
        """ Custom class instantiation utilizing instance cache.

        If a Task has already been instantiated with the same parameters,
        the previous instance is returned to reduce number of object instances.
        """
        def instantiate():
            return super(Register, cls).__call__(*args, **kwargs)

        h = Register.__instance_cache
        if h is None:
            return instantiate()

        params = cls.get_params()
        param_values = cls.get_param_values(params, args, kwargs)
        k = (cls, tuple(param_values))

        try:
            hash(k)
        except TypeError:
            logger.debug(
                "Not all parameter values are hashable so instance isn't coming from the cache")
            return instantiate()  # unhashable types in parameters

        if k not in h:
            h[k] = instantiate()
        return h[k]

    @classmethod
    def clear_instance_cashe(self):
        """Clear/Reset the instance cache."""
        Register.__instance_cache = {}

    @classmethod
    def disable_instance_cache(self):
        """Disables the instance cashe."""
        Register.__instance_cache = None

    @property
    def task_family(cls):
        """The task family for the given class.

        If ``cls.task_namespace is None`` then it's the name of the class.
        Otherwise, ``<task_namespace>.`` is prefixed to the class name.
        """
        if cls.task_namespace is None:
            return cls.__name__
        else:
            return '{0}.{1}'.format(cls.task_namespace, cls.__name__)

    @classmethod
    def get_reg(cls):
        """Return all of the registery classes.

        :return: a ``dict`` of task_family -> class
        """
        # We have to do this on-demand in case task names have changed later
        reg = {}
        for cls in cls._reg:
            if cls.run != NotImplemented:
                name = cls.task_family
                if name in reg and reg[name] != cls and \
                        reg[name] != cls.AMBIGUOUS_CLASS and \
                        not issubclass(cls, reg[name]):
                    # Registering two different classes - this means we can't instantiate them by name.
                    # The only exception is if one class is subclass of the other. In that case, we
                    # instantiate the most-derived class (this fixes some issues with decorator wrappers).
                    reg[name] = cls.AMBIGUOUS_CLASS
                else:
                    reg[name] = cls
        return reg

    @classmethod
    def get_global_params(cls):
        """Compiles and returns the global parameters for all :py:class:`Task`.

        :return: a ``dict of parameter name -> parameter.
        """
        global_params = {}
        for t_name, t_cls in cls.get_reg().iteritems():
            if t_cls == cls.AMBIGUOUS_CLASS:
                continue
            for param_name, param_obj in t_cls.get_global_params():
                if param_name in global_params and global_params[param_name] != param_obj:
                    # Could be registered multiple times in case there are subclasses
                    raise Exception(
                        "Global parameter %r registered by multiple classes".format(param_name))
                global_params[param_name] = param_obj
        return global_params.iteritems()
