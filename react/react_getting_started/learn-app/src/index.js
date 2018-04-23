import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
//import App from './App';
import Button from './App';
import GithubApp from './Github';

import registerServiceWorker from './registerServiceWorker';

ReactDOM.render(<GithubApp />, document.getElementById('root'));
/*
 *ReactDOM.render(<Button/>, document.getElementById('root'));
 */

/*
 *class Hello extends React.Component {
 *  constructor(props) {
 *    super(props);
 *    this.state = {hoverCount: 0, name: '', inputValue: ''};
 *  };
 *
 *  handleHover = () =>  {
 *    this.setState({hoverCount: this.state.hoverCount + 1})
 *  };
 *
 *  handleKeyDown = (evt) => {
 *    this.setState({inputValue: evt.target.value})
 *    console.log("hello");
 *  }
 *
 *  render() {
 */
    /*
     *return (<p onMouseOver={this.handleHover}>Hello {this.props.name} {this.state.hoverCount}</p>)
     */
/*
 *    return (<div>
 *            <input type="text" onKeyDown={this.handleKeyDown} />
 *            <p> {this.state.inputValue}</p>
 *            </div>
 *           )
 *  }
 *}
 */

/*
 *ReactDOM.render(<Hello name="Goutham"/>, document.getElementById('root'));
 */


registerServiceWorker();
