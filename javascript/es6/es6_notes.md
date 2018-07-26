## TODO
- compile es6 to es2015 using babel - Done
- setup linter - Done
- write tests to es6

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