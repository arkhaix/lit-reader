// +build integration

package archiveofourown_test

import (
	"crypto/sha256"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/arkhaix/lit-reader/pkg/scraper/sites/archiveofourown"
	"github.com/arkhaix/lit-reader/pkg/scraper/wrapper"
)

var storyURL string

func init() {
	storyURL = "https://archiveofourown.org/works/11478249/chapters/25740126"
}

func TestArchiveOfOurOwnIntegration(t *testing.T) {
	s := NewScraper(wrapper.NewScraperWrapper())

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
	c, err := s.FetchChapter(storyURL, 0)
	assert.Nil(t, err)

	expectedChapterURL := "https://archiveofourown.org/works/11478249/chapters/25740126"
	expectedChapterTitle := "Taking the Fall"
	expectedHTMLSum := "16109d786c5210586ccd1315fa25bb3f4a0edc219df2d472642e7783c5de1944"
	htmlSum := sha256.Sum256([]byte(c.HTML))
	htmlSumStr := fmt.Sprintf("%x", htmlSum)

	assert.Equal(t, expectedChapterURL, c.URL, "Chapter URL must match")
	assert.Equal(t, expectedChapterTitle, c.Title, "Chapter title must match")
	assert.Equal(t, expectedHTMLSum, htmlSumStr, "Chapter HTML must match")
}

func TestFetchStoryWithWrongDomainRewrites(t *testing.T) {
	s := NewScraper(wrapper.NewScraperWrapper())

	story, err := s.FetchStoryMetadata("https://example.com/works/11478249/chapters/25740126")
	assert.Nil(t, err)
	assert.Equal(t, "https://archiveofourown.org/works/11478249/chapters/25740126", story.URL,
		"Incorrect domain must be rewritten")
}

func TestFetchChapterWithOutOfBoundsChapterIndexFails(t *testing.T) {
	s := NewScraper(wrapper.NewScraperWrapper())

	_, err := s.FetchChapter("https://archiveofourown.org/works/11478249/chapters/25740126", -1)
	assert.NotNil(t, err)

	_, err = s.FetchChapter("https://archiveofourown.org/works/11478249/chapters/25740126", 99999999)
	assert.NotNil(t, err)
}
