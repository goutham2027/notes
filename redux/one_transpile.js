'use strict';

console.log("hello");

var counter = function counter() {
  var state = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : 0;
  var action = arguments[1];

  switch (action.type) {
    case 'INCREMENT':
      return state + 1;
    case 'DECREMENT':
      return state - 1;
    default:
      return state;
  }
};

// const { createStore } = Redux;
var createStore = function createStore(reducer) {
  var state = void 0;
  var listeners = [];

  var getState = function getState() {
    return state;
  };

  var dispatch = function dispatch(action) {
    state = reducer(state, action);
    listeners.forEach(function (listener) {
      return listener();
    });
  };

  var subscribe = function subscribe(listener) {
    listeners.push(listener);
    return function () {
      listeners = listeners.filter(function (l) {
        return l !== listener;
      });
    };
  };

  dispatch({});
  return { getState: getState, dispatch: dispatch, subscribe: subscribe };
};

var store = createStore(counter);

// get current state
console.log(store.getState());

var render = function render() {
  document.body.innerText = store.getState();
};

store.subscribe(render);
render();

document.addEventListener('click', function () {
  store.dispatch({ type: 'INCREMENT' });
});
