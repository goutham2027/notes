### Getting started with Redux.
https://egghead.io/lessons/javascript-redux-the-single-immutable-state-tree


3 principles of redux

1) state is maintained as a single javascript object - state tree
2) state tree is read only. To change state dispatch an action. Action
is a plain javascript object.
3) to describe state mutuations we need to write a function which takes
the current state of the application and the action to dispatch and
returns the next state of the app. This function is pure. This function
is called the Reducer.

Notes to principle 3:
pure and impure functions

pure functions - return values are dependent on the input. predictable
passed values are not changed. In redux we write pure fucntions to the
most part.

eg:
```fucntion square(x) {
  return x + x;
}```

impure functions - network calls, db calls, dom modifications change
object.

Redux: The Reducer Function
The state mutuations in the app need to be described as pure function.

Takes previous state and the action to dispatch and returns the new
state of the application.

Inside any Redux application there is one particular function that takes
the state of the whole application and the action being dispatched and
returns the next state of the whole application. It is important that it
doesn't modify the state given to it, it has to be pure so it has to
return a new object.

#### Writing a Counter Reducer with Tests

