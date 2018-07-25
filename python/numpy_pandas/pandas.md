pandas is designed for working with tabular or heterogeneous data. NumPy
is best suites for working with homegeneous numerical array data.

### pandas data structures

#### Series
1 dimensional array and an associated array of data labels - index.

obj = pd.Series([4, 7, -5])

obj2 = pd.Series([4, 7, -5], index=['d', 'b', 'a', 'c'])

Another way to think about a Series is as a fixed-length, ordered dict,
as it is a mapping of index values to data values.
