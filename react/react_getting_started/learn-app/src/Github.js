import React, { Component } from 'react';
import axios from 'axios';

const alertStyle = {
  background: "#DA3849",
  padding: "10px",
  color: "#fff",
  width: "300px"
};

const infoStyle = {
  background: "#25A2B7",
  padding: "10px",
  color: "#fff",
  width: "300px"
};

class Notification extends React.Component {

  constructor(props) {
    super(props);
    if(props.type === 'alert') {
      this.style = alertStyle;
    }
    else if(props.type === 'info') {
      this.style = infoStyle;
    }
  };

  render() {
    return (
      <div>
      <p style={this.style}> {this.props.message} </p>
      </div>
    );
  }
}

const Card = (props) => {
  return (
    <div style={{margin: '1em'}}>
      <img src={props.avatar_url} height="80px" width="80px" />
      <div style={{display: 'inline-block', marginLeft: 10}}>
        <div> {props.name}  </div>
        <div> {props.company} </div>
      </div>
    </div>
  );
};



const CardList = (props) => {
  return (
    <div>
      {props.cards.map(card => <Card {...card} key={card.id} />)}
    </div>
  );
};

class Form extends React.Component {
  state = {userName: ''}
  handleSubmit = (event) => {
    event.preventDefault();
    axios.get(`https://api.github.com/users/${this.state.userName}`)
    .then(resp => {
      this.props.onSubmit(resp.data);
      this.setState({userName: ''});
    });
  };

  render() {
    return (
      <form onSubmit={this.handleSubmit}>
        <input type="text"
          value={this.state.userName}
          onChange={(event) => this.setState({userName: event.target.value})}
          /*
           *ref={(input) => this.userNameInput = input}
           */
          placeholder="Github username" />
        <button type="submit">Add card </button>
      </form>
    );
  }
}

class GithubApp extends React.Component {
  state = {
    cards: [],
    is_duplicate: false,

  };

  addNewCard = (cardInfo) => {
    if (! this.userCardExists(cardInfo['login'])) {
      this.setState({is_duplicate: false})
      this.setState(prevState => ({
        cards: prevState.cards.concat(cardInfo)
      }));
    }
    else {
      this.setState({is_duplicate: true})
    }
  };

  userCardExists = (userName) => {
    const cards = this.state.cards;
    for(let card of cards) {
      if(card['login'] === userName) {
        return true;
      }
    }
  };

  render() {
    return (
      <div>
        <Form onSubmit={this.addNewCard}/>
        <CardList cards={this.state.cards} />
        {this.state.is_duplicate &&
          <Notification message="Duplicate Entry" type="alert" />
        }
      </div>
    )
  }
}

export default GithubApp;
