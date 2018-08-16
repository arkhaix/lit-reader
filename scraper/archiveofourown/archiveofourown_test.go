package archiveofourown_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	lit "github.com/arkhaix/lit-reader/common"
	. "github.com/arkhaix/lit-reader/scraper/archiveofourown"
)

var s Scraper

func init() {
	s = NewScraper()
}

func TestURLWithValidStoryURLSucceeds(t *testing.T) {
	assert.Equal(t, true, s.IsSupportedStoryURL("https://archiveofourown.org/works/11478249/chapters/25740126"))
}

func TestURLWithHTTPProtocolSucceeds(t *testing.T) {
	// This should succeed because the protocol is ignored and forced to https
	assert.Equal(t, true, s.IsSupportedStoryURL("http://archiveofourown.org/works/11478249/chapters/25740126"))
}

func TestURLWithEmptyProtocolSucceeds(t *testing.T) {
	// This should succeed because the protocol is ignored and forced to https
	assert.Equal(t, true, s.IsSupportedStoryURL("archiveofourown.org/works/11478249/chapters/25740126"))
}

func TestURLWithInvalidProtocolSucceeds(t *testing.T) {
	// This should succeed because the protocol is ignored and forced to https
	assert.Equal(t, true, s.IsSupportedStoryURL("ftp://archiveofourown.org/works/11478249/chapters/25740126"))
}

func TestUnparseableURLFails(t *testing.T) {
	assert.Equal(t, false, s.IsSupportedStoryURL("ht&tps://archiveofourown.org/works/11478249/chapters/25740126"))
	assert.Equal(t, false, s.IsSupportedStoryURL("https://archiveofourown.org/%^&works/11478249/chapters/25740126"))
}

func TestURLWithInvalidSubdomainFails(t *testing.T) {
	assert.Equal(t, false, s.IsSupportedStoryURL("https://invalid.archiveofourown.org/works/11478249/chapters/25740126"))
}

func TestURLWithEmptyDomainFails(t *testing.T) {
	assert.Equal(t, false, s.IsSupportedStoryURL("https://.org/works/11478249/chapters/25740126"))
}

func TestURLWithInvalidDomainFails(t *testing.T) {
	assert.Equal(t, false, s.IsSupportedStoryURL("https://example.com/works/11478249/chapters/25740126"))
}

func TestURLWithEmptyTLDFails(t *testing.T) {
	assert.Equal(t, false, s.IsSupportedStoryURL("https://archiveofourown./works/11478249/chapters/25740126"))
}

func TestURLWithInvalidTLDFails(t *testing.T) {
	assert.Equal(t, false, s.IsSupportedStoryURL("https://archiveofourown.net/works/11478249/chapters/25740126"))
}

func TestURLWithEmptyPrefixFails(t *testing.T) {
	assert.Equal(t, false, s.IsSupportedStoryURL("https://archiveofourown.org//11478249/chapters/25740126"))
}

func TestURLWithInvalidPrefixFails(t *testing.T) {
	assert.Equal(t, false, s.IsSupportedStoryURL("https://archiveofourown.org/invalid/11478249/chapters/25740126"))
}

func TestURLWithEmptyStoryIDFails(t *testing.T) {
	assert.Equal(t, false, s.IsSupportedStoryURL("https://archiveofourown.org/works//chapters/25740126"))
}

func TestURLWithAlphaStoryIDFails(t *testing.T) {
	assert.Equal(t, false, s.IsSupportedStoryURL("https://archiveofourown.org/works/A11478249/chapters/25740126"))
}

func TestURLWithEmptyChapterPrefixFails(t *testing.T) {
	assert.Equal(t, false, s.IsSupportedStoryURL("https://archiveofourown.org/works/11478249//25740126"))
}

func TestURLWithInvalidChapterPrefixFails(t *testing.T) {
	assert.Equal(t, false, s.IsSupportedStoryURL("https://archiveofourown.org/works/11478249/invalid/25740126"))
}

func TestURLWithEmptyChapterIDFails(t *testing.T) {
	assert.Equal(t, false, s.IsSupportedStoryURL("https://archiveofourown.org/works/11478249/chapters//"))
}

func TestURLWithInvalidChapterIDFails(t *testing.T) {
	assert.Equal(t, false, s.IsSupportedStoryURL("https://archiveofourown.org/works/11478249/chapters/A25740126"))
}

// FetchStoryMetadata failure tests

func TestFetchStoryWithUnparseableURLFails(t *testing.T) {
	_, err := s.FetchStoryMetadata("ht&tps://archiveofourown.org/works/11478249/chapters/25740126")
	assert.NotNil(t, err)

	_, err = s.FetchStoryMetadata("https://archiveofourown.org/%^&works/11478249/chapters/25740126")
	assert.NotNil(t, err)
}

func TestFetchStoryWithInvalidPathFails(t *testing.T) {
	_, err := s.FetchStoryMetadata("https://archiveofourown.org/invalid/11478249/chapters/25740126")
	assert.NotNil(t, err)

	_, err = s.FetchStoryMetadata("https://archiveofourown.org/works/invalid11478249/chapters/25740126")
	assert.NotNil(t, err)

	_, err = s.FetchStoryMetadata("https://archiveofourown.org/works/11478249/invalid/25740126")
	assert.NotNil(t, err)

	_, err = s.FetchStoryMetadata("https://archiveofourown.org/works/11478249/chapters/invalid25740126")
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
		Chapters: []lit.Chapter{lit.Chapter{URL: "ht&tps://archiveofourown.org/works/11478249/chapters/25740126"}},
	}
	err := s.FetchChapter(&story, 0)
	assert.NotNil(t, err)

	story = lit.Story{
		Chapters: []lit.Chapter{lit.Chapter{URL: "https://archiveofourown.org/%^&works/11478249/chapters/25740126"}},
	}
	err = s.FetchChapter(&story, 0)
	assert.NotNil(t, err)
}
