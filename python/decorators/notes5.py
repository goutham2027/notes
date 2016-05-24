# Understanding Python decorators in 12 easy steps
# http://simeonfranklin.com/blog/2012/jul/1/python-decorators-in-12-steps/

# 1. Functions
def foo():
    return 1

# 2. Scope
# In Python functions create a new scope (i.e; functions have their own namespace)
# This means Python looks first in the namespace of the function to find variable
# names when it encounters them in the function body.
a_string = "This is a global variable"

def foo():
    print locals()

print globals()
# The builtin globals function returns a dictionary containing all the
# variable names Python knows about.

# 3. Variable resolution rules
# Python's scope rule is that varianle creation always creates a new local
# variable but variable access (including modification) looks in the local
# scope and then searches all the enclosing scopes to find a match.
a_string = "Global variable"
def foo():
    print a_string

foo() # prints Global variable

def foo():
    a_string = 'test'
    print locals()

foo() # { 'a_string': 'test'}
a_string # prints Global variable

# 4. Variable Lifetime
def foo():
    x = 1
# print x # name space error
# The namespace created for foo function is created from scratch each time
# the function is called and it is destroyed when the function ends.

# 5. Function arguments and parameters
# The parameter names become local variables in our function.
# Function parameters can be either
# * positional parameters that are mandatory
# * named, default value parameters that are optional
{'x': 1}
def foo(x):
    print locals()

foo(1)
{'x': 1}

# 6. Nested functions
# Python allows the creation of nested functions. This means we can declare
# functions inside of functions and all the scoping and lifetime rules still
# apply normally
def outer():
    x = 1
    def inner():
        print x
    inner()

outer() # 1
# Python looks for a local variable named x, failing it then looks in the
# enclosing scope which is another function. It finds there.

# 7. Functions are first class objects in Python
# In Python, functions are objects like everything else.
# All objects in Python inherit from a common baseclass
issubclass(int, object) # True

def foo():
    pass

foo.__class__ # <type 'function'>
issubclass(foo.__class__, object) # True
# Functions has attributes - functions are objects in Python, just like
# everything else. Classes are also objects in Python.
# Functions are just regular values like any other kind of value in Python.
# That means we can pass functions to functions as arguments or return
# functions from functions as return values.
def add(x,y):
    return x + y

def sub(x,y):
    return x - y

def apply(func, x, y):
    return func(x,y)

apply(add, 2, 1) # 3
apply(sub, 2, 1) # 1

# we can also return functions
def outer():
    def inner():
        print "Inside Inner"
    return inner

foo = outer() # returns value which is the function inner.
foo() # Inside Inner.

# 8. Closures
def outer():
    x = 1
    def inner():
        print x # 1
    return inner

foo = outer()
foo.func_closure

# Everything works according to Python's scoping rules - x is a local variable
# in function outer. When inner printx x, Python looks for a local variable to
# inner and not finding it looks in the enclosing scope which is the function
# outer (finds it there).

# But what about variable lifetime? variable x is local to the function outer
# which means it only exists while the function outer is running. We aren't able
# to call inner till after the return of outer so according to our model of
# how Python works, x shouldn't exist anymore by the time we call inner and
# perhaps a runtime error of some kind should occur.

# But inner function does work. Python supports a feature called function
# closures which means that inner functions defined in non-global scope
# remembers what their enclosing namespaces looked like at definition time.
# This can be seen by looking at the func_closure attribute of our inner
# function which contains the variables in the enclosing scopes.

# 9. Decorators
# A decorator is just a callable that takes a function as an argument and
# returns a replacement function.
def outer(some_func):
    def inner():
        print "before some func"
        ret = some_func()
        return ret + 1
    return inner

def foo():
    return 1

decorated = outer(foo)
decorated()
# before some func
# 2

# We could say that variable decoated is a decorated version of foo - it's foo
# plus something.
foo = outer(foo)
# Now any calls to foo() won't get the original foo, they'll get the decorated
# version.
class Coordinate(object):
    def __init__(self, x, y):
        self.x = x
        self.y = y

    def __repr__(self):
        return "Coord: " + str(self.__dict__)

def add(a,b):
    return Coordinate(a.x + b.x, a.y + b.y)

def sub(a,b):
    return Coordinate(a.x - b.x, a.y - b.y)

one = Coordinate(100, 200)
two = Coordinate(300, 200)

add(one,two)
# Coord: {'y': 400, 'x': 400}

# But what if our add, sub functions have some bounds checking behavior.
# We can only sum or subtract based on positive coordinates and any result
# should be limited to positive coordinates as well
# currently
one = Coordinate(100, 200)
two = Coordinate(300, 200)
three = Coordinate(-100, -100)
sub(one, two)
# Coord: {'y': 0, 'x': -200} # expected y, x to be 0
add(one, three)
# Coord: {'y': 100, 'x': 0} # expected y 200, x 100

# Instead of adding bounds checking to the input arguments of each function
# and return value of each function let's write a bounds checking decorator.
def wrapper(func):
    def checker(a,b):
        if a.x < 0 or a.y < 0:
            a = Coordinate(a.x if a.x > 0 else 0, a.y if a.y > 0 else 0)
        if b.x < 0 or b.y < 0:
            b = Coordinate(b.x if b.x > 0 else 0, b.y if b.y > 0 else 0)
        ret = func(a, b)
        if ret.x < 0 or ret.y < 0:
            ret = Coordinate(ret.x if ret.x > 0 else 0, ret.y if ret.y > 0 else 0)
        return ret
    return checker
add = wrapper(add)
sub = wrapper(sub)
sub(one, two)
add(one, three)

# 10. The @ symbol applies a decorator to a function
@wrapper
def add(a,b):
    return Coordinate(a.x + b.x, a.y + b.y)

# 11. *args and **kwargs
# Instead of fixed params accept all args

# 12 More generic decorators
def logger(func):
    def inner(*args, **kwargs):
        print "Arguments were: %s, %s" % (args, kwargs)
        return func(*args, **kwargs)
    return inner

@logger
def foo1(x, y=1):
    return x + y

