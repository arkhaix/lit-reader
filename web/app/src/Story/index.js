import React, { Component } from 'react';

import Button from '@material-ui/core/Button';
import Paper from '@material-ui/core/Paper';
import TextField from '@material-ui/core/TextField';

import Chapter from './Chapter';

import './index.css';

class Story extends Component {

  constructor(props) {
    super(props);
    this.state = {
      Id: null,
      Url: "http://example.com",
      Author: "Author",
      Title: "Select a story",
      NumChapters: 0,

      DesiredUrl: "",
    };
  }

  render() {
    return (
      <div className="Story">
        <Paper elevation={5}>
          <header className="Story-header">
              <h1 className="Story-title">{this.state.Title}</h1>
              <h3 className="Story-author">by {this.state.Author}</h3>
          </header>
        </Paper>
        <div className="debug">
          <form onSubmit={this.handleSubmit}>
            <TextField
              label="Story url"
              placeholder="wanderinginn.com"
              margin="normal"
              name="DesiredUrl"
              value={this.state.DesiredUrl}
              onChange={this.handleChange}
            />
            <div className="debug-fetch">
            <Button variant="outlined" color="inherit" size="small" type="submit" value="submit">
              Fetch Story
            </Button>
            </div>
          </form>
        </div>
        <Chapter readerApi={this.props.readerApi} storyId={this.state.Id} chapterId={0}/>
      </div>
    );
  }

  handleChange = (event) => {
    const { target: { name, value } } = event;
    this.setState({ [name]: value });
  }

  handleSubmit = (event) => {
    event.preventDefault();
    this.fetchStory(this.state.DesiredUrl);
  }

  fetchStory(url) {
    fetch(this.props.readerApi + '/story', {
      method: 'post',
      headers: {
        'Accept': 'application/json, text/plain, */*',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({url: url})
    }).then(res => res.json())
      .then(res => this.updateStateFromStory(res));
  }

  updateStateFromStory(storyResponse) {
    if (storyResponse.Status.Code !== 200) {
      console.log("Bad story");
      return;
    }
    console.log('setting new story state:');
    console.log(storyResponse);
    this.setState({
      Id: storyResponse.Id,
      Url: storyResponse.Url,
      Author: storyResponse.Author,
      Title: storyResponse.Title,
      NumChapters: storyResponse.NumChapters,
    });
  }
}

export default Story;