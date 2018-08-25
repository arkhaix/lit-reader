import React, { Component } from 'react';
import Chapter from './Chapter';
import './index.css';

class Story extends Component {

  constructor(props) {
    super(props);
    this.state = {
      Id: null,
      Url: "http://example.com",
      Author: "Nobody",
      Title: "Select a story",
      NumChapters: 0,
    };
  }

  render() {
    return (
      <div className="Story">
        <header className="Story-header">
          <h1 className="Story-title">{this.state.Title}</h1>
          <h3 className="Story-author">by {this.state.Author}</h3>
        </header>
        <button onClick={() => this.fetchStory('wanderinginn.com')}>Fetch Story</button>
        <div className="Story-chapter"><Chapter storyId={this.state.Id} chapterId={0}/></div>
      </div>
    );
  }

  fetchStory(url) {
    fetch('http://localhost:8081/story', {
      method: 'post',
      headers: {
        'Accept': 'application/json, text/plain, */*',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({url: url})
    }).then(res => res.json())
      .then(res => this.updateStateFromStory(res));
  }

  updateStateFromStory(story) {
    this.setState({
      Id: story.Id,
      Url: story.Url,
      Author: story.Author,
      Title: story.Title,
      NumChapters: story.NumChapters,
    });
  }
}

export default Story;