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
    console.log('render');
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
    console.log('in updateChapter with props:')
    console.log(this.props)
    if (this.props.storyId !== this.state.StoryId || this.props.chapterId !== this.state.ChapterId) {
      this.fetchChapter()
    }
  }

  fetchChapter() {
    console.log('in fetchChapter');
    if(!this.props.storyId || this.props.chapterId === null || this.props.chapterId < 0) {
      console.log('bad props');
      return
    }
    console.log('props ok. fetching');
    fetch('/chapter/' + this.props.storyId + '/' + this.props.chapterId, {
      method: 'get',
      headers: {
        'Accept': 'application/json, text/plain, */*',
        'Content-Type': 'application/json'
      }
    }).then(res => res.json())
      .then(res => this.updateStateFromChapter(res));
  }

  updateStateFromChapter(chapter) {
    console.log('updateStateFromChapter:');
    console.log(chapter);
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