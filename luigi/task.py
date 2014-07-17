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


class Task(object):
    """
    This is the base class of all Luigi Tasks, the base unit of work in Luigi.

    A Luigi Task describes a unit of work. The key methods of a Task, which must
    be implemented in a subclass are:

    * :py:meth:`run` - the computation done by this task.
    * :py:meth:`requires` - the list of Tasks that this Task depends on.
    * :py:meth:`output` - the output :py:class:`Target` that this Task creates.

    Parameters to the Task should be declared as members of the class, e.g.::

        class MyTask(luigi.Task):
            count = luigi.IntParameter()

    Each Task exposes a constructor accepting all :py:class:`Parameter` (and
    values as kwargs. e.g. ``MyTask(count=10)`` would instantiate `MyTask`.

    In addition to any declared properties and methods, there are a few
    non-declared properties, which are created by the :py:class:`Register`
    metaclass:

    ``Task.task_namespace``
       optional string which is prepended to the task name for the sake of
       scheduling. If it isn't overridden in a Task, whatever was last declared
       using `luigi.namespace` will be used.

    ``Task._parameters``
       list of ``(parameter_name, parameter)`` tuples for this task class.
    """
    __metaclass__ = Register
    _event_callbacks = {}

    # Priority of the task: the scheduler should favor available
    # tasks with higher priority value first.
    priority = 0

    @classmethod
    def event_handler(cls, event):
        """ Decorator for adding event handlers """
        def wrapped(callback):
            cls._event_callbacks.setdefault(cls, {}).setdefault(event, set()).add(callback)
            return callback
        return wrapped

    def trigger_event(self, event, *args, **kwargs):
        """
        Trigger that calls all of the specified events associated with this class.
        """
        for event_class, event_callbacks in self._event_callbacks.iteritems():
            if not isinstance(self, event_class):
                continue
            for callback in event_callbacks.get(event, []):
                try:
                    # callbacks are protected
                    callback(*args, **kwargs)
                except KeyboardInterrupt:
                    return
                except:
                    logger.exception('Error in event callback for {}.'.format(event))
                    pass

    @property
    def task_family(self):
        """
        Convenience method since a property on the metaclass isn't directly
        accessible through the class instances.
        """
        return self.__class__.task_family

    @classmethod
    def get_params(cls):
        """
        Returns all of the Parameters for this Task.
        """
        # We want to do this here and not at class instantiation,
        # or else there is no room to extend classes dynamically

        params = []
        for param_name in dir(cls):
            param_obj = getattr(cls, param_name)
            if not isinstance(param_obj, Parameter):
                continue
            params.append((param_name, param_obj))
        # The order the parameters are created matters. See Parameter class.
        params.sort(key=lambda t: t[1].counter)
        return params

    @classmethod
    def get_global_params(cls):
        """
        Returns the global parameters for this Task.
        """
        return [(param_name, param_obj)
                for param_name, param_obj in cls.get_params() if param_obj.is_global]

    @classmethod
    def get_nonglobal_params(cls):
        """
        Returns the non-global parameters for this Task.
        """
        return [(param_name, param_obj)
                for param_name, param_obj in cls.get_params() if not param_obj.is_global]

    @classmethod
    def get_param_values(cls, params, args, kwargs):
        """
        Get the values of the parameters from the args and kwargs.

        :param params: list of (param_name, Parameter).
        :param args: positional arguments
        :param kwargs: keyword arguments.
        :returns: list of `(name, value)` tuples, one for each parameter.
        """
        result = {}
        params_dict = dict(params)
        # In case any exception are thrown, create a helpful description of how the Task was invoked
        # TODO: should we detect non-reprable arguments? There will lead to mysterious errors
        exec_desc = '{0}[args={1}, kwargs={2}]'.format(cls.__name__, args, kwargs)

        # Fill in the positional arguments
        positional_params = [(n, p) for n, p in params if not p.is_global]
        for i, arg in enumerate(args):
            if i >= len(positional_params):
                raise parameter.UnknownParameterException(
                    '{0}: takes at most {1} parameters ({2} given)'.format(
                        exec_desc, len(positional_params), len(args))
                )
            param_name, param_obj = positional_params[i]
            result[param_name] = arg

        # Then the optional arguments
        for param_name, arg in kwargs.iteritems():
            if param_name in result:
                raise parameter.DuplicateParameterException(
                    '{0}: parameter {1} was already sset as a positional parameter'.format(
                        exec_desc, param_name)
                )
            if param_name not in params_dict:
                raise parameter.UnknownParameterException(
                    '{0}: unknown parameter {1}'.format(exec_desc, param_name)
                )
            if params_dict[param_name].is_global:
                raise parameter.ParameterException(
                    '{0}: can not override global parameter {1}'.format(exec_desc, param_name)
                )
            result[param_name] = arg

        # Then use the defaults for anything not filled in.
        for param_name, param_obj in params:
            if param_name not in result:
                if not param_obj.has_default:
                    raise parameter.MissingParameterException(
                        '{0}: requires the {1} parameter to be set'.format(
                            exec_desc, param_name)
                    )

        def list_to_tuple(x):
            """ Make tuples out of lists and sets to allow hashing.
            """
            if isinstance(x, list) or isinstance(x, set):
                return tuple(x)
            else:
                return x
        # Sort it by the correct order and make a list
        return [(param_name, list_to_tuple(result[param_name])) for param_name, param_obj in params]

    def __init__(self, *args, **kwargs):
        """
        Constructor to resolve values for all Parameters.
        For example, the Task::

            class MyTask(luigi.Task):
                count = luigi.IntParameter()

        can be instantiated as ``MyTask(count=10)``.
        """
        params = self.get_params()
        param_values = self.get_param_values(params, args, kwargs)

        # Set all values on class instance
        for key, value in param_values:
            setattr(self, key, value)

        # Register args and kwargs as an attribute on the class. Might be useful
        self.param_args = tuple(value for key, value in param_values)
        self.param_kwargs = dict(param_values)

        # Build up task id
        task_id_parts = []
        param_objs = dict(params)
        for param_name, param_value in param_values:
            if dict(params)[param_name].significant:
                task_id_parts.append(
                    '{0}={1}'.format(param_name, param_objs[param_name].serialize(param_value)))

        self.task_id = '{0}({1})'.format(self.task_family, ', '.join(task_id_parts))
        self.__hash = hash(self.task_id)

    def initialized(self):
        """ Returns ``True`` if the Task is initialized and ``False`` otherwise. """
        return hasattr(self, 'task_id')

    @classmethod
    def from_input(cls, params, global_params):
        """
        Creates an instance from a str->str hash

        This method is for parsing of command line arguments or other
        non-programmatic invocations.

        :param params: dict of param name -> value.
        :param global_params: dict of param name -> value, the global params.
        """
        for param_name, param in global_params:
            value = param.parse_from_input(param_name, params[param_name])
            param.set_default(value)

        kwargs = {}
        for param_name, param in cls.get_nonglobal_params():
            value = param.parse_from_input(param_name, params[param_name])

        return cls(**kwargs)

    def clone(self, cls=None, **kwargs):
        """
        Creates a new instance from an existing instance where some of the args have changed.

        There's at least two scenarios where this is useful (see test/clone_test.py)
        - Remove a lot of biler plate when you have recursive dependencies and lots of args
        - There's task inheritance and some logic is on the base class
        """
        k = self.param_kwargs.copy()
        k.update(kwargs.items())

        if cls is None:
            cls = self.__class__

        new_k = {}
        for param_name, param_class in cls.get_nonglobal_params():
            if param_name in k:
                new_k[param_name] = k[param_name]

        return cls(**new_k)

    def __hash__(self):
        return self.__hash

    def __repr__(self):
        return self.task_id

    def complete(self):
        """
        If the task has an outputs, return ``True`` if all outputs exist.
        Otherwise, return whether or not the task has run or not.
        """
        outputs = flatten(self.output())
        if len(outputs) == 0:
            # TODO: unclear if tasks without outputs should always run or never run
            warnings.warn('Task {} without outputs has no custom complete() method'.format(self))
            return False

        for output in outputs:
            if not output.exists():
                return False
        else:
            return True

    def output(self):
        """
        The output that this task produces.

        The output of the Task determines if the Task needs to be run --the task
        is considered finished if the outputs all exist. Subclasses should
        override this method to return a sinble :py:class:`Target` or list of
        :py:class:`Target` instances.
        """
        return []

    def requires(self):
        """
        The Tasks that this Task depends on.

        A Task will only run if all of the Tasks that it requires are completed.
        If your Task does not require any other Tasks, then you don't need to
        override this method. Otherwise, a Subclass can override this method
        to return a single Task, a list of Task instances, or a dict whose
        values are Task instances.
        """
        return []

    def _requires(self):
        """
        Override in "template" tasks which themselves are supposed to be
        subclassed and thus have their requires() overridden (name preserved to
        provide consistent end-user experience), yet need to introduce
        (non-input) dependencies.

        Must return an iterable which among others contains the _requires() of
        the subclass.
        """
        return flatten(self.requires())

    def input(self):
        """
        Returns the outputs of the Tasks returned by :py:metho:`requires`

        :return: a list of :py:class:`Target` objects which are specified as
                 outputs of all required Tasks.
        """
        return getpaths(self.requires())

    def deps(self):
        """
        Internal method used by the scheduler.
        Returns the flattend lsit of requires.
        """
        return flatten(self._requires())

    def run(self):
        """
        The task run method, to be overridden in a subclass.
        """
        pass

    def on_failure(self, exception):
        """
        Override for custom error handling

        This method gets called if an exception is raised in run.
        Return value of this method is json encoded and sent to the scheduler as the `expl` argument.
        Its string representation will be used as the body of the error email sent out if any.

        Default behavior is to return a string representation of the stack trace.
        """
        traceback_string = traceback.format_exc()
        return 'Runtime error\:\n{}'.format(traceback_string)

    def on_success(self):
        """
        Override for doing custom completion handling for a larger class of tasks.

        This method gets called when :py:metho:`run` 
        """
