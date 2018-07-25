what's the difference between numpy array and python list?


x = np.arange(10)

np.full((3,3), True, dtype=bool) # create a 3x3 numpy array of all
True's

Python for Data Analysis
  Data wrangling with Pandas, NumPy, and Ipython
  Wes McKinney

Github repository: http://github.com/wesm/pydata-book


Numpy: Numerical Python. Numerical computing in Python.

Pandas: panel data. provide high-level data structures and functions designed to
make working with structured or tabluar data fast, easy and expressive.

Matplotlib: for producing plots and other 2 dimensional data
visualizations.

Ipython and Jupyter

SciPy: collection of packages addressing a number of different standard
problem domains in scientific computing.

scikit-learn: premier general purpose machine learning toolkit for
Python programmers.

statsmodels: is a statistical analysis package


=======================

Jupyter Notebook: One of the major components of the Jupyter project is
the notebook, a type of interactive document for code, text, data
visualizations and other output. The Jupyter notebook interacts with
kernels, which are implementations of the Jupyter interactive computing
protocol in any number of programming languages. Python's Jupyter kernel
use the Ipython system for its underlying behavior.

`jupyter notebook`

Object introspection: add `?` in the end. and `??` will also show the
function's source code if possible.

`%run` command to execute a python file inside IPython environment.

`%paste` and `%cpaste` foolproof methods.


### shortcuts
Ctrl+A - Move cursor to beginning of line
Ctrl+E - Move cursor to end of line
Ctrl+K - Delete text from cursor until end of line
Ctrl+U - Discard all text on current line
Ctrl+F - Move cursor forward one character
Ctrl+B - Move cursor back one character
Ctrl+L - Clear screen


### IPython's Magic commands
To faciliate common tasks. A magic command is prefixed by `%`.

Execution of any Python statement: `%timeit <statement`

`%matplotlib` magic function configures its integration with the IPython
shell or Jupyter notebook.

`isinstance` can accept a typle of types.
eg: `isinstance(a, (int, float))`

You can concatenate tuples using the + operator to produce longer tuples

Multiplying a tuple by an integer, as with lists, has the effect of concatenating together that many copies of the tuple

The built-in bisect module implements binary search and insertion into a
sorted list.
`bisect.bisect` finds the location where an element should be inserted
to keepit sorted.
`bisect.insort` actually inserts the element into that location.


enumerate: to track index of the current item.

Generators: a generator is a concise way to construct a new iterable
object. Whereas normal functions execute and return a single result at a
time, Generators return a sequence of multiple results lazily, pausing
after each one until the next one is requested.

```
def squares(n=10):
  for i in range(1, n+1):
    yield i ** 2

```

Generator expressions:
`gen = (x ** 2 for x in range(100))`

### NumPy
NumPy - Numerical Python

NumPy is designed for efficiency on large array of data.

NumPy internally stores data in a contiguous block of memory,
independent of other built-in Python objects.NumPy's library of algorithms written in C
can operate on this memory without any type checking or other overhead.
NumPy arrays also use less memory than built-in Python sequences.

NumPy operations perform complext computations on entire arrays without
the need for Python `for` loops.

Eg:
```
import numpy as np
my_arr = np.arange(100000)
my_list = list(range(100000))
%time for _ in range(10): my_arr2 = my_arr * 2
%time for _ in range(10): my_list2 = [x*2 for x in my_list]
```

- NumPy based algorithms are generally 10 to 100 times faster than their
  pure Python counterparts and use significantly less memory.


ndarray - A multidimensional array object for homogeneous data.
Arrays enable us to perform mathematical operations on whole blocks of
data using syntax equivalent operations between scalar elements
```
data = np.random.randn(2,3)
data * 10 # all elements have been multiplied by 10
```

shape: Every array has a shape, a tuple indicating the size of each
dimension

dtype: data type of the array

ndim: dimensions of the array

creating ndarrays
====================
```
data = [1, 2, 3]
arr1 = np.array(data)

data2 = [[1, 2, 3], [4, 5, 6]]
arr2 = np.array(data2)
```

```
arr = np.zeros(10)
arr = np.zeros((10, 10))
np.ones(10)
np.empty(10) # np.empty will not always return an array of all zeros. In
some cases, it may return uniinitialized garbage values.

np.arange(10) # array valued version of the built-in Python range
function.

np.array([1, 2, 3], dtype=np.flaot64)
arr.astype(np.float64)
```

* Calling `astype` always creates a new array (a copy of the data), even
  if the new dtype is the same as the old type.

* Arrays enable us to express batch operations on data without writing
  any for loops. NumPy uses vectorization. Any arithmetic operations
  between equal-size arrays applies the operation element-wise.

* Operations between differently sized arrays is called broadcasting.

Basic indexing and slicing
==============================
Indexing can be done in multiple ways. For 1-d arrays it is same as
Python lists.

```
arr = np.arange(10)
arr[5:8]
arr[5:8] = 12 # the value is propogated (broadcasted) to the entire
selection.
```

* An important distinction from Python's built in lists is that array
slices are views on the original array. This means that the data is not
copied, and any modifications to the view will be reflected in the
source array.
eg:
```
arr_slice = arr[5:8]
arr_slice # array([12,12,12])
arr_slice[1] = 12345
arr # array([0,1,2,3,4,12, 12345,12,8])
arr_slice[:] = 64 # bare slice: will assign all the values in an array.
```

If you want a copy of a slice of an ndarray instead of a view use
arr[5:8].copy()

Boolean indexing
```
names = np.array(['Bob', 'Joe', 'Will', 'Bob', 'Will', 'Joe'])
names == 'Bob'

data = np.random.randn(6,4)
data[names == 'Bob']

names != 'Bob'
data[~(names=='Bob')]

mask = (names == 'Bob') | (names == 'Will')
```

Selecting data from an array by boolean indexing always creates a copy
of the data, even if the returned array is unchanged.

Python Keywords `and` and `or` do not work with boolean arrays. Use `&`
and `|` instead.

Setting values with boolean arrays.
```
data[data < 0] = 0
```

Fancy indexing - is a term adopted by NumPy to describe indexing using
integer arrays.
eg: 8x4 array
```
arr = np.empty(8,4)

for i in range(8):
  arr[i] = i

# to select out a subset of the rows in a particular order, we can
simply pass a list of ndarray of integers specifying the desired order
arr[[4, 3, 0, 6]]
```

Transposing arrays and swapping axes

Transposing is a special form of reshaping that similarly returns a view
on the underlying data without copying anything.
Arrays have the `transpose` method and also the `T` attribute.
```
arr = np.arange(15).reshape((3,5))
arr.T
```

For higher dimensional arrays, `transpose` will accept a tuple of axis
numbers to permute the axes.
```
arr = np.arange(16).reshape((2,2,4))
arr.transpose((1,0,2))
```

### Universal functions: Fast element wise array functions
A universal function (`ufunc`) is a function that performs element-wise
operations on data in ndarrays.

Many `ufuncs` are simple element-wise transformations like `sqrt` or
`exp`.

## May 16, 2018 - Wednesday
### Array-Oriented programming with Arrays
- Numpy array enables us to express many kinds of data processing tasks as
concise array expressions that might otherwise require writing loops.
This practice of replacing explicit loops with array expressions is
commonly referred to as vectorization.

- Vectorized array operations will often be one or two times faster than
  pure Python equivalents.

- The `numpy.where` function is vectorized version of the ternary
  expression `x if condition else y`.

```
xarr = np.arange(1.1, 1.6, 0.1)
yarr = np.arange(2.1, 2.6, 0.1)
cond = np.array([True, False, True, True, False])

# take a value form xarr whenever the corresponding value in cond is True else take value from yarr

# list comprehension implementation
# cons with this approach:
# - Will not be very fast for large arrays because all the work is being done in interpreted Python code.
# - This implementation will not work with multi dimensional arrays.
result = [(x if c else y) for x, y, c in zip(xarr, yarr, cond)]
print("list comprehension: {}".format(result))
%timeit [(x if c else y) for x, y, c in zip(xarr, yarr, cond)]

# np.where implementation
result = np.where(cond, xarr, yarr)
print("np.where: {}".format(result))
%timeit np.where(cond, xarr, yarr)
```

- 2nd and 3rd arguments to `np.where` can also be scalars.

### Mathematical and Statistical methods
```
# Mathematical and Statistical methods
arr = np.random.randn(4,5)
print(arr)
print("Mean: {}".format(arr.mean()))
print("Sum: {}".format(arr.sum()))

# Functions like mean, sum take an optional axis argument that computes the
# statistic over the given axis.

# the following 2 lines compute mean, sum across the columns
print("Mean on axis=1: {}".format(arr.mean(axis=1)))
print("Sum on axis=1: {}".format(arr.sum(axis=1)))

# compute sum across the rows
print("Sum on axis=0: {}".format(arr.sum(axis=0)))
```

### Methods for Boolean arrays
```
# methods for boolean array

arr = np.random.randn(100)
bools = arr > 0

# boolean values are coerced to 1 (True) or 0 (False).
# sum is often used as a means of counting True values in a boolean arrzay.
(arr > 0).sum()

# any, all

print("any: {}".format(bools.any()))

print("all: {}".format(bools.all()))
```

### unique, intersection
```
arr1 = np.random.randint(10, size=(100))
arr2 = np.random.randint(5, size=(100))

# unique
print(np.unique(arr))

# intersection
print(np.intersect1d(arr1, arr2))
```

### File input and output with Arrays
