Frontend masters course: https://frontendmasters.com/courses/react-state/

Libraries: Flux, Redux, Mobx, Redux-thunk, Redux-Saga

Managing UI state is not a solved problem.

## Understanding State
The main job of React is to take your application state and turn it into
DOM nodes.

### What state is not?
- If it is passed from a parent via props - then it is not a state
- Does it remain unchanged over time? - then it is not a state
- Can you compute it based on any other state or props - then it is not
  a state.

React is all about one-way data flow fown the component hierarchy. It
may not be immediately clear which component shown own what state.

State is created in the component and stays in the component. It can be
pased to a children as its props.

### Types of state
* Model data: Coming from the backend. Nouns in the application.

* View/UI state: are those nouns sorted in ascending or descending
  order?

* Session state: is the user even logged in?

* Communication: are we in the process of fetching the nouns from the
  server?

* Location: Where are we in the app? which nouns are we looking at?



### Long lasting state vs Ephemeral state

Long lasting state: This is likely the data in app

Ephemeral state input field value will be wiped when user hits enter.


```javascript
this.setState({count: this.state.count + 1})
this.setState({count: this.state.count + 1})
this.setState({count: this.state.count + 1})

console.log(this.state.count);

// output is 0
this.setState() is asyncronous.

React is trying to avoid unnecessary re-renders.
```

```javascript
increment() {
this.setState({count: this.state.count + 1})
this.setState({count: this.state.count + 1})
this.setState({count: this.state.count + 1})
}

render() {...}

// What will the count be after the user clicks the increment button?
// output: 1
// We are queuing up state changes. React will batch them up, figure of
// the result and then effeciently make that change.
```

```javascript
increment() {
  this.setState((state) => { return { count: state.count + 1}});
  this.setState((state) => { return { count: state.count + 1}});
  this.setState((state) => { return { count: state.count + 1}});
}
```

`this.setState() takes a callback as a second argument.

```javascript
increment() {
  this.setState(
    {count: this.state.count + 1},
    () => { console.log(this.state);}
  )
}
```

## State patterns and anti patterns
State should be considered private data. If you are not using state in
render.
Don't use `this.state` for derivations of props.

### Shared component state problems
- Identify every component that renders something based on that state.
- Find a common owner component - a single component above all the
  components that need the state in the hierarchy.
- If you can't find a component where it makes sense to own the state,
  create a new component simply for holding the state and add it
  somewhere in the hierarchy above the common owner component.

The lowest common ancestor.

### Container pattern
- Draw a line between state and presentation.
- Pesentational components receive props and render UI.
- They also receive actions and pass them back to the container.
- Presentational components have a render() method or they are stateless
  functional components.

### Higher order components
Higher order components take a component as an argument return a wrapped
component. Take a presentational component and return a wrapped in
container component. A container factory.

https://frontendmasters.com/courses/react-state/providing-a-display-name-to-hocs/



## Reading Material
Thinking in React - https://reactjs.org/docs/thinking-in-react.html
