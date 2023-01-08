import logo from './logo.svg';
import './App.css';

import showResults from './showResults';
import MyForm from './MyForm';
import { Container, Row, Col } from 'react-bootstrap';

function App() {
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
            <MyForm onSubmit={showResults}></MyForm>
          </Col>
        </Row>
      </Container>
    </div>
  );
}

export default App;
