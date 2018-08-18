package wanderinginn_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/arkhaix/lit-reader/pkg/scraper/sites/wanderinginn"
	"github.com/arkhaix/lit-reader/pkg/scraper/wrapper"
)

var s Scraper

func init() {
	s = NewScraper(wrapper.NewScraperWrapper())
}

func TestURLWithValidStoryURLSucceeds(t *testing.T) {
	assert.Equal(t, true, s.CheckStoryURL("https://wanderinginn.com"))
}

func TestURLWithHTTPProtocolSucceeds(t *testing.T) {
	// This should succeed because the protocol is ignored and forced to https
	assert.Equal(t, true, s.CheckStoryURL("http://wanderinginn.com"))
}

func TestURLWithEmptyProtocolSucceeds(t *testing.T) {
	// This should succeed because the protocol is ignored and forced to https
	assert.Equal(t, true, s.CheckStoryURL("wanderinginn.com"))
}

func TestURLWithInvalidProtocolSucceeds(t *testing.T) {
	// This should succeed because the protocol is ignored and forced to https
	assert.Equal(t, true, s.CheckStoryURL("ftp://wanderinginn.com"))
}

func TestUnparseableURLFails(t *testing.T) {
	assert.Equal(t, false, s.CheckStoryURL("ht&tps://wanderinginn.com"))
	assert.Equal(t, false, s.CheckStoryURL("https://wanderinginn.com/%^&"))
}

func TestURLWithInvalidSubdomainFails(t *testing.T) {
	assert.Equal(t, false, s.CheckStoryURL("https://invalid.wanderinginn.com"))
}

func TestURLWithEmptyDomainFails(t *testing.T) {
	assert.Equal(t, false, s.CheckStoryURL("https://.com"))
}

func TestURLWithInvalidDomainFails(t *testing.T) {
	assert.Equal(t, false, s.CheckStoryURL("https://example.com"))
}

func TestURLWithEmptyTLDFails(t *testing.T) {
	assert.Equal(t, false, s.CheckStoryURL("https://wanderinginn."))
}

func TestURLWithInvalidTLDFails(t *testing.T) {
	assert.Equal(t, false, s.CheckStoryURL("https://wanderinginn.net"))
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
	_, err := s.FetchStoryMetadata("https://example.com")
	assert.NotNil(t, err)
}

// FetchChapter failure tests

func TestFetchChapterWithEmptyStoryFails(t *testing.T) {
	_, err := s.FetchChapter("", 0)
	assert.NotNil(t, err)
}

func TestFetchChapterWithUnparseableChapterURLFails(t *testing.T) {
	_, err := s.FetchChapter("ht&tps://wanderinginn.com", 0)
	assert.NotNil(t, err)

	_, err = s.FetchChapter("https://wanderinginn.com/%^&", 0)
	assert.NotNil(t, err)
}
