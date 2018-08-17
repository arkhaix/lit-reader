package wanderinginn_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	lit "github.com/arkhaix/lit-reader/common"
	. "github.com/arkhaix/lit-reader/scraper/wanderinginn"
)

var s Scraper

func init() {
	s = NewScraper()
}

func TestURLWithValidStoryURLSucceeds(t *testing.T) {
	assert.Equal(t, true, s.IsSupportedStoryURL("https://wanderinginn.com"))
}

func TestURLWithHTTPProtocolSucceeds(t *testing.T) {
	// This should succeed because the protocol is ignored and forced to https
	assert.Equal(t, true, s.IsSupportedStoryURL("http://wanderinginn.com"))
}

func TestURLWithEmptyProtocolSucceeds(t *testing.T) {
	// This should succeed because the protocol is ignored and forced to https
	assert.Equal(t, true, s.IsSupportedStoryURL("wanderinginn.com"))
}

func TestURLWithInvalidProtocolSucceeds(t *testing.T) {
	// This should succeed because the protocol is ignored and forced to https
	assert.Equal(t, true, s.IsSupportedStoryURL("ftp://wanderinginn.com"))
}

func TestUnparseableURLFails(t *testing.T) {
	assert.Equal(t, false, s.IsSupportedStoryURL("ht&tps://wanderinginn.com"))
	assert.Equal(t, false, s.IsSupportedStoryURL("https://wanderinginn.com/%^&"))
}

func TestURLWithInvalidSubdomainFails(t *testing.T) {
	assert.Equal(t, false, s.IsSupportedStoryURL("https://invalid.wanderinginn.com"))
}

func TestURLWithEmptyDomainFails(t *testing.T) {
	assert.Equal(t, false, s.IsSupportedStoryURL("https://.com"))
}

func TestURLWithInvalidDomainFails(t *testing.T) {
	assert.Equal(t, false, s.IsSupportedStoryURL("https://example.com"))
}

func TestURLWithEmptyTLDFails(t *testing.T) {
	assert.Equal(t, false, s.IsSupportedStoryURL("https://wanderinginn."))
}

func TestURLWithInvalidTLDFails(t *testing.T) {
	assert.Equal(t, false, s.IsSupportedStoryURL("https://wanderinginn.net"))
}

// FetchStoryMetadata failure tests

func TestFetchStoryWithUnparseableURLFails(t *testing.T) {
	_, err := s.FetchStoryMetadata("ht&tps://wanderinginn.com")
	assert.NotNil(t, err)

	_, err = s.FetchStoryMetadata("https://wanderinginn.com/%^&")
	assert.NotNil(t, err)
}

func TestFetchStoryWithWrongDomainFails(t *testing.T) {
	// For this scraper, the domain is the story link, so wrong domains are
	// invalid and should not be rewritten
	s := NewScraper()
	_, err := s.FetchStoryMetadata("https://example.com")
	assert.NotNil(t, err)
}

// FetchChapter failure tests

func TestFetchChapterWithNilStoryFails(t *testing.T) {
	err := s.FetchChapter(nil, 0)
	assert.NotNil(t, err)
}

func TestFetchChapterWithOutOfBoundsChapterIndexFails(t *testing.T) {
	story := lit.Story{
		Chapters: []lit.Chapter{lit.Chapter{}},
	}

	err := s.FetchChapter(&story, -1)
	assert.NotNil(t, err)

	err = s.FetchChapter(&story, 1)
	assert.NotNil(t, err)
}

func TestFetchChapterWithUnparseableChapterURLFails(t *testing.T) {
	story := lit.Story{
		Chapters: []lit.Chapter{lit.Chapter{URL: "ht&tps://wanderinginn.com"}},
	}
	err := s.FetchChapter(&story, 0)
	assert.NotNil(t, err)

	story = lit.Story{
		Chapters: []lit.Chapter{lit.Chapter{URL: "https://wanderinginn.com/%^&"}},
	}
	err = s.FetchChapter(&story, 0)
	assert.NotNil(t, err)
}
