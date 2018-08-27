import React, { Component } from 'react';

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
    fetch(this.props.readerApi + '/chapter/' + this.props.storyId + '/' + this.props.chapterId, {
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

export default Chapter;