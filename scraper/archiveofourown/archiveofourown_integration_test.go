// +build integration

package archiveofourown_test

import (
	"crypto/sha256"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/arkhaix/lit-reader/scraper/archiveofourown"
)

var storyURL string

func init() {
	storyURL = "https://archiveofourown.org/works/11478249/chapters/25740126"
}

func TestArchiveOfOurOwnIntegration(t *testing.T) {
	s := NewScraper()
	story, err := s.FetchStoryMetadata(storyURL)
	if err != nil {
		t.Fatal("Failed to fetch the story", storyURL, err)
	}

	expectedURL := storyURL
	expectedTitle := "Worth the Candle"
	expectedAuthor := "cthulhuraejepsen"
	expectedChapters := 118

	// Validate the story metadata
	assert.Equal(t, expectedURL, story.URL, "URL must match")
	assert.Equal(t, expectedTitle, story.Title, "Title must match")
	assert.Equal(t, expectedAuthor, story.Author, "Author must match")
	assert.Equal(t, expectedChapters, len(story.Chapters), "Number of chapters must match")

	// Validate the data for a chapter
	s.FetchChapter(&story, 0)
	c := story.Chapters[0]
	expectedChapterURL := "https://archiveofourown.org/works/11478249/chapters/25740126"
	expectedChapterTitle := "Taking the Fall"
	expectedTextSum := "0c203b899cd3e7913af050e7fc26e4d2faf67995353ea178f74b814d53f122e9"
	expectedHTMLSum := "98655b975c671efb178bdb3b43a0d3a46b892ef3dc1bc5c02f8024b5e3269497"
	textSum := sha256.Sum256([]byte(c.Text))
	textSumStr := fmt.Sprintf("%x", textSum)
	htmlSum := sha256.Sum256([]byte(c.HTML))
	htmlSumStr := fmt.Sprintf("%x", htmlSum)

	assert.Equal(t, expectedChapterURL, c.URL, "Chapter URL must match")
	assert.Equal(t, expectedChapterTitle, c.Title, "Chapter title must match")
	assert.Equal(t, expectedTextSum, textSumStr, "Chapter text must match")
	assert.Equal(t, expectedHTMLSum, htmlSumStr, "Chapter HTML must match")
}

func TestFetchStoryWithWrongDomainRewrites(t *testing.T) {
	s := NewScraper()
	story, err := s.FetchStoryMetadata("https://example.com/works/11478249/chapters/25740126")
	assert.Nil(t, err)
	assert.Equal(t, "https://archiveofourown.org/works/11478249/chapters/25740126", story.URL,
		"Incorrect domain must be rewritten")
}
