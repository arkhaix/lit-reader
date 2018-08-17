// +build integration

package archiveofourown_test

import (
	"crypto/sha256"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/arkhaix/lit-reader/scraper/sites/archiveofourown"
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
	expectedTextSum := "29af0585a00bfbf187a1ac142b7f0648ff217daafc56766c4cdccfaa41a560ba"
	expectedHTMLSum := "16109d786c5210586ccd1315fa25bb3f4a0edc219df2d472642e7783c5de1944"
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
