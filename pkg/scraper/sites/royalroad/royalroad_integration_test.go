// +build integration

package royalroad_test

import (
	"crypto/sha256"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/arkhaix/lit-reader/pkg/scraper/sites/royalroad"
)

var storyURL string

func init() {
	storyURL = "https://www.royalroad.com/fiction/15130/threadbare"
}

func TestRoyalRoadIntegration(t *testing.T) {
	s := NewScraper()
	story, err := s.FetchStoryMetadata(storyURL)
	if err != nil {
		t.Fatal("Failed to fetch the story", storyURL, err)
	}

	expectedURL := storyURL
	expectedTitle := "Threadbare"
	expectedAuthor := "Andrew Seiple"
	expectedChapters := 78

	// Validate the story metadata
	assert.Equal(t, expectedURL, story.URL, "URL must match")
	assert.Equal(t, expectedTitle, story.Title, "Title must match")
	assert.Equal(t, expectedAuthor, story.Author, "Author must match")
	assert.Equal(t, expectedChapters, len(story.Chapters), "Number of chapters must match")

	// Validate the data for a chapter
	c, err := s.FetchChapter(storyURL, 0)
	assert.Nil(t, err)

	expectedChapterURL := "https://www.royalroad.com/fiction/15130/threadbare/chapter/175199/awakening-1"
	expectedChapterTitle := "Awakening 1"
	expectedHTMLSum := "2ab31d5b1b2052070978e8c413cbb5c4fb611f2daaa5d69de931e773a35a7e2c"
	htmlSum := sha256.Sum256([]byte(c.HTML))
	htmlSumStr := fmt.Sprintf("%x", htmlSum)

	assert.Equal(t, expectedChapterURL, c.URL, "Chapter URL must match")
	assert.Equal(t, expectedChapterTitle, c.Title, "Chapter title must match")
	assert.Equal(t, expectedHTMLSum, htmlSumStr, "Chapter HTML must match")
}

func TestFetchStoryWithWrongDomainRewrites(t *testing.T) {
	s := NewScraper()
	story, err := s.FetchStoryMetadata("https://www.example.com/fiction/15130/threadbare")
	assert.Nil(t, err)
	assert.Equal(t, "https://www.royalroad.com/fiction/15130/threadbare", story.URL,
		"Incorrect domain must be rewritten")
}

func TestFetchChapterWithOutOfBoundsChapterIndexFails(t *testing.T) {
	_, err := s.FetchChapter("https://www.royalroad.com/fiction/5701/savage-divinity", -1)
	assert.NotNil(t, err)

	_, err = s.FetchChapter("https://www.royalroad.com/fiction/5701/savage-divinity", 99999999)
	assert.NotNil(t, err)
}
