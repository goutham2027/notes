Reference:
https://developer.mozilla.org/en-US/docs/Web/JavaScript/A_re-introduction_to_JavaScript

A reintroduction to JS
* JS supports object oriented programming with object prototypes,
  instead of classes.
* JS also supports functional programming - functions are objects,
  giving functions the capacity to hold executable code and be passed
  around like any other object.

  Types:
  * Number
  * String
  * Boolean
  * Function
  * Object
  * Symbol (new in ES2015)
  * undefined
  * null

  functions are special type of object.

  New variables in js are declared using one of 3 keywords: let, const,
  var

  let - block-level variables. available from the function block it is
  enclosed in.
  eg: for(let aLetVariable = 0; aLetVariable < 5; aLetVariable++) {
    // aLetVariable is accessible here.
  }
  // aLetVariable is not accessible here.

  const - constants.

  var - common declarative keyword. It does not have the restrictions
  that the other two keywords have. A variable declared with the var
  keyword is available from the function block it is declared in.

  Note: In js blocks doesn't have scope, only functions have. So if a
  variable is defined using var in a compount statement (eg: within if),
  it will be visible to the entire function. In ECMAScript205, let and
  const declarations allow us to create block-scoped variables.

  Equality: == performs type coercion. To avoid type coercion using ===.
  123 === '123' // false
  1 == true // true
  1 === true // false

for(var i=0; i<5; i++) {}
for(let value of array)  {}
for(let property in object) {}

switch(action) {
  case '':
    statements
    break
   default:
}

Objects: can be thought as simple collections of name-value pairs.

to create a new object:
```
var obj = new Object();
// or
var obj = {};
```

```
var obj = {
  name: 'Carrot',
  'for': 'Max',
  'details': {
    'color': 'orange',
    'size': 12
   }
 }

 // attributes access can be chained together.
 obj.details.color //orange
 obj['details']['color']
```
```
 function Person(name, age){
  this.name = name;
  this.age = age;
}

// defining an object
var you = new Person('goutham', 28);
```

### Arrays
Arrays in js are a special type of object.
var arr = new Array();

for ... in is not a recommended way to loop over an array. Because this
loops if a property is added to Array.prototype.
Another way to iterate over an array is using forEach in ECMA2015.
```
['dog', 'cat', 'hen'].forEach(function(currentValue, index, array) {
}
```

js uses functions as classes.
