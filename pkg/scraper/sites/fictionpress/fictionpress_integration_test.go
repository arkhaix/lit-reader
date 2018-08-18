// +build integration

package fictionpress_test

import (
	"crypto/sha256"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/arkhaix/lit-reader/pkg/scraper/sites/fictionpress"
	"github.com/arkhaix/lit-reader/pkg/scraper/wrapper"
)

var storyURL string

func init() {
	storyURL = "https://www.fictionpress.com/s/2922431/1/A-Lucky-Apocalypse"
}

func TestFictionPressIntegration(t *testing.T) {
	s := NewScraper(wrapper.NewScraperWrapper())

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
	c, err := s.FetchChapter(storyURL, 0)
	assert.Nil(t, err)

	expectedChapterURL := "https://www.fictionpress.com/s/2922431/1/A-Lucky-Apocalypse"
	expectedChapterTitle := "Chapter 1"
	expectedHTMLSum := "20ba190838a61d6ecfbdfaeb810366326220f6a2d9bb774955d9382060be888d"
	htmlSum := sha256.Sum256([]byte(c.HTML))
	htmlSumStr := fmt.Sprintf("%x", htmlSum)

	assert.Equal(t, expectedChapterURL, c.URL, "Chapter URL must match")
	assert.Equal(t, expectedChapterTitle, c.Title, "Chapter title must match")
	assert.Equal(t, expectedHTMLSum, htmlSumStr, "Chapter HTML must match")
}

func TestFetchStoryWithWrongDomainRewrites(t *testing.T) {
	s := NewScraper(wrapper.NewScraperWrapper())

	story, err := s.FetchStoryMetadata("https://www.example.com/s/2922431/1/A-Lucky-Apocalypse")
	assert.Nil(t, err)
	assert.Equal(t, "https://www.fictionpress.com/s/2922431/1/A-Lucky-Apocalypse", story.URL,
		"Incorrect domain must be rewritten")
}

func TestFetchChapterWithOutOfBoundsChapterIndexFails(t *testing.T) {
	s := NewScraper(wrapper.NewScraperWrapper())

	_, err := s.FetchChapter("https://www.fictionpress.com/s/2922431/1/A-Lucky-Apocalypse", -1)
	assert.NotNil(t, err)

	_, err = s.FetchChapter("https://www.fictionpress.com/s/2922431/1/A-Lucky-Apocalypse", 99999999)
	assert.NotNil(t, err)
}
