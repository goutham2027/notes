```
npm install create-react-app -g

create-react-app react-app

cd react-app

npm start
```


// One way of creating a component. Stateless way of defining a
componenet. We can not have a state.
/*
 *const App = () => <h1> Hi </h1>
 *
 *export default App
 */

// other way of creating a compoent.
class App extends React.Component {
  render() {
    /*
     *return React.createElement('h1', null, 'Hello Eggheads')
     */
    return (
      <h1> Hello World </h1>
    )
  }
}

* render method is allowed to return a single node. Use a top level
  element.
        <div>
        <h1> Hello World </h1>
        <b> bold </b>
      </div>

The following will not work, as it translates to
React.createElement('h1', null, 'Hello World') React.createElement('b',
null, 'bold')
      return (
        <h1> Hello World </h1> <b> Bold </b>
      )


* set properties to react components

* setting what data types props will have
App.propTypes = {
  txt: React.PropTypes.string,
  cat: React.PropTypes.number.isRequired
}

* setting default values for properties
App.defaultProps = {
  txt: "this is the default txt"
}

* state:
unlike props, which are a collection of values that are meant to be
passed into component as static values not meant to be changed by our
component, state represents a collection of values that is meant to be
managed as well as updated by our component.

```
class App extends React.Component {
  constructor() {
    super();
    this.state = {
      'txt': 'this is the state txt'
    }
  }

  update(e) {
    this.setState({txt: e.target.value})
  }
  render() {
      let txt = this.state.txt
      let cat = this.props.cat
    return (
      <div>
      <input type="text" onChange={this.update.bind(this)}/>
        <h1>{ txt }</h1>
        <h1>{ cat }</h1>
      </div>
    )
  }
}

App.propTypes = {
  txt: React.PropTypes.string,
  cat: React.PropTypes.number.isRequired
}

App.defaultProps = {
  txt: "this is the default txt"
}

export default App

```
using other components
```
import React from 'react';

// other way of creating a compoent.
class App extends React.Component {
  constructor() {
    super();
    this.state = {
      'txt': 'this is the state txt'
    }
  }

  update(e) {
    this.setState({txt: e.target.value})
  }
  render() {
      let txt = this.state.txt
    return (
      <div>
        <h1>{ txt }</h1>
        <Widget update={this.update.bind(this)} />
      </div>
    )
  }
}

const Widget = (props) =>
  <input type="text" onChange={props.update}/>

/*
 *App.propTypes = {
 *  txt: React.PropTypes.string,
 *  cat: React.PropTypes.number.isRequired
 *}
 *
 *App.defaultProps = {
 *  txt: "this is the default txt"
 *}
 */

export default App
```

### accessing child properties
```
class App extends React.Component {
  render() {
    return <Button>I <Heart /> React </Button>
  }
}

const Button = (props) => <button>{props.children}</button>

class Heart extends React.Component {
  render() {
    return <span>&hearts;</span>
  }
}
```

### custom validation
```
class App extends React.Component {
  render() {
    return <Title text="goutham"/>
  }
}

const Title = (props) => <h1>Title: {props.text}</h1>

Title.propTypes = {
  text(props, propName, component) {
    if(!(propName in props)) {
      return new Error(`missing ${propName}`)
    }
    if(props[propName].length < 6) {
      return new Error(`${propName} was too short`)
    }
  }
}
```

### React synthetic event system
to support all browsers

```
class App extends React.Component {
  constructor() {
    super();
    this.state = {currentEvent: '---'}
    this.update = this.update.bind(this)
  }

  update(e) {
    this.setState({currentEvent: e.type})
  }
  render() {
    return (
      <div>
        <textarea
        onKeyPress={this.update}
        onCopy={this.update}
        onCut={this.update}
        onPaste={this.update}
        onFocus={this.update}
        onBlur={this.update}
        onDoubleClick={this.update}
        onTouchStart={this.update}
        onTouchMove={this.update}
        onTouchEnd={this.update}
        cols="30"
        rows="10" />
        <h1>{this.state.currentEvent}</h1>
      </div>
    )
  }
}
```

### Use refs to access components
```
class App extends React.Component {
  constructor() {
    super();
    this.state = {a: '', b: ''}
  }

  update() {
    this.setState({
      a: this.refs.a.value,
      b: this.refs.b.value
    })
  }
  render() {
    return (
      <div>
      <input
      ref="a"
      type="text"
      onChange={this.update.bind(this)}
      /> {this.state.a}
      <hr />
      <input
      ref="b"
      type="text"
      onChange={this.update.bind(this)}
      /> {this.state.b}
      </div>
    )
  }
}
```

### React component lifecycle methods
When component is added to DOM it is called mounting and when it is
removed it is called un-mounting.

componentWillMount() - Fired right before it gets mounted into the dom.
Runs only once.
componentDidMount()
componentWillUnmout()
example
```js
class App extends React.Component {
  constructor() {
    super();
    this.state = {val: 0}
    this.update = this.update.bind(this)
  }

  update() {
    this.setState({val: this.state.val + 1})
  }

  componentWillMount() {
    console.log('componentWillMount')
  }


  render() {
    console.log('render');
    return <button onClick={this.update}>{this.state.val}</button>
  }

  componentDidMount() {
    console.log('componentDidMount')
  }

  componentWillUnmount() {
    console.log('componentWillUnmount')
  }
}

class Wrapper extends React.Component {
  mount() {
    ReactDOM.render(<App />, document.getElementById('a'))
  }

  unmount() {
    ReactDOM.unmountComponentAtNode( document.getElementById('a'))
  }
  render() {
    return (
      <div>
        <button onClick={this.mount.bind(this)}>Mount</button>
        <button onClick={this.unmount.bind(this)}>UnMount</button>
      </div>
    )
  }
}
```

### Manage react component state with lifecycle methods
clear the values in componentWillUnmount() that are set in componentDidMount()

### Control react component updates when new props are received
componentWillReceiveProps
shouldComponentUpdate
componentDidUpdate

```
import React from 'react';
import ReactDOM from 'react-dom';

// other way of creating a compoent.
class App extends React.Component {
  constructor() {
    super();
    this.state = {increasing: false}
  }
  update() {
    ReactDOM.render(<App val={this.props.val+1}/>, document.getElementById('root'))
  }

  compoentWillReceiveProps(nextProps) {
    this.setState({increasing: nextProps.val > this.props.val })
  }

  shouldComponentUpdate(nextProps, nextState) {
    return nextProps.val % 5 === 0;
  }

  componentDidUpdate(prevProps, prevState) {
    console.log(`prevProps: ${prevProps.val}`)
  }
  render() {
    return (
      <button onClick={this.update.bind(this)}>
      {this.props.val}
      </button>
    )
  }
}


App.defaultProps = {val: 0}
```

### React dynamically generated components
```
class App extends React.Component {
  constructor() {
    super();
    this.state = {items: []}
  }
  componentWillMount() {
    fetch('http://swapi.co/api/people/?format=json')
    .then(response => response.json())
    .then( ({results: items}) => this.setState({items}))
  }
  render() {
    let items = this.state.items
    return (
      <div>
        {items.map(item => <h4 key={item.name}>{item.name}</h4>)}
      </div>
    )
  }
}


export default App
```
#### Dynamic filter from list
```
import React from 'react';

// other way of creating a compoent.
class App extends React.Component {
  constructor() {
    super();
    this.state = {items: []}
  }
  componentWillMount() {
    fetch('http://swapi.co/api/people/?format=json')
    .then(response => response.json())
    .then( ({results: items}) => this.setState({items}))
  }
  filter(e) {
    this.setState({filter: e.target.value})
  }
  render() {
    let items = this.state.items;
    if(this.state.filter) {
      items = items.filter(item =>
                           item.name.toLowerCase()
                           .includes(this.state.filter.toLowerCase()))
    }
    return (
      <div>
        <input type="text"
        onChange={this.filter.bind(this)} />
        {items.map(item => <Person key={item.name} person={item} />)}
      </div>
    )
  }
}


const Person = (props) => <h4> {props.person.name} </h4>
export default App
```

### Compose React Component Behavior with Higher Order Components
The purpose of higher order components is to share common functionality,
information between multiple components. And the sole functionality of
higher order component is to take in a component and return a component.
```
import React from 'react';

const HOC = (InnerComponent) => class extends React.Component {
  componentWillMount() {
    console.log("Will mount")
  }

  constructor() {
    super();
    this.state = {count:0}
  }

  update() {
    this.setState({count: this.state.count +1})
  }
  render() {
    return (
      <InnerComponent
        {...this.props}
        {...this.state}
        update={this.update.bind(this)}
      />
    )
  }
}

class App extends React.Component {
  render() {
    return (
      <div>
        <Button>button</Button>
        <hr/>
        <LabelHOC>label</LabelHOC>
      </div>
    )
  }
}

const Button = HOC((props) => <button onClick={props.update}>{props.children}-{props.count}</button>)

class Label extends React.Component {
  componentWillMount() {
    console.log("Will mount")
  }
  render() {
    return (
      <label onMouseMove={this.props.update}>{this.props.children}- {this.props.count}</label>
    )
  }
}

const LabelHOC = HOC(Label)

export default App
```

