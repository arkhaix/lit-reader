import React, { Component } from 'react';
import Story from './Story';
import './App.css';


/* App */
class App extends Component {
  render() {
    return (
      <Story readerApi={process.env.REACT_APP_READER_API}/>
    );
  }
}

export default App;
