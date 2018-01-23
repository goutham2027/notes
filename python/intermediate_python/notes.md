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


