import React from 'react';
import ReactDOM from 'react-dom';

// other way of creating a compoent.
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

export default Wrapper
