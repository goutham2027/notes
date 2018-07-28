## TODO
- compile es6 to es2015 using babel - Done
- setup linter - Done
- write tests to es6 - Done
- require vs import

## JavaScript Fundamentals for ES6
PluralSight course - Scott Allen, Joe Eames

1) Variables and Parameters
`let` will be the new `var`
2) Classes
3) Functional JS
  arrow functions
  iterators
  generators
  list comprehension
4) Built-in Objects
  Set, map collections
5) Asynchronous JS
  Promise
6) ES6 Objects
  Metaprogramming capability
7) Modules in JS
8) Tools

ES6 Compatibility table.

Github Course Repository: joeeames/ES6FunamentalsCourseFiles

### 1. Variables and Parameters
`let` - To define variables.

`var` has limitations. There are only 2 scopes defined using `var`. Global scope and functional scope.
`let` will give block scope.

`const` - To declare constant/read only variables. Block Scope.

#### Destructuring assignment like python/ruby.
```
let x = 2;
let y = 3;
[x, y] = [y, x];

let [x, y] = [3, 2];

let [, x, y] = [1, 3, 2]; # empty commans means skip
let [, x, y, z] = [1, 3, 2]; # z is undefined
```

Can also destructure objects
```
person = {firstName: 'a', lastName: 'b'}
let { firstName: first, lastName: last} = person
console.log(first) is 'a'
console.log(last) is 'b'
```
if variable name and attribute name is same
```
{ firstName, lastName}
```

For functions
```
let dowork = function(url, {data, cache}) {
}

doWork('api/test', {data: 'test', cache: false}
```

#### Default Parameter Values
```
function(name="abc") {
  return name;
}
```

#### Rest Parameters
variable parameters

```
sum(...numbers) # last parameter in the function

let sum = function(...numbers) {
    let result = 0;
    numbers.forEach(function(n) {
        result += n;
    })

    console.log(result)
};

sum(1, 21, 3) # 25
```

In es5 this is achieved using `arguments`.

#### Spread parameters
```
Same as rest parameters but this will assign to variables

sum = function(x, y, z) {
}

sum(...[1,2,3])
x = 1
y = 2
z = 3
```

We can also use to build arrays
```
var a = [4, 5, 6]
var b = [1, 2, 3, ...a, 7, 8, 9]
```

#### Template Literals
Variable interpolation
Can avoid string concatenation
In between backticks (`)

```
`http://apiserver/${category}/${id}`;
```

### 2. Classes

class vs prototype

#### Prototype
```
// constructor
var Employee = function() {
};

Employee.prototype = {
  doWork: function() {
    return "complete!";
  }
};

var e new Employee();
e.doWork();
```

#### Class
```
class Employee {
  doWork() {
    return "complete!";
  }
}

var e = new Employee();
e.doWork()
```

#### Getters and Setters
```
get name() {

}

set name(newValue) {
}
```

#### Inheritance
Use Inheritance when there is "is-a" relationship.

```
class Employee extends Person {
  // employee is-a person
}
```

#### Super
```
class Employee extends Person {
  constructor(name, title) {
    super(name);
    this._title = title;
  }
}
```

#### Overriding method behavior
```
class Person {
  doWork = {}
}

class Employee extends Person {
  doWork() {
    super()
    ...
  }
}
```

### 3. Functional Programming
#### Arrow functions
```
let add = (x, y) => x + y;

let three = () => 3;

// multiline
let add = (x, y) => {
  let temp = x + y;
  return temp
}

var numbers = [1,2,3,4];
var sum = 0;
numbers.forEach(n => sum += n);

var doubled = numbers.map(n => n * 2);

// arrow functions and async
```

Arrow functions will always capture the `this` value of the context that
they are inside. It lexically binds to `this`.

#### Iterables and Iterators
Iterable `[1, 2, 3, 4]`

Retrieve an iterator on Iterable.
Iterator has next() which returns `{value: <>, done: <bool>}`
```
eg: [1, 2, 3, 4]
next()
{value: 1, done: false}
{value: 2, done: false}
{value: 3, done: false}
{value: 4, done: false}
{value: undefined, done: true}
```

#### for of
`for of` uses iterator
```
let numbers = [1, 2, 3, 4];
for (let i of numbers) {
    console.log(i);
}
```

Building an Iterable

#### Generators
A generator function is a function which generates an iterator.

```
let numbers = function*() {
  yield 1;
  yield 2;
}
```

```
// implementation of range generator
let range = function*(start, end) {
  let current = start;
  while(current <= end) {
    let delta = yield current;
    current += delta || 1;
  }
}
```

#### Comprehensions
```
var numbers = [for (n of [1, 2, 3, 4]) n * n]
console.log(numbers)

yield* array // yield each element in the array instead of entire
array

```
Use Generator Comprehensions whenever possible.

### 4. Built-in Objects

#### Numbers
```
// ES5
var a = 10;
var hexd = 0xa;
var octal = 071;

// ES6
var octal = 0o71;
var bin = 0b1101;
var octNum = Number("0o71");
Number.parseInt("3") // 3
Number.parseFloat("3.5") // 3.5
```

### Math
New Math functions are introduced in ES6.

### Array
```
var match = [1, 5, 10]
match.find(5)
match.find(item => item > 8);

match.findIndex(item => item > 3);

match.fill('a');
match.fill('a', 3); // starting index
match.fill('a', 2, 3); // starting, ending index

match.copyWithin(2, 0) // [1,2,1,2]
match.copyWithin(2, 0, 1) //[1,2,1,4]
```

#### Set
New data structure added in ES6.

```
var set = new Set();

set.size();
set.add("some value");

var key = {}
set.add(key)
set.has(key)

var set = new Set([1, 2, 3, 4])
set.clear() # size becomes zero
set.delete(1)

set.values() # is an iterator
```

#### Map
```
var map = new  Map();
map.size

map.set("name", "Foo")
map.get("name") # Foo

var map = new Map([['name', 'john'],['age', 15]])
map.has('name')
map.clear()

for ([k, v] opf map) {}

map.entries() # [k, v]

map.values() # enumerates values
```

#### WeakMap and WeakSet
No size, entries, values and forEach on WeakSet
No size, entries, keys, values and forEach on WeakMap
```
var set = new WeakSet();
set.has('')
set.add(item)
set.delete(item)
set.clear()
```

```
var map = new WeakMap()
```


### 5. Asynchronous Develiopment in ES6

#### Promise Primer
Callbacks

```
function getCompanyFromOrderId(orderId) {
  try {
    getOrder(orderId, function(order) {
      getUser(order.userId, function(user) {
        getCompany(user.companyId, function(company) {
          //
        });
      });
    });
    }
    catch(ex) {}
}
```

Each callback has not only process the data by calling the appropriate
next function, but it also has to pass in an appropriate callback into
the next function, so each callback is not only handling processing, but
also flow control, and when we add exception handling it gets even
worse. When we add Exception Handling this may go even worse.

Each callback goes onto the call stack separately from the caller.
Because of that, any exceptions they throw aren't caught by the original
initiator of the async request.

Promise

Promises have long been recognized as a solution to what is often known
as callback hell.

A promise is an object which represents a handle to listen to the
results to the async operation whether it succeeds or fails. In other
words, the promise promises to alert when the async operation is done
and give the results of theat operation.

One of the main benefits of promises are that they are composable. Two
promises can be chained together so that one happens after the other. Or
we can wait for them both to run, and then run a new operation when both
are complete.

A promise is made up of two parts. The first part is the control of the
promise. In many libraries this is called a Deferred, and it is a
separate object. In other implementations, this is simply a callback
itself. This gives the creator of the promise the ability to mark the
promise as succeeded or failed.

The second part is the promise itself. This object can be passed around
and enables interested parties to register listeners who can take
actions when the asynchronous operation completes. That way, flow
control is no longer the responsibility of the same function which will
handle the result of the operation, and it can focus solely on doing its
job and others can worry about their own responsibilities.

Promise will also notify if the async process has a failure, so error
handling is no longer swallowed up into the great Asynchronous abyss
(deep) of javascript call stacks.

A promise exists in on of three states: Pending, Fulfilled and Rejected.

```
function getCompanyFromOrderId(orderId) {
  getOrder(orderId).then(function(order) {
    return getUser(order.userId);
  }).then(function(user) {
    return getCompany(user.companyId);
  }).then(undefined, function(error) {})
}
```

#### Promise Basics
