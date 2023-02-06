https://www.youtube.com/watch?v=BwuLxPH8IDs&t=8541s&ab_channel=Academind

### What is TypeScript?

- A Javascript superset
- A language building up on Javascript
- Adds new features + advantages to JS
- Browser can't execute it

- TS compiled to JS
- TS features are compiled to JS workarounds, possible errors are thrown.
- Typed
- Can be caught errors during compile time instead of runtime

```javascript
function add(num1, num2) {
  return num1 + num2;
}

console.log(add("2", "3"));
// Unwanted behavior at Runtime
// Mitigation strategies: add check to add function validate
```

```javascipt

```

### Install

`npm install -g typescript`

### Compule

`tsc file_name.js`

TypeScript adds

- Types
- Next-gem JS features (compiled down for older browsers)
- Just like Babel
- Non-JS features like Interfaces or Generics
- Meta-programming features like decorators
- Rich configuration options

### VS code extensions

- ESLint
- Material Icon Theme
- Path intellisense
- Prettier

### Working with Types

#### Core types

Both JS and TS

- number - 1, 5.3, -10
- string - 'Hi', "Hi", `Hi`
- boolean - true, false

Type inference
let number1 = 5 // type is inferred
Type assignment
`let number1: number`

#### TS types

JS uses "dynamic types" resolved at runtime, TypeScript uses "static types" set during development
