// +build integration

package fictionpress_test

import (
	"crypto/sha256"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/arkhaix/lit-reader/scraper/fictionpress"
)

var storyURL string

func init() {
	storyURL = "https://www.fictionpress.com/s/2922431/1/A-Lucky-Apocalypse"
}

func TestFictionPressIntegration(t *testing.T) {
	s := NewScraper()
	story, err := s.FetchStoryMetadata(storyURL)
	assert.Nil(t, err, "Failed to fetch the story")

	expectedURL := storyURL
	expectedTitle := "A Lucky Apocalypse"
	expectedAuthor := "ShaperV"
	expectedChapters := 4

	// Validate the story metadata
	assert.Equal(t, expectedURL, story.URL, "URL must match")
	assert.Equal(t, expectedTitle, story.Title, "Title must match")
	assert.Equal(t, expectedAuthor, story.Author, "Author must match")
	assert.Equal(t, expectedChapters, len(story.Chapters), "Number of chapters must match")

	// Validate the data for a chapter
	s.FetchChapter(&story, 0)
	c := story.Chapters[0]
	expectedChapterURL := "https://www.fictionpress.com/s/2922431/1/A-Lucky-Apocalypse"
	expectedChapterTitle := "Chapter 1"
	expectedTextSum := "9c8fa415b39ae1a863c0a82161360d89d7595ff900e608ac661aa74928d803bb"
	expectedHTMLSum := "20ba190838a61d6ecfbdfaeb810366326220f6a2d9bb774955d9382060be888d"
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
	story, err := s.FetchStoryMetadata("https://www.example.com/s/2922431/1/A-Lucky-Apocalypse")
	assert.Nil(t, err)
	assert.Equal(t, "https://www.fictionpress.com/s/2922431/1/A-Lucky-Apocalypse", story.URL,
		"Incorrect domain must be rewritten")
}
