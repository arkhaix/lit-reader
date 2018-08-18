package royalroad_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/arkhaix/lit-reader/pkg/scraper/sites/royalroad"
)

var s Scraper

func init() {
	s = NewScraper()
}

func TestURLWithValidStoryURLSucceeds(t *testing.T) {
	assert.Equal(t, true, s.CheckStoryURL("https://www.royalroad.com/fiction/5701/savage-divinity"))
}

func TestURLWithHTTPProtocolSucceeds(t *testing.T) {
	// This should succeed because the protocol is ignored and forced to https
	assert.Equal(t, true, s.CheckStoryURL("http://www.royalroad.com/fiction/5701/savage-divinity"))
}

func TestURLWithEmptyProtocolSucceeds(t *testing.T) {
	// This should succeed because the protocol is ignored and forced to https
	assert.Equal(t, true, s.CheckStoryURL("www.royalroad.com/fiction/5701/savage-divinity"))
}

func TestURLWithInvalidProtocolSucceeds(t *testing.T) {
	// This should succeed because the protocol is ignored and forced to https
	assert.Equal(t, true, s.CheckStoryURL("ftp://www.royalroad.com/fiction/5701/savage-divinity"))
}

func TestUnparseableURLFails(t *testing.T) {
	assert.Equal(t, false, s.CheckStoryURL("ht&tps://www.royalroad.com/fiction/5701/savage-divinity"))
	assert.Equal(t, false, s.CheckStoryURL("https://www.royalroad.com/%^&fiction/5701/savage-divinity"))
}

func TestURLWithEmptySubdomainFails(t *testing.T) {
	assert.Equal(t, false, s.CheckStoryURL("https://royalroad.com/fiction/5701/savage-divinity"))
}

func TestURLWithInvalidSubdomainFails(t *testing.T) {
	assert.Equal(t, false, s.CheckStoryURL("https://invalid.royalroad.com/fiction/5701/savage-divinity"))
}

func TestURLWithEmptyDomainFails(t *testing.T) {
	assert.Equal(t, false, s.CheckStoryURL("https://www..com/fiction/5701/savage-divinity"))
}

func TestURLWithInvalidDomainFails(t *testing.T) {
	assert.Equal(t, false, s.CheckStoryURL("https://www.example.com/fiction/5701/savage-divinity"))
}

func TestURLWithEmptyTLDFails(t *testing.T) {
	assert.Equal(t, false, s.CheckStoryURL("https://www.royalroad./fiction/5701/savage-divinity"))
}

func TestURLWithInvalidTLDFails(t *testing.T) {
	assert.Equal(t, false, s.CheckStoryURL("https://www.royalroad.net/fiction/5701/savage-divinity"))
}

func TestURLWithEmptyPrefixFails(t *testing.T) {
	assert.Equal(t, false, s.CheckStoryURL("https://www.royalroad.com//5701/savage-divinity"))
}

func TestURLWithInvalidPrefixFails(t *testing.T) {
	assert.Equal(t, false, s.CheckStoryURL("https://www.royalroad.com/invalid/5701/savage-divinity"))
}

func TestURLWithEmptyStoryIDFails(t *testing.T) {
	assert.Equal(t, false, s.CheckStoryURL("https://www.royalroad.com/fiction//savage-divinity"))
}

func TestURLWithAlphaStoryIDFails(t *testing.T) {
	assert.Equal(t, false, s.CheckStoryURL("https://www.royalroad.com/fiction/A5701/savage-divinity"))
}

func TestURLWithEmptySuffixFails(t *testing.T) {
	assert.Equal(t, false, s.CheckStoryURL("https://www.royalroad.com/fiction/5701/"))
}

// FetchStoryMetadata failure tests

func TestFetchStoryWithUnparseableURLFails(t *testing.T) {
	_, err := s.FetchStoryMetadata("ht&tps://www.royalroad.com/fiction/5701/savage-divinity")
	assert.NotNil(t, err)

	_, err = s.FetchStoryMetadata("https://www.royalroad.com/%^&fiction/5701/savage-divinity")
	assert.NotNil(t, err)
}

func TestFetchStoryWithInvalidPathFails(t *testing.T) {
	_, err := s.FetchStoryMetadata("https://www.royalroad.com/invalid/5701/savage-divinity")
	assert.NotNil(t, err)

	_, err = s.FetchStoryMetadata("https://www.royalroad.com/fiction/invalid5701/savage-divinity")
	assert.NotNil(t, err)

	_, err = s.FetchStoryMetadata("https://www.royalroad.com/fiction/5701//")
	assert.NotNil(t, err)
}

// FetchChapter failure tests

func TestFetchChapterWithEmptyStoryFails(t *testing.T) {
	_, err := s.FetchChapter("", 0)
	assert.NotNil(t, err)
}

func TestFetchChapterWithUnparseableChapterURLFails(t *testing.T) {
	_, err := s.FetchChapter("ht&tps://www.royalroad.com/fiction/5701/savage-divinity/chapter/58095/chapter-1-new-beginnings", 0)
	assert.NotNil(t, err)

	_, err = s.FetchChapter("https://www.royalroad.com/%^&fiction/5701/savage-divinity/chapter/58095/chapter-1-new-beginnings", 0)
	assert.NotNil(t, err)
}
