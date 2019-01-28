Learning Log:

### Aug 31, 2017 - Thursday

pluralsight course

amd - asynchronous module definition

#### why typescript and not regular javascript?
to write clean js we need to follow patterns otherwise js can be messy.
js provides a dynamic type system

pros
 - variables can hold any object
 - types determined on the fly
 - implicit type coercion (eg: string to number)

cons
  - difficult to ensure proper types are passed without tests
  - not all developers use ===
  - enterprise scale apps can have 1000s of lines of code to maintain

typescript alternatives
  - pure js
  - apply js patterns
  - CoffeScript
  - Dart (Google)

what is typescript? typescriptlang.org
 - typed superset of js that compiles to plain js.
 - open source

typescript features
  - supports js code
  - provides static typing
  - encapsulation through classes and modules
  - support for constructors, properties and functions
  - define interfaces
  - arrow functions, es6 (for anonymous functions)

typescript compiler
  ts files -> tsc <ts_file.ts> -> outputs js file

syntax
  - {} for blocks
  - semi-colons end code expressions

code hierarchy
  - module/namespace
    - class
      - fields, properties, constructor, functions
      - interfaces

tooling and framework options
  - nodejs
  - editors: sublime/emacs/vi/visual studio
  - typescript playground


tsc --init

tsc --watch -out bundle.js app.ts


### Sep 1, 2017 - Friday

installed typescript vim plugin - https://github.com/leafgarland/typescript-vim

#### Core features of Typescript: typing, variables and functions

Grammar, Declarations and Annotations

Grammar: type inference
`var num = 2; // type is declared from the input`

Grammar: type annotations
`var num: number;`

`var any1;` - type could be any type


Static and Dynamic Typing

how to?
* compile minified file
* disable comments in the generated js

