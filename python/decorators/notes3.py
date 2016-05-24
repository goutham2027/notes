# Decorating methods
# In Python, methods are functions that expect their first parameter to be a
# reference to the current object. We can build decorators for methods the same
# way, while taking self into consideration in the wrapper function.

def p_decorate(func):
    def func_wrapper(self):
        return "<p>{0}</p>".format(func(self))

    return func_wrapper

class Person(object):
    def __init__(self):
        self.name = 'John'
        self.family = 'Doe'

    @p_decorate
    def get_fullname(self):
        return self.name + " " + self.family

my_person = Person()
print my_person.get_fullname()

# Better approach would be to make our decorator useful for functions and
# methods alike. This can be done by putting *args, **kwargs as parameters
# for the wrapper, then it can accept any arbitrary number of arguments and
# keyword arguments.

def p_decorate(func):
    def func_wrapper(*args, **kwargs):
        return "<p>{0}</p>".format(func(*args, **kwargs))

    return func_wrapper

class Person(object):
    def __init__(self):
        self.name = "John"
        self.family = "Doe"

    @p_decorate
    def get_fullname(self):
        return self.name + " " + self.family

my_person = Person()
print my_person.get_fullname()

# Passing arguments to decorators
# In one of the examples, we decorated a function with 3 decorators. (p, div,
# strong). All have same functionality but wrapping the string with different
# tags. We can have a general implementation.
def tags(tag_name):
    def tags_decorator(func):
        # works with *args, **kwargs also.
        # def func_wrapper(name):
        def func_wrapper(*args, **kwargs):
            return "<{0}>{1}</{0}>".format(tag_name, func(*args, **kwargs))
        return func_wrapper
    return tags_decorator

# tags("tag") returns tags_decorator function.
# tags_decorator takes function as argument and returns func_wrapper
@tags("div")
@tags("strong")
@tags("p")
def get_text(name):
    return "Hello" + name

# Decorators expect to receive a function as an argument, that why we will have
# to build a function that takes those extra arguments and generate our
# decorator on the fly. In the example above tags is our decorator generator.


print get_text("Goutham")
print get_text.__name__
print get_text.__module__
print get_text.__doc__


# Debugging decorated functions
# decorators are just wrapping our functions, in case of debugging that can
# be problamatic since the wrapper function does not carry the name, module
# and doc string of the original function.
# print get_text.__name__ # outputs func_wrapper.
# the output was expected to be get_text yet, the attributes __name__, __doc__,
# and __module__ of get_text got overridden by those of the
# wrapper(func_wrapper). We can re-set them within func_wrapper but Python
# provides a much nicer way.

# Functools to the resuce.
# functools module contains functools.wraps
# Wraps is a decorator for updating the attributes of the wrapping function
# (func_wrapper) to those of the original function(get_text).
# This is as simple as decorating func_wrapper by @wraps(func).
# Example:
from functools import wraps

def tags(tag_name):
    def tags_decorator(func):
        @wraps(func)
        def func_wrapper(name):
            return "<{0}>{1}</{0}>".format(tag_name, func(name))
        return func_wrapper
    return tags_decorator

@tags("p")
def get_text(name):
    """ returns some text """
    return "Hello" + name

print "After functools wraps"
print get_text.__name__
print get_text.__doc__
print get_text.__module__
