- Ruby is a genuine object-oriented language. Everything we manipulate is
an object, and the results of manipulations are themselves objects.

- Every object has a unique object identifier - object ID

- methods are invoked by sending a message to an object. The message
  contains the method's name, params method need. When an object
  receives a message, it looks into its own class for a corresponding
  method. If found, that method is executed.
  eg: `"goutham".length' -> 7
  the thing before the period is called the receiver, and the name after
  the period is the method to be invoked.

  In Java, we'd find the abs value of a number by calling a separate
  function and passing in that number. eg: `Math.abs(number)`

  In Ruby, the ability to determine an absolute number is built into
  numbers. We simply send the message abs to a number object and it will
  do its work.

  In double quotes strings, substitutions, expression interpolation can take place.

  `$var` - global variable
  `@var` - an instance variable
  `@@var` - class variable


