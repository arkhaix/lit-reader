import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';

class Chapter extends Component {
  constructor(props) {
    super(props);
    this.state = {
      StoryId: props.storyId || "",
      ChapterId: -1,
      Url: "http://example.com",
      Title: "Select a chapter",
      Html: "<p>Select a chapter to continue</p>",
    };
  }
  render() {
    this.updateChapter();
    return (
      <div className="Chapter">
        <div className="invisible"></div>
        <header className="Chapter-header">
          {/*Not using Chapter-title because the title is included in the html*/}
          {/*<h2 className="Chapter-title">{this.state.Title}</h2>*/}
        </header>
        <div className="Chapter-html" dangerouslySetInnerHTML={this.getChapterHtml()}></div>
      </div>
    );
  }
  getChapterHtml() {
    return {__html: this.state.Html};
  }
  updateChapter() {
    if (this.props.storyId !== this.state.StoryId || this.props.chapterId !== this.state.ChapterId) {
      this.fetchChapter()
    }
  }
  fetchChapter() {
    if(!this.props.storyId || this.props.chapterId === null || this.props.chapterId < 0) {
      return
    }
    fetch('http://localhost:8082/story/' + this.props.storyId + '/chapter/' + this.props.chapterId, {
      method: 'get',
      headers: {
        'Accept': 'application/json, text/plain, */*',
        'Content-Type': 'application/json'
      }
    }).then(res => res.json())
      .then(res => this.updateStateFromChapter(res));
  }
  updateStateFromChapter(chapter) {
    this.setState({
      StoryId: chapter.StoryId,
      ChapterId: chapter.ChapterId,
      Url: chapter.Url,
      Title: chapter.Title,
      Html: chapter.Html,
    });
  }
}

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

class App extends Component {
  render() {
    return (
      <Story />
    );
  }

}

export default App;
