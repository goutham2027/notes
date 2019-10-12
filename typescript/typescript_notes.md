Frontend masters course - https://frontendmasters.com/courses/typescript/introduction/
By Mike North

### What's TypeScript?
- A typed superset of JavaScript
- Developed by Microsoft
- Compiles to Javascript for either browsers or Node
- 3 parts: Language, Language Service and Compiler

### JS Types
Primitive Types -  Primitive values aren't objects and have no methods
* Primitive types are immutable.

6 primitive types in JS
* null
* undefined
* boolean
* number
* string
* symbol (ES2015)

Everything else extends from Object

### Type conversion
3 different ways in JS
* var
- var declarations are hoisted - it's as if they are all decleared at the top of the global or function scope in which they are defined.
eg:
```javascript
function foo() {
    console.log(x);
    var x = 15;
}
```
Browser sees as following
```javascript
function foo() {
    var x = 15;
    console.log(x);
}
```
- `var` declarations are not block scoped. Belong to the function or global scope they're defined in.

* let
- `let` declarations are not hoisted
- `let` declarations are block scoped.
- polyfills and transpilers check for, and enforce this at build time.

* const
- `const` declarations are not hoisted
- `const` declarations are block scoped
- must be initialized at the time of declaration
- re-assignment is not allowed
- constant variable does not mean "immutable value"
- use Object.freeze for mutable values

### Type conversion
- In dynamically typed languages, things are converted as needed
- The `+` operator, when used with strings, converts all other operands to strings
eg: `'37' + 7 // 377`
`'37' - 7 // 30`
- The unary `+` operator converts the operand to a number
eg: `(+ 37) // 37`


This type conversion is confusing. This sets a clear need of type system.

## Types
### Implicit Typing
- The typescript compiler can make good guesses at types, just through assignment
- After assigning a value to a variable, we are not allowed to change the type.
```javascript
let age = 35;
age = "35"; // not acceptable in typescript
```

### Explicit typing
Type annotation
```typescript
let teacherAge: number = 34;
```

Casting
```javascript
let input = document.querySelector('input#name_field') as HTMLInputElement;
```
another way to do is, but it doesn't mix well with JSX
```javascript
let input = <HTMLInputElement>document.querySelector('input#name_field');
```

Function Parameters & Return
```javascript
function login(username: string, password: string) : User {

}

// arrow functions
const login = (username: string, password: string) : User => {

}
```
### Object Shapes

Type systems & Type equivalence
```javascript
function validateInputField(input: HTMLInputElement) {

}

validateInputField(x); // can we regard x as an HTMLInputElement
```

Nominal Type Systems answers this question based on whether x is an instance of a class/type named HTMLInputElement.

Structural Type Systems only care about the shape of an object. This is how typescript works.

Shape of an object - names of properties and types of their values.

```javascript
let myCar: {make: string, model: string, year: number};

myCar = {
    make: 'Honda',
    model: 'Accord'
    year: 1992,
    color: {r: 255, g: 0, b: 0}
};
```
Excess properties are fine.


Return type `void` means the function retuns `undefined`

### Interfaces
Allow us to define a structure and refer to it by name
```typescript
interface Car {
    make: string;
    model: string;
    year: number;
}
```
- Interfaces only describe structure, they have no implementation
- They don't compile to any JS code
- DRY type definition allows for easy refactoring later

### Any & Never Types

- `any` type allows for a value of any kind.
- Back to regular JS if we use any..
```typescript
let age = 34;
let myAge = age as any;
myAge = '35'; // myAge can take any value
```

- `never` type is compatible with nothing.
```javascript
let y: never = "hello"; // not possible
```
- Use case?

Interfaces values can be set to a literal and when an interface is used it looks for exact literal
eg:
```typescript
interface ConfirmedUser {
    email: string;
    password: string;
    isActive: true;
}
```

## Classes
- Javascript uses prototypes for inheritance. Prototypical inheritances are more flexible.

```javascript
class Person {
    constructor(name) {
        this.name = name;
    }

    toString() {
        return `Person: ${this.name}`;
    }

    static createRandom() {
        return new Person(`${Math.random()}`);
    }
}

let stranger = Person.createRandom();
console.log(stranger.name)
```

### Inheritnace
- Using `extends` keyword.
- The `super` keyword can be used to call methods on the parent class.

```typescript
class Person {
    constructor(name) {
        this.name = name;
    }
    toJSON() {
        return {
            name: this.name;
        };
    }
}

class Employee extends Person {
    constructor(id, name) {
        super(name);
        this._em = id;
    }
}
```
### Species
- There is a special property on classes called Symbol.species that's used when building derived objects.

Eg: An array that doesn't print any private info via `console.log`, but `map` and `filter` return regular arrays that don't have this restriction.

```javascript
class MyArray extends Array {
    toString() {
        return '[PRIVATE]';
    }

    static get [Symbol.species]() {
        return Array;
    }
}


let a = new MyArray(1, 2, 3);
console.log(`${a}`); // toString is called

let filtered = a.filter((y)) => y <= 2);
console.log(`${filtered`}); // returns original array
```

Example on static functions
```typescript
class Example {
    static abc() {
        return "static function"
    }

    def() {
        return "regular function"
    }
}

alert(Example.abc())

let e = new Example()
alert(e.def())

```

```javascript

function add(a, b) {
    return a + b;
}

// using arrow function
// it will carry calling context
const add = (a, b) => a+b;

// to carry context
function add(a, b) {
    return a + b;
}.bind(this);
```

### Mixins
Abstract classes or templates for classes

### Enums
ordered set

```typescript
enum AccType {
    Checking,
    Savings,
    MoneyMarket
}

type Acct = [number, AccType]

let account: Acct = [
    9142.14, AcctType.Checking
]
```

Note: When initializing class properties with empty arrays, provide a type.

```typescript
class ShoppingCart {
    items = [];
    constructor() {
        this.items.push(5)
    }
}

// we should be doing something like the following
class ShoppingCart {
    items: number[] = [];
    constructor() {
        this.items.push(5)
    }
}
```

### Tuples
- Arrays of fixed length
- Typically represent values that are related in some way.
- Consumers need to know about order

## Object Literals

### Rest and Spread Properties
```javascript
// Rest - destructuring
let data = { x: 34, y: 21, z: 0.1};
let { x, ...others } = data;
console.log(others); //  { y: 21, z: 0.1}

// Spread
let values = { ...others, a: 99, b: 77};
console.log(values) // { y: 21, z: 0.1, a: 99, b:77 }
```

### Getters (accessor) & Setters (mutator)
```typescript
let name = {
    first: 'Michael',
    last: 'North',
    get full() {
        return `${this.first} ${this.last}`
    }
    set full(newVal) {
        let [a, b] = newVal.split(/\s+\g);
        this.first = a;
        this.last = b;
    }
}

console.log(name.first, name.last);
name.full = "Michael North";
console.log(name.first, name.first);
```

### Function Types
- Functions have a type just like any other value.
```typescript
let login: (username: string, password: string) => User; // variable declaration. A function type.

login = (username, password) => { return new User(); }; // variable assignment. A function value

// interfaces can be used to define functions
interface ClickListener {
    (this: Window, e: MouseEvent): void
}

const myListener: ClickListener =
    e => {
        console.log('mouse clicked', e);
    }

addEventListener('click', myListener);
```

### Functions: parameters
```typescript
let criticalURL: URL | undefined;
// optional parameters - ?
function createTwitterPost(body: string, username: string, imageUrl ?: URL) {
    criticalURL = imageURL
}

// default values
function createTwitterPost(body: string,
                         username: string = 'MichaelNorth)


// rest parameters
// A boundless group of optional parameters
function orderSandwich(bread: string, name: string, ...toppings: string[])
```

## Generics
- Generics allow us to reuse code across many types, interfaces and functions.
- We still get compile-time safety
```typescript
// <T> is Type Parameter
// (x: T) - T is determined by argument type
function gimmieFive<T>(x: T): T[] {
    return [x, x, x, x, x];
}

let threes: number[] = gimmeFive(3);
let eggs: string[] = gimmeFive('egg');

// Arrays can be expressed this way too
let cards = Array<[Suit, CardNumber]>(52);

// promises
let data: Promise<Response> = fetch('http://example.com');

// react components
interface MyProps {title: string}
interface MyState {isClicked: boolean}

class MyComponent extends Component<MyProps, MyState> {}
```

### Access modifier Keywords
public - Anyone can access

protected - self and subclasses can access

private - self can access

static

readonly

```typescript
class Account {
    protected email: string;
    private password: string;
    public accountId: number;
}
```

### Function Overloading
- Typescript allows us to have more than one function "head", although we are still limited to a single implementation.

```typescript
function add(x: number, y: number): number;
function add(x: string, y: string, radix: number): number;

function add(x: number|string, y: number|string, radix: number=10): number {
    return parseInt(`${x}`, radix) + parseInt(`${y}`, radix);
}

```