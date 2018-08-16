package fictionpress_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/arkhaix/lit-reader/scraper/fictionpress"
)

func TestValidStoryURLSucceeds(t *testing.T) {
	s := NewScraper()
	assert.Equal(t, true, s.IsSupportedStoryURL("https://www.fictionpress.com/s/2961893/1/Mother-of-Learning"))
}

func TestHTTPProtocolSucceeds(t *testing.T) {
	// This should succeed because the protocol is ignored and forced to https
	s := NewScraper()
	assert.Equal(t, true, s.IsSupportedStoryURL("http://www.fictionpress.com/s/2961893/1/Mother-of-Learning"))
}

func TestEmptyProtocolSucceeds(t *testing.T) {
	// This should succeed because the protocol is ignored and forced to https
	s := NewScraper()
	assert.Equal(t, true, s.IsSupportedStoryURL("www.fictionpress.com/s/2961893/1/Mother-of-Learning"))
}

func TestInvalidProtocolSucceeds(t *testing.T) {
	// This should succeed because the protocol is ignored and forced to https
	s := NewScraper()
	assert.Equal(t, true, s.IsSupportedStoryURL("ftp://www.fictionpress.com/s/2961893/1/Mother-of-Learning"))
}

func TestEmptySubdomainFails(t *testing.T) {
	s := NewScraper()
	assert.Equal(t, false, s.IsSupportedStoryURL("https://fictionpress.com/s/2961893/1/Mother-of-Learning"))
}

func TestInvalidSubdomainFails(t *testing.T) {
	s := NewScraper()
	assert.Equal(t, false, s.IsSupportedStoryURL("https://invalid.fictionpress.com/s/2961893/1/Mother-of-Learning"))
}

func TestEmptyDomainFails(t *testing.T) {
	s := NewScraper()
	assert.Equal(t, false, s.IsSupportedStoryURL("https:///s/2961893/1/Mother-of-Learning"))
}

func TestInvalidDomainFails(t *testing.T) {
	s := NewScraper()
	assert.Equal(t, false, s.IsSupportedStoryURL("https://www.example.com/s/2961893/1/Mother-of-Learning"))
}

func TestEmptyTLDFails(t *testing.T) {
	s := NewScraper()
	assert.Equal(t, false, s.IsSupportedStoryURL("https://www.fictionpress./s/2961893/1/Mother-of-Learning"))
}

func TestInvalidTLDFails(t *testing.T) {
	s := NewScraper()
	assert.Equal(t, false, s.IsSupportedStoryURL("https://www.fictionpress.net/s/2961893/1/Mother-of-Learning"))
}

func TestEmptyPrefixFails(t *testing.T) {
	s := NewScraper()
	assert.Equal(t, false, s.IsSupportedStoryURL("https://www.fictionpress.com//2961893/1/Mother-of-Learning"))
}

func TestInvalidPrefixFails(t *testing.T) {
	s := NewScraper()
	assert.Equal(t, false, s.IsSupportedStoryURL("https://www.fictionpress.com/st/2961893/1/Mother-of-Learning"))
}

func TestEmptyStoryIDFails(t *testing.T) {
	s := NewScraper()
	assert.Equal(t, false, s.IsSupportedStoryURL("https://www.fictionpress.com/s//1/Mother-of-Learning"))
}

func TestAlphaStoryIDFails(t *testing.T) {
	s := NewScraper()
	assert.Equal(t, false, s.IsSupportedStoryURL("https://www.fictionpress.com/s/A2961893/1/Mother-of-Learning"))
}

func TestEmptyChapterIDFails(t *testing.T) {
	s := NewScraper()
	assert.Equal(t, false, s.IsSupportedStoryURL("https://www.fictionpress.com/s/2961893//Mother-of-Learning"))
}

func TestAlphaChapterIDFails(t *testing.T) {
	s := NewScraper()
	assert.Equal(t, false, s.IsSupportedStoryURL("https://www.fictionpress.com/s/2961893/A/Mother-of-Learning"))
}

func TestEmptySuffixFails(t *testing.T) {
	s := NewScraper()
	assert.Equal(t, false, s.IsSupportedStoryURL("https://www.fictionpress.com/s/2961893/1//"))
}
