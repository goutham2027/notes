writing proto files

* `package` - to avoid naming conflicts
* `message` - aggregate containing a set of typed fields
  standard simple data types - `bool`, `int32`, `float`, `double`, `string`
  other messages can also be used field types
  we can also define `enum` types if we want values to be from
  predefined list
* `= 1`, `= 2` markers on each element identify the unique tag that field
uses in binary encoding. Tag numbers 1-15 require one less byte to encode than higher
  numbers.

  tip: for commonly used or repeated elements use tag numbers 1-15, for
  less frequent fields use 16 and higher.
* Each field must be annotated with one of the following modifiers
  + `required`
  + `optional`
  + `repeated`

  tip: using `required` does more harm than good.



References:
https://developers.google.com/protocol-buffers/docs/pythontutorial
