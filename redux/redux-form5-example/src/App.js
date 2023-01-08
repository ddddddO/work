import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';

import showResults from './showResults';
import MyForm from './MyForm';
import MyFamilyForm from './MyFamilyForm';
import { Container, Row, Col } from 'react-bootstrap';

class App extends Component {
  constructor(props) {
    super(props)
    this.nextPage = this.nextPage.bind(this)
    this.state = {
      page: 1,
    }
  }

  nextPage() {
    this.setState({ page: this.state.page + 1 })
  }

  render() {
    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <h1 className="App-title">Example redux-form5✖</h1>
        </header>
        <Container style={{padding: 15}}>
          <Row>
            <Col sm={10}>
              <h3>フォーム画面</h3>
              {
                this.state.page === 1 &&
                <MyForm onSubmit={this.nextPage}/>
              }
              {
                this.state.page === 2 &&
                <MyFamilyForm onSubmit={showResults}/>
              }
            </Col>
          </Row>
        </Container>
      </div>
    );
  }
}

export default App;
