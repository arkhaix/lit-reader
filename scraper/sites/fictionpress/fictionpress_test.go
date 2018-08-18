package fictionpress_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	lit "github.com/arkhaix/lit-reader/common"
	. "github.com/arkhaix/lit-reader/scraper/sites/fictionpress"
)

var s Scraper

func init() {
	s = NewScraper()
}

// CheckStoryURL tests

func TestURLWithValidStoryURLSucceeds(t *testing.T) {
	assert.Equal(t, true, s.CheckStoryURL("https://www.fictionpress.com/s/2961893/1/Mother-of-Learning"))
}

func TestURLWithHTTPProtocolSucceeds(t *testing.T) {
	// This should succeed because the protocol is ignored and forced to https
	assert.Equal(t, true, s.CheckStoryURL("http://www.fictionpress.com/s/2961893/1/Mother-of-Learning"))
}

func TestURLWithEmptyProtocolSucceeds(t *testing.T) {
	// This should succeed because the protocol is ignored and forced to https
	assert.Equal(t, true, s.CheckStoryURL("www.fictionpress.com/s/2961893/1/Mother-of-Learning"))
}

func TestURLWithInvalidProtocolSucceeds(t *testing.T) {
	// This should succeed because the protocol is ignored and forced to https
	assert.Equal(t, true, s.CheckStoryURL("ftp://www.fictionpress.com/s/2961893/1/Mother-of-Learning"))
}

func TestUnparseableURLFails(t *testing.T) {
	assert.Equal(t, false, s.CheckStoryURL("ht&tps://www.fictionpress.com/s/2961893/1/Mother-of-Learning"))
	assert.Equal(t, false, s.CheckStoryURL("https://www.fictionpress.com/%^&s/2961893/1/Mother-of-Learning"))
}

func TestURLWithEmptySubdomainFails(t *testing.T) {
	assert.Equal(t, false, s.CheckStoryURL("https://fictionpress.com/s/2961893/1/Mother-of-Learning"))
}

func TestURLWithInvalidSubdomainFails(t *testing.T) {
	assert.Equal(t, false, s.CheckStoryURL("https://invalid.fictionpress.com/s/2961893/1/Mother-of-Learning"))
}

func TestURLWithEmptyDomainFails(t *testing.T) {
	assert.Equal(t, false, s.CheckStoryURL("https:///s/2961893/1/Mother-of-Learning"))
}

func TestURLWithInvalidDomainFails(t *testing.T) {
	assert.Equal(t, false, s.CheckStoryURL("https://www.example.com/s/2961893/1/Mother-of-Learning"))
}

func TestURLWithEmptyTLDFails(t *testing.T) {
	assert.Equal(t, false, s.CheckStoryURL("https://www.fictionpress./s/2961893/1/Mother-of-Learning"))
}

func TestURLWithInvalidTLDFails(t *testing.T) {
	assert.Equal(t, false, s.CheckStoryURL("https://www.fictionpress.net/s/2961893/1/Mother-of-Learning"))
}

func TestURLWithEmptyPrefixFails(t *testing.T) {
	assert.Equal(t, false, s.CheckStoryURL("https://www.fictionpress.com//2961893/1/Mother-of-Learning"))
}

func TestURLWithInvalidPathFails(t *testing.T) {
	assert.Equal(t, false, s.CheckStoryURL("https://www.fictionpress.com/st/2961893/1/Mother-of-Learning"))
}

func TestURLWithEmptyStoryIDFails(t *testing.T) {
	assert.Equal(t, false, s.CheckStoryURL("https://www.fictionpress.com/s//1/Mother-of-Learning"))
}

func TestURLWithAlphaStoryIDFails(t *testing.T) {
	assert.Equal(t, false, s.CheckStoryURL("https://www.fictionpress.com/s/A2961893/1/Mother-of-Learning"))
}

func TestURLWithEmptyChapterIDFails(t *testing.T) {
	assert.Equal(t, false, s.CheckStoryURL("https://www.fictionpress.com/s/2961893//Mother-of-Learning"))
}

func TestURLWithAlphaChapterIDFails(t *testing.T) {
	assert.Equal(t, false, s.CheckStoryURL("https://www.fictionpress.com/s/2961893/A/Mother-of-Learning"))
}

func TestURLWithEmptySuffixFails(t *testing.T) {
	assert.Equal(t, false, s.CheckStoryURL("https://www.fictionpress.com/s/2961893/1//"))
}

// FetchStoryMetadata failure tests

func TestFetchStoryWithUnparseableProtocolFails(t *testing.T) {
	_, err := s.FetchStoryMetadata("ht&tps://www.fictionpress.com/s/2961893/1/Mother-of-Learning")
	assert.NotNil(t, err)
}

func TestFetchStoryWithInvalidPathFails(t *testing.T) {
	_, err := s.FetchStoryMetadata("https://www.fictionpress.com/%^&s/2961893/1/Mother-of-Learning")
	assert.NotNil(t, err)

	_, err = s.FetchStoryMetadata("https://www.fictionpress.com/s/A2961893/1/Mother-of-Learning")
	assert.NotNil(t, err)

	_, err = s.FetchStoryMetadata("https://www.fictionpress.com/s/2961893/A1/Mother-of-Learning")
	assert.NotNil(t, err)

	_, err = s.FetchStoryMetadata("https://www.fictionpress.com/s/2961893/1//")
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
		Chapters: []lit.Chapter{lit.Chapter{URL: "ht&tps://www.fictionpress.com/s/2922431/1/A-Lucky-Apocalypse"}},
	}
	err := s.FetchChapter(&story, 0)
	assert.NotNil(t, err)

	story = lit.Story{
		Chapters: []lit.Chapter{lit.Chapter{URL: "https://www.fictionpress.com/%^&s/2922431/1/A-Lucky-Apocalypse"}},
	}
	err = s.FetchChapter(&story, 0)
	assert.NotNil(t, err)
}
