package royalroad_test

import (
	"testing"

	. "github.com/arkhaix/lit-reader/scraper/royalroad"
)

func TestValidStoryURLSucceeds(t *testing.T) {
	s := NewScraper()
	if !s.IsSupportedStoryURL("https://www.royalroad.com/fiction/5701/savage-divinity") {
		t.Fail()
	}
}

func TestHTTPProtocolSucceeds(t *testing.T) {
	// This should succeed because the protocol is ignored and forced to https
	s := NewScraper()
	if !s.IsSupportedStoryURL("http://www.royalroad.com/fiction/5701/savage-divinity") {
		t.Fail()
	}
}

func TestEmptyProtocolSucceeds(t *testing.T) {
	// This should succeed because the protocol is ignored and forced to https
	s := NewScraper()
	if !s.IsSupportedStoryURL("www.royalroad.com/fiction/5701/savage-divinity") {
		t.Fail()
	}
}

func TestInvalidProtocolSucceeds(t *testing.T) {
	// This should succeed because the protocol is ignored and forced to https
	s := NewScraper()
	if !s.IsSupportedStoryURL("ftp://www.royalroad.com/fiction/5701/savage-divinity") {
		t.Fail()
	}
}

func TestEmptySubdomainFails(t *testing.T) {
	s := NewScraper()
	if s.IsSupportedStoryURL("https://royalroad.com/fiction/5701/savage-divinity") {
		t.Fail()
	}
}

func TestInvalidSubdomainFails(t *testing.T) {
	s := NewScraper()
	if s.IsSupportedStoryURL("https://invalid.royalroad.com/fiction/5701/savage-divinity") {
		t.Fail()
	}
}

func TestEmptyDomainFails(t *testing.T) {
	s := NewScraper()
	if s.IsSupportedStoryURL("https://www..com/fiction/5701/savage-divinity") {
		t.Fail()
	}
}

func TestInvalidDomainFails(t *testing.T) {
	s := NewScraper()
	if s.IsSupportedStoryURL("https://www.example.com/fiction/5701/savage-divinity") {
		t.Fail()
	}
}

func TestEmptyTLDFails(t *testing.T) {
	s := NewScraper()
	if s.IsSupportedStoryURL("https://www.royalroad./fiction/5701/savage-divinity") {
		t.Fail()
	}
}

func TestInvalidTLDFails(t *testing.T) {
	s := NewScraper()
	if s.IsSupportedStoryURL("https://www.royalroad.net/fiction/5701/savage-divinity") {
		t.Fail()
	}
}

func TestEmptyPrefixFails(t *testing.T) {
	s := NewScraper()
	if s.IsSupportedStoryURL("https://www.royalroad.com//5701/savage-divinity") {
		t.Fail()
	}
}

func TestInvalidPrefixFails(t *testing.T) {
	s := NewScraper()
	if s.IsSupportedStoryURL("https://www.royalroad.com/invalid/5701/savage-divinity") {
		t.Fail()
	}
}

func TestEmptyStoryIDFails(t *testing.T) {
	s := NewScraper()
	if s.IsSupportedStoryURL("https://www.royalroad.com/fiction//savage-divinity") {
		t.Fail()
	}
}

func TestAlphaStoryIDFails(t *testing.T) {
	s := NewScraper()
	if s.IsSupportedStoryURL("https://www.royalroad.com/fiction/A5701/savage-divinity") {
		t.Fail()
	}
}

func TestEmptySuffixFails(t *testing.T) {
	s := NewScraper()
	if s.IsSupportedStoryURL("https://www.royalroad.com/fiction/5701/") {
		t.Fail()
	}
}
