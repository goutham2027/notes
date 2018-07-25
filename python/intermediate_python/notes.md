Note: Most of the text is from Intermediate Python book by Obi Ike-Nwosu

# Intermediate Python
  - Obi Ike-Nwosu

## Short Tutorial
Python supports anonymous functions with the lambda keyword.

```
square_of_number = lambda x: x ** 2
```

Tuple is an immutable object but the object in the tuple is mutable it
can be changed.


## Object 201
Python objects are the basic abstraction over data in Python; every
object has an identity (id), a type and a value.
id(obj) - is operator compares identity of two objects. In CPython id()
function returns an integer that is a memory location for the object.

Python is a pass by object reference language which means the values of
object references are the values passed to function or method calls and
names bound to variables refer to these reference values.
eg: x = [1,2,3]
y = x
 now changing y will also change x.

### Strong and Weak object references
Python objects get references when they are bound to names. This binding
can be in the form of an assignment, function/method call. Everytime an
object gets a reference, the reference count is increased.
`sys.getrefcount` gives the reference count for an object.

There are 2 types of references in Python: Strong, and Weak references.
But when discussing references, it is always strong reference.

Strong Reference: Whenever a strong reference is created, reference
count is increment. Garbage collector will collect only objects that
have a reference count of 0.

```
import sys
li = []
m = li
sys.getrefcount(li) # 3
```

Weak Reference: Do not increase reference count of the referenced
object. Weak referencing is provided by `weakref` module.

```
class Foo:
  pass

a = Foo()
b = a

sys.getrefcount(a) # 3

c = weakref.ref(a)
sys.getrefcount(a) # 3
```

When all the strong references to an object have deleted then the weak
reference looses it reference to the original object and the object is
ready for garbage collection.


### The Type Hierarchy
* None Type
* NotImplemented Type
* Ellipsis Type (used in numpy)
* Numeric Type
  + int
  + float
  + boolean
  + complex numbers
* Sequence Type
  + tuple
  + list
  + bytearray
  + strings
  + bytes
* Set
* Mapping
  + dict
* Callable Types
  These are the types that support the function call operation.
  Functions are not the only callable types in Python; any object type
  that implements the __call__ - special method is a callable type.
  The function callable(type) is used to check if the given type is
  callable.
  built-in callable types:
    - user defined functions
    - methods
    - classes
* Custom Type
  are created using class statement. Custom class objects have a type of
  type.
* Module Type
* File/IO Types
* Built-in Types
  - traceback objects
  - code objects
  - frame objects
  - slice objects
  - generator objects



## Ojbect Oriented Programming
The execution of a class statement creates a class object. At the start
of the execution of a class statement, a new name-space is created and
this serves as the name-space into which all class attributes go. Unlike
languages like Java, in Python does not create a new local scope that
can be used by class methods.

If the class created is an object then what is the class of the class
object? type class

Instance objects: result of instantiating class objects.

Method objects: If x is an instance of the Account class, x.deposit is
an example of a method object. Method objects are similar to functions
however during a method definition, an extra argument self is included in the
arguments list. This self argument refers to an instance of the class.
Why do we have to pass an instance as an argument to a method?
```
x = Account()
x.inquiry() # this is equal to Account.inquiry(x)
```
The method object is a wrapper around a function object; when the method
object is called with an argument list, a new argument list is
constructed from the instance object and the argument list, and the
underlying function object is called with this new argument list. This
applies to all instance method objects including __init__ method. self
argument is just a convention, anything can be used there.


### Special or magic methods
Python special methods are just ordinary methods with double underscores
as prefix and suffix to the method names.
eg: __init__
__getuten__ method invoked by the index, [] operator. a[i] is translated
to type(a).__getitem__(a,i)

Special methods for instance creation
The __new__ and __init__ special methods are the two methods that are
integral to instance creation.
New class instances creation is a two step process:
1) First static method __new__ is called to create and return a new
class instance
2) __init__ method is called to initialize

When do we need to override __new__ method is when sub-classing built-in
immutable types.

```
import math

class NextInteger(int):
    def __new__(cls, val):
        return int.__new__(cls, math.ceil(val))

print(NextInteger(5.2))
```

Attempting to do math.ceil operation in an __init__ method will cause
the object initialization to fail.

The __new__ method can also be overridden to create a Singleton super
class; subclasses of this class can only ever have a single instance
throughout the execution of a program.

# TODO
check `singleton.py`. Didn't understand this concept.


Special methods for attribute access

For customizing attribute references; this maybe inorder to access or
set such an attribute.

1) __getattr__: used when a referenced attribute cannot be found. Only
called when an attribute that is referenced is neither an instance
attribute nor is it found in the class tree of that object. This method
should return some value for the attribute or raise an AttributeError
exception.
2) __setattr__
3) __delattr__
4) __dir__: this is implemented to customize the list of object
attributes returned by a call dir(obj)


Special methods for Type emulation
eg: __add__, __sub__ etc

Sequence and Mapping types special methods
Sequences and mapping are often referred to as container types because
they can hold references to other objects.
1) __len__(obj)
2) __getitem__(obj, key)
3) __setitem__(obj, key, value)
4) __delitem__(obj, key)
5) __contains__(obj, key) # invoked by key in obj.
6) __iter__(self) # called when an iterator is required for a container.
for mappings, it should iterate over the keys of the container. Also
used by for..in construct.


Emulating Callable Types
Callable types support the function call syntax (args). Classes that
implement the __call__(self, args) method are callable.


### Inheritance
super - `super().__init__(name, balance)

Multiple Inheritance
In multiple inheritance, class can have multiple parent classes. This
type of hierarchy is strongly discouraged. Complexity is resolving the
methods.

A method resolution algorithm determines how a method is found in a
class or any of the class' base classes. In Python, the resoluton order
is calculated at class definition time and stored in the class__dict__
as the __mro__ attribute.

eg:
```
class A:
  pass

class B(A):
  pass


class C(A):
  pass


class D(B, C):
  pass
```
To obtain an mro, the interpreter method resolution alogrithm carries
out a left to right depth first listing of all classes in the hierarchy.
In the example above, this results in the following class list [D, B, A,
C, A, object]

By default all classes inherit from the root `object` class. For classes
that occurs multiple times, all occrrences are removed except the last
occurrence resulting in an mro of [D, B, C, A, object].

For eg `A` defines a method say `save` that is overridden by `B, C and D`. Suppose
that there is a requirement that all the methods are called. A
combination of `super` and `__mro__` provide the ammunition for solving
the problem. This solution is referred to as the `call-next` method by
Guido van Roussum.

```
class A(object):
  def meth(self):
    print("Saving A's data")

class B(A):
  def meth(self):
    super(B, self).meth()
    print("Saving B's data")

class C(A):
  def meth(self):
    super(C, self).meth()
    print("Saving C's data")

class D(B, C):
  def meth(self):
    super(D, self).meth()
    print("Saving D's data")
```
When self.meth() is called by an instance of D -
`super(D, self).meth()` will find and call `B.meth(self)`, since B is
the first base class following D that defines meth in `D.__mro__`= [D,
B, C, A].
In `B.meth`, `super(B, self).meth()` is called and self is an instance
of `D`, the next after B is C in the mro. This finds `C.meth` and which
inturn calls `super(C, self).meth()`. In mro, the next after `C` is `A`
and thus `A.meth` is called.

### Static methods
Static methods are decorated with `@staticmethod`.

Static methods are normal functions that exist in the name-space of a
class. Referencing a class method from a class shows that rather than an
unbound method type, a function type is returned.
Example in `ex_static.py`.

Static methods provide a mechanism for better organization to put
related code in a class and can be overridden in a sub-class.

### Class methods
Class methods are decorated with `@classmethod`. Class methods operate
on classes and the first parameter passed to class methods is a class.

The usage of class methods is a factory for object creaton. For eg, the
data for the Account class comes in different formats such as tuples,
json string etc. It is not possible to define multiple __init__ methods
in a class so class methods come in handy.

In the Account example if the initial data comes from a json string
object we define a class factory method, `from_json` that takes a json
string object and handles the extraction of params and creation of the
account object using the extracted params.
```
@classmethod
def from_json(cls, params_json):
  params = json.loads(params_json)
  return cls(params.get("name"), params.get("balance"))
```


### Descriptors and Properties
Descriptors are used widely in the code of the Python language.

Different cases where Descriptors could be used.
1) Type checking of object attributes. It becomes cumbersome to check
for many attributes.
```
def __init__(self, name, age):
  if isinstance(name, str):
    self.name = name
  else:
    raise TypeError("Must be a string")
  if isinstance(age, int):
    self.age = age
  else:
    raise TypeError("Must be an int")
```

2) A program that needs object attributes to be read-only once
initialized.

3) A program in which the attribute access need to be customized. This
may be to log such attribute access or to even perform some kind of
transformation of the attribute.


All the above mentioned issues are all linked together by attribute
reference.

### Python Descriptors
A Descriptor is an object that represents the value of an attribute.
This means that if an account object has an attribute name, a descriptor
is another object that can be used to represent the value held by the
attribute, name.

Descriptor object can implement __get__, __set__ or __delete__ special
methods of the descriptor protocol.
```
descr.__get__(self, obj, type=None) returns a value
descr.__set__(self, obj, value) returns None
descr.__delete__(self, obj) returns None
```

Objects implements only the __get__ method are non-data descriptors so
they can only be read from after initialization while objects
implementing the __get__ and __set__ are data descriptors - writable.

implementing type checking: `ex_descriptors.py`

Descriptors provide the mechanism behind properties, static methods,
class methods, super and a host of other functionality in Python
classes. In fact, descriptors are the first type of object searched for
during an attribute reference.
Wehn an object is referenced, a reference, b.x is transformed into
type(b).__dict__['x'].__get__(b, type(b))


### Class Properties
Defining descriptor classes each time a descriptor is required is
cumbersome. Python properties provide a concise way of adding data
descriptors to attributes.
`property(fget=None, fset=None, fdel=None, doc=None) returns property
attribute`
fget, fset and fdel are the getter, setter and deleter methods.
example `ex_class_properties.py`

Python also provides the @property decorator that can be used to create
read only attribute. A property object has getter, setter and deleter
decorator methods that can be used to create a copy of the property with
the corresponding accessor function set to the decorated function.
example: `ex_property.py`


methods are stored as ordinary functions in a class dictionary. Where as
object methods are of bound method type

eg:
```
Account.inquiry # <function Account.inquiry at ..>

x = Account("Goutham", 10)
x.inquiry # <bound method Account.inquiry of <Account object at ...>>

A bound methods is just a thin wrapper around the class function.
Functions are descriptors because they have the __get__ method attribute
so a reference to a function will result in a call to the __get__ method
of the function and this returns the desired type.

Is this why functions are also objects in Python.

### Abstract Base Classes
Sometimes, it is necessary to enforce a contrat betweeen classes in a
program.
For eg, it may be necessary for all classes to implement a set of
methods.

Interfaces, abstract classes in Java.

One way of implementing is define a base class and each of the methods
raise "NotImplementedError" exception. But this doesn't solve the
problem (we will not know whether there is an implementation for the
method) completely as we will not know until the method call.
```

Python's Abstract base classes solve these issues. This functionality is
provided by the abc module. This module defines a meta-class and a set
of decorators that are used in the creation of Abstract base classes.
@abstractmethod and @abstractproperty decorators to create methods and
properties that must be implemented by non-abstract subclasses.

** If a subclass doesn't implement any of the abstract methods or
properties then it is also an abstract class and cannot be instantiated.

eg: `ex_abc_1.py`
This will throw an error.

The second implementation `ex_abc_2.py` implements all the abstract base
class methods.

Abstract base classes also allow classses to register as part of its
hierarchy but it performs no check on whether the sub class implement
all the methods and properties that have been marked as abstract.

eg: `Vehicle.register(Car)`

## The Function
Python functions are either named or anonymous. In Python, functions
are first class objects.

Which means, functions can be:
 - used as values,
 - assigned to variables
 - used as arguments to other function
 - returned from method/function calls just like python value eg:
   string, numbers.


### Function Definitions
```
def square(x):
  return x**2
```

When a function definiton gets encountered, only the function definition
statement, the is `def square(x)` is executed, this implies that all
arguments are evaluated. The function definition does not execute
the function body; this gets executed only when the function is called.

# TODO
** Try to understand this paragraph

The execution of a function definition binds the function name in
the current name-space to a function object which is a wrapper around
the executable code for the function. This function object contains a
reference to the current global namespace which is the global
name-space that is used when the function is called.

Python also supports anonymous functions which are created using
`lambda` keyword.
```
square = lambda x: x**2
```

Lambda expressions return function objects after evaluation and
have the same attributes as named functions. Used for very simple
functions.

### Functions are objects
`type(square) # <class function>`

Like every other object, introspection on function using dir()
function provides a full list of function attributes.

Functions that take other functions as arguments are commonly referred
to as higher order functions and these form a very important part of
functional programming.
Example for higher order functions is the map function takes a function and an iterable and
applies the function to each item in the iterable returning a new
list.
eg: `map(square, range(10)) # [0, 1, 4, ... 81]`

A function can be defined inside another function as well as returned
from a function call.

### Functions are descriptors
Functions has `__get__` method making them non-data descriptors.

The `__get__` method is called whenever a function is referenced and
provides a mechanism for handling method calls from object and ordinary
function calls. This descriptor characteristic of a function enables
functions to return either itself or a bound method when referenced
depending on where and how it is referenced.

### Calling functions

Python functions can also be called using variable number of arguments
than the regular way with normal arguments.
These variable number of arguments come in 3 flavors.

+ Default Argument Values:
eg:
```
def def_args(arg, arg1=1, arg2=2)
```

To be careful when using mutable data structures as default arguments.
Function definition get executed only once so these mutable data
structures are created once at definition time. This means that the same
data structure is used for all function calls.

eg:
```
def show_args_using_mutable_defualts(arg, def_arg=[]):
  def_arg.append('Hello world!")
  return "arg={}, def_arg={}".format(arg, def_arg)

>>> show_args_using_mutable_defualts("test")
"arg=test, def_arg=['Hello World']"

>>> show_args_using_mutable_defualts("test")
"arg=test, def_arg=['Hello World', 'Hello World']"
```

On every function call, "Hello World~" is added to the `def_arg` list.

+ Keyword Argument:
Functions can be called using keyword argumets of the form
`kwarg=value`.
In a function call, keyword arguments must not come before non-keyword
arguments.

+ Arbitary arguments list:
`*args`
The arguments are all bundled together into a tuple that can be
accessed via the `args` argument.
`def abc(file, sep, *args)`

### Unpacking function argument
Arguments for a function call are either in a tuple, a list or a dict.
These arguments can be unpacked into functions for function calls
using * or ** operators.

If the values for a function call are in a list then these values can be
unpacked directly into the function using *list
eg:
```
def print_args(a,b):
  print a
  print b

>>> args = [1,2]
>>> print_args(*args)
```

Dictionaries can be used to store `keyword to value` mapping and the **
operator is used to unpack the keyword arguments to the functions
eg:
```
def parrot(voltage, state='a stiff', action='voom'):
  pass

>>> d = {'voltage': '4', 'state': 'abc', 'action': 'VOOM'}
>>> parrot(**d)
```

*args - represents unknown length of sequence of postiional
arguments while **kwargs represents a dict of keyword name value
mappings which may contain any amount of keyword name value mapping.

*args must come before **kwargs in the function definition.

### Nested functions and closures
eg: `ex_nested_functions.py`
In nested functions such as in the example, a new instance of the nested
function is created on each call to outer function. This is because
during each execution of the make_counter function, the definition of
the counter function is executed but the body is not executed.

When nested functions reference variables from the outer function in
which they are defined, the nested function is said to be closed over
the referenced variable.

python3 introduced `nonlocal` keyword to fix the closure scoping issue.
Closures can be used for maintaining states and for simple cases provide
a more succinct and readable solutions than classes.
eg: `ex_closures.py`

### A byte of functional programming
side effects: object values do not change once they are created and to
reflect a change in an object value, a new object with the changed value
is created.
example of a function with side effects
```
def squares(numbers):
  for i, v in enumerate(numbers):
    numbers[i] = v**2
  return numbers
```

The same function can be implemented without any modification to
arguments and create new values.
```
def squares(numbers):
  return map(lambda x: x**x, numbers)
```

Python provides built-in functions such as map, filter and reduce that
aid in functional programming.

1) map(func, iterable): takes a function and an iterable and returns an
interator that applies the function to each item in the iterable.

2) filter(func, iterable): returns an iterator that applies func to each
element of the iterable and returns elements of the iterable for which
the result of the application is True.
eg:
```
even = lambda x: x%2 == 0
even(10) # True
list(filter(even, range(10)) # [0, 2, 4, 6, 8]
```

3) reduce(func, iterable[, initializer]): moved to functools in python3.
reduce function applies func cumulatively to the items in iterable in
order to get a single value and returns that value. `func` takes two
positional arguments.
eg:
```
reduce(lambda x, y: x+y, [1,2,3,4,5])
# calculates ((((1+2)+3)+4)+5)
```
reduce function starts with reducing the first 2 args then reduces the
third with the result of the first two and soon.
eg:
```
def flatten_list(nested_list):
  add_lists = lambda x,y: x + y
  return reduce(add_lists, nested_list, [])

flatten_list([1,2,3], [4,5,6]) # [1,2,3,4,5,6]
```

#### Comprehensions
Python comprehensions are syntactic constructs that enable sequences to
be built from other sequences in a clear and concise manner.
3 types of comprehensions:
- List comprehension
- Set comprehension
- Dictionary comprehension

List comprehension

provide a concise way to create new list of elements that satisfy a
given condition from an iterable.
eg:
`squares = [x*x for x in range(10)]

Set Comprehensions
eg: `x = {i**2 for i in range(10)}`

Dict Comprehensions
eg: `x = {i: i**2 for i in range(10)}`

#### Functools
`functools` module in Python contains a few higher order functions that
act on and return other functions.

partial(func, *args, **kwargs): when called returns an object that can
be called like the original `func` argument with *args and **kwargs as
arguments.
eg:
```
from functools import partial
basetwo = partial(int, base=2)
basetwo.__doc__ = "convert base 2 string to an int"
basetwo('100') # 4
```
A new callable `basetwo` takes a number in binary and converts
to decimal. int() functions that takes two arguments has been wrapped by
a callable, `basetwo` that takes only one argument.

singledispatch: decorator that changes a function into a single
dispatch generic function. Handle dynamic overloading in which a
single function can handle multiple types.
A generic function is defined with the @singledispatch function,
the register decorator is then used to define functions for each type
that is handled. Dispatch to the correct function is carried out based
on the type of the first arg to the function call - hence the name
single generic dispatch
eg:
```
@singledispatch
def fun(arg, verbose=False):
  if verbose:
    print("Let me just say", end="")
  print(arg)

@fun.register(int)
def _(arg, verbose=False):
  if verbose:
    print("Strength in number", end="")
  print(arg)
```

## Iterators and Generators

