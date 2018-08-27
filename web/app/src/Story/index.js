import React, { Component } from 'react';

import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';

import Button from '@material-ui/core/Button';
import Paper from '@material-ui/core/Paper';
import TextField from '@material-ui/core/TextField';

import ChevronLeftIcon from '@material-ui/icons/ChevronLeft';
import ChevronRightIcon from '@material-ui/icons/ChevronRight';
import FirstPageIcon from '@material-ui/icons/FirstPage';
import LastPageIcon from '@material-ui/icons/LastPage';

import Chapter from './Chapter';

import './index.css';
const styles = theme => ({
  button: {
    margin: theme.spacing.unit,
  },
  leftIcon: {
    marginRight: theme.spacing.unit,
  },
  rightIcon: {
    marginLeft: theme.spacing.unit,
  },
  iconSmall: {
    fontSize: 20,
  },
});

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
      CurrentChapter: -1,
    };
  }

  render() {
    const { classes } = this.props;
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
        <Chapter readerApi={this.props.readerApi} storyId={this.state.Id} chapterId={this.state.CurrentChapter}/>
        <Button variant="contained" disabled={!this.firstButtonEnabled()} size="small" className={classes.button}>
          <FirstPageIcon />
          First
        </Button>
        <Button variant="contained" disabled={!this.prevButtonEnabled()} size="small" className={classes.button}>
          <ChevronLeftIcon />
          Previous
        </Button>
        <Button variant="contained" disabled={!this.nextButtonEnabled()} size="small" className={classes.button}>
          Next
          <ChevronRightIcon />
        </Button>
        <Button variant="contained" disabled={!this.lastButtonEnabled()} size="small" className={classes.button}>
          Last
          <LastPageIcon />
        </Button>
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
    this.setState({
      Id: storyResponse.Id,
      Url: storyResponse.Url,
      Author: storyResponse.Author,
      Title: storyResponse.Title,
      NumChapters: storyResponse.NumChapters,
      CurrentChapter: 0,
    });
  }

  firstButtonEnabled() {
    return this.state.CurrentChapter > 0;
  }
  lastButtonEnabled() {
    return this.state.CurrentChapter < this.state.NumChapters-1;
  }
  prevButtonEnabled() {
    return this.firstButtonEnabled();
  }
  nextButtonEnabled() {
    return this.lastButtonEnabled();
  }
}

Story.propTypes = {
  classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(Story);