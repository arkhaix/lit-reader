package royalroad_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/arkhaix/lit-reader/scraper/royalroad"
)

func TestValidStoryURLSucceeds(t *testing.T) {
	s := NewScraper()
	assert.Equal(t, true, s.IsSupportedStoryURL("https://www.royalroad.com/fiction/5701/savage-divinity"))
}

func TestHTTPProtocolSucceeds(t *testing.T) {
	// This should succeed because the protocol is ignored and forced to https
	s := NewScraper()
	assert.Equal(t, true, s.IsSupportedStoryURL("http://www.royalroad.com/fiction/5701/savage-divinity"))
}

func TestEmptyProtocolSucceeds(t *testing.T) {
	// This should succeed because the protocol is ignored and forced to https
	s := NewScraper()
	assert.Equal(t, true, s.IsSupportedStoryURL("www.royalroad.com/fiction/5701/savage-divinity"))
}

func TestInvalidProtocolSucceeds(t *testing.T) {
	// This should succeed because the protocol is ignored and forced to https
	s := NewScraper()
	assert.Equal(t, true, s.IsSupportedStoryURL("ftp://www.royalroad.com/fiction/5701/savage-divinity"))
}

func TestEmptySubdomainFails(t *testing.T) {
	s := NewScraper()
	assert.Equal(t, false, s.IsSupportedStoryURL("https://royalroad.com/fiction/5701/savage-divinity"))
}

func TestInvalidSubdomainFails(t *testing.T) {
	s := NewScraper()
	assert.Equal(t, false, s.IsSupportedStoryURL("https://invalid.royalroad.com/fiction/5701/savage-divinity"))
}

func TestEmptyDomainFails(t *testing.T) {
	s := NewScraper()
	assert.Equal(t, false, s.IsSupportedStoryURL("https://www..com/fiction/5701/savage-divinity"))
}

func TestInvalidDomainFails(t *testing.T) {
	s := NewScraper()
	assert.Equal(t, false, s.IsSupportedStoryURL("https://www.example.com/fiction/5701/savage-divinity"))
}

func TestEmptyTLDFails(t *testing.T) {
	s := NewScraper()
	assert.Equal(t, false, s.IsSupportedStoryURL("https://www.royalroad./fiction/5701/savage-divinity"))
}

func TestInvalidTLDFails(t *testing.T) {
	s := NewScraper()
	assert.Equal(t, false, s.IsSupportedStoryURL("https://www.royalroad.net/fiction/5701/savage-divinity"))
}

func TestEmptyPrefixFails(t *testing.T) {
	s := NewScraper()
	assert.Equal(t, false, s.IsSupportedStoryURL("https://www.royalroad.com//5701/savage-divinity"))
}

func TestInvalidPrefixFails(t *testing.T) {
	s := NewScraper()
	assert.Equal(t, false, s.IsSupportedStoryURL("https://www.royalroad.com/invalid/5701/savage-divinity"))
}

func TestEmptyStoryIDFails(t *testing.T) {
	s := NewScraper()
	assert.Equal(t, false, s.IsSupportedStoryURL("https://www.royalroad.com/fiction//savage-divinity"))
}

func TestAlphaStoryIDFails(t *testing.T) {
	s := NewScraper()
	assert.Equal(t, false, s.IsSupportedStoryURL("https://www.royalroad.com/fiction/A5701/savage-divinity"))
}

func TestEmptySuffixFails(t *testing.T) {
	s := NewScraper()
	assert.Equal(t, false, s.IsSupportedStoryURL("https://www.royalroad.com/fiction/5701/"))
}
