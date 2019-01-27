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

## Adding Types
