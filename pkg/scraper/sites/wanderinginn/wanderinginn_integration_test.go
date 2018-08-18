// +build integration

package wanderinginn_test

import (
	"crypto/sha256"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/arkhaix/lit-reader/pkg/scraper/sites/wanderinginn"
)

var storyURL string

func init() {
	storyURL = "https://wanderinginn.com"
}

func TestWanderingInnIntegration(t *testing.T) {
	s := NewScraper()
	story, err := s.FetchStoryMetadata(storyURL)
	if err != nil {
		t.Fatal("Failed to fetch the story", storyURL, err)
	}

	expectedURL := storyURL
	expectedTitle := "The Wandering Inn"
	expectedAuthor := "pirateaba"
	expectedChapters := 249

	// Validate the story metadata
	assert.Equal(t, expectedURL, story.URL, "URL must match")
	assert.Equal(t, expectedTitle, story.Title, "Title must match")
	assert.Equal(t, expectedAuthor, story.Author, "Author must match")
	assert.Equal(t, expectedChapters, len(story.Chapters), "Number of chapters must match")

	// Validate the data for a chapter
	c, err := s.FetchChapter(storyURL, 0)
	assert.Nil(t, err)

	expectedChapterURL := "https://wanderinginn.com/2016/07/27/1-00/"
	expectedChapterTitle := "1.00"
	expectedHTMLSum := "29543cfb7c01ea294688bca8aeb60afe0badf591aa202d0d51b17876e7ac98cf"
	htmlSum := sha256.Sum256([]byte(c.HTML))
	htmlSumStr := fmt.Sprintf("%x", htmlSum)

	assert.Equal(t, expectedChapterURL, c.URL, "Chapter URL must match")
	assert.Equal(t, expectedChapterTitle, c.Title, "Chapter title must match")
	assert.Equal(t, expectedHTMLSum, htmlSumStr, "Chapter HTML must match")
}

func TestFetchStoryWithInvalidPathRewrites(t *testing.T) {
	story, err := s.FetchStoryMetadata("https://wanderinginn.com/invalid")
	assert.Nil(t, err)
	assert.Equal(t, "https://wanderinginn.com", story.URL)
}

func TestFetchChapterWithOutOfBoundsChapterIndexFails(t *testing.T) {
	_, err := s.FetchChapter(storyURL, -1)
	assert.NotNil(t, err)

	_, err = s.FetchChapter(storyURL, 99999999)
	assert.NotNil(t, err)
}
