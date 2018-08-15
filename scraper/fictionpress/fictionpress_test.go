package fictionpress_test

import (
	"testing"

	. "github.com/arkhaix/lit-reader/scraper/fictionpress"
)

func TestValidStoryURLSucceeds(t *testing.T) {
	s := NewScraper()
	if !s.IsSupportedStoryURL("https://www.fictionpress.com/s/2961893/1/Mother-of-Learning") {
		t.Fail()
	}
}

func TestHTTPProtocolSucceeds(t *testing.T) {
	// This should succeed because the protocol is ignored and forced to https
	s := NewScraper()
	if !s.IsSupportedStoryURL("http://www.fictionpress.com/s/2961893/1/Mother-of-Learning") {
		t.Fail()
	}
}

func TestEmptyProtocolSucceeds(t *testing.T) {
	// This should succeed because the protocol is ignored and forced to https
	s := NewScraper()
	if !s.IsSupportedStoryURL("www.fictionpress.com/s/2961893/1/Mother-of-Learning") {
		t.Fail()
	}
}

func TestInvalidProtocolSucceeds(t *testing.T) {
	// This should succeed because the protocol is ignored and forced to https
	s := NewScraper()
	if !s.IsSupportedStoryURL("ftp://www.fictionpress.com/s/2961893/1/Mother-of-Learning") {
		t.Fail()
	}
}

func TestEmptySubdomainFails(t *testing.T) {
	s := NewScraper()
	if s.IsSupportedStoryURL("https://fictionpress.com/s/2961893/1/Mother-of-Learning") {
		t.Fail()
	}
}

func TestInvalidSubdomainFails(t *testing.T) {
	s := NewScraper()
	if s.IsSupportedStoryURL("https://invalid.fictionpress.com/s/2961893/1/Mother-of-Learning") {
		t.Fail()
	}
}

func TestEmptyDomainFails(t *testing.T) {
	s := NewScraper()
	if s.IsSupportedStoryURL("https:///s/2961893/1/Mother-of-Learning") {
		t.Fail()
	}
}

func TestInvalidDomainFails(t *testing.T) {
	s := NewScraper()
	if s.IsSupportedStoryURL("https://www.example.com/s/2961893/1/Mother-of-Learning") {
		t.Fail()
	}
}

func TestEmptyTLDFails(t *testing.T) {
	s := NewScraper()
	if s.IsSupportedStoryURL("https://www.fictionpress./s/2961893/1/Mother-of-Learning") {
		t.Fail()
	}
}

func TestInvalidTLDFails(t *testing.T) {
	s := NewScraper()
	if s.IsSupportedStoryURL("https://www.fictionpress.net/s/2961893/1/Mother-of-Learning") {
		t.Fail()
	}
}

func TestEmptyPrefixFails(t *testing.T) {
	s := NewScraper()
	if s.IsSupportedStoryURL("https://www.fictionpress.com//2961893/1/Mother-of-Learning") {
		t.Fail()
	}
}

func TestInvalidPrefixFails(t *testing.T) {
	s := NewScraper()
	if s.IsSupportedStoryURL("https://www.fictionpress.com/st/2961893/1/Mother-of-Learning") {
		t.Fail()
	}
}

func TestEmptyStoryIDFails(t *testing.T) {
	s := NewScraper()
	if s.IsSupportedStoryURL("https://www.fictionpress.com/s//1/Mother-of-Learning") {
		t.Fail()
	}
}

func TestAlphaStoryIDFails(t *testing.T) {
	s := NewScraper()
	if s.IsSupportedStoryURL("https://www.fictionpress.com/s/A2961893/1/Mother-of-Learning") {
		t.Fail()
	}
}

func TestEmptyChapterIDFails(t *testing.T) {
	s := NewScraper()
	if s.IsSupportedStoryURL("https://www.fictionpress.com/s/2961893//Mother-of-Learning") {
		t.Fail()
	}
}

func TestAlphaChapterIDFails(t *testing.T) {
	s := NewScraper()
	if s.IsSupportedStoryURL("https://www.fictionpress.com/s/2961893/A/Mother-of-Learning") {
		t.Fail()
	}
}

func TestEmptySuffixFails(t *testing.T) {
	s := NewScraper()
	if s.IsSupportedStoryURL("https://www.fictionpress.com/s/2961893/1//") {
		t.Fail()
	}
}
