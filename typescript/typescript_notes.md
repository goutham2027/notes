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