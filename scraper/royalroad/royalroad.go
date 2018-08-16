package royalroad

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/gocolly/colly"

	lit "github.com/arkhaix/lit-reader/common"
)

// Scraper implements common.Scraper
type Scraper struct {
}

// NewScraper returns an empty Scraper
func NewScraper() Scraper {
	return Scraper{}
}

var baseURL *url.URL
var storyPattern *regexp.Regexp

func init() {
	baseURL, _ = url.Parse("https://www.royalroad.com")
	storyPattern = regexp.MustCompile("/fiction/[0-9]+/[^/]+$")
}

// IsSupportedStoryURL returns true if the specified URL matches the expected
// pattern of a story supported by this parser
func (Scraper) IsSupportedStoryURL(path string) bool {
	if !strings.Contains(path, "://") {
		path = "https://" + path
	}

	u, err := url.Parse(path)
	if err != nil {
		return false
	}
	if u.Hostname() != baseURL.Hostname() {
		return false
	}
	if storyPattern.FindStringIndex(path) == nil {
		return false
	}
	return true
}

// FetchStoryMetadata fetches the title, author, and chapter index of a story
func (scraper Scraper) FetchStoryMetadata(path string) (lit.Story, error) {
	story := lit.Story{}

	// validate
	path, err := forceBaseURL(path)
	if err != nil {
		return story, errors.New("Invalid story URL: " + path)
	}
	if scraper.IsSupportedStoryURL(path) == false {
		return story, errors.New("Invalid story URL: " + path)
	}

	// init
	story.URL = path

	c := colly.NewCollector(
		colly.AllowedDomains(baseURL.Hostname()),
	)

	// title
	c.OnHTML(".fic-title h1", func(e *colly.HTMLElement) {
		story.Title = strings.TrimSpace(e.Text)
	})

	// author
	c.OnHTML(".fic-title h4 span a", func(e *colly.HTMLElement) {
		story.Author = strings.TrimSpace(e.Text)
	})

	// chapter index
	c.OnHTML("#chapters tbody tr td a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		linkURL, err := url.Parse(link)
		if err != nil {
			fmt.Println("Error parsing link", link)
			fmt.Println(err)
		}
		absoluteLink := baseURL.ResolveReference(linkURL)
		linkText := strings.TrimSpace(e.Text)
		story.Chapters = append(story.Chapters, lit.Chapter{
			Title: linkText,
			URL:   absoluteLink.String(),
			HTML:  "",
			Text:  "",
		})
	})

	c.Visit(path)

	return story, nil
}

// FetchChapter fetches the text of one chapter of a story, inserting it into the Story
func (Scraper) FetchChapter(story *lit.Story, index int) error {
	// validate
	if story == nil {
		return errors.New("Story must not be nil")
	}
	if index < 0 || index >= len(story.Chapters) {
		return errors.New("Chapter index out of bounds")
	}
	chapterURL, err := forceBaseURL(story.Chapters[index].URL)
	if err != nil {
		return err
	}

	// init
	c := colly.NewCollector(
		colly.AllowedDomains(baseURL.Hostname()),
	)

	// parse
	var parseError error
	c.OnHTML(".chapter-content", func(e *colly.HTMLElement) {
		story.Chapters[index].Text = strings.TrimSpace(e.Text)
		story.Chapters[index].HTML, parseError = e.DOM.Html()
	})

	// fetch
	c.Visit(chapterURL)

	return parseError
}

// forceBaseURL rewrites the url to start with baseURL.
// This forces an https connection instead of whatever protocol is in the url.
func forceBaseURL(path string) (string, error) {
	origURL, err := url.Parse(path)
	if err != nil {
		return "", err
	}

	pathOnly, err := url.Parse(origURL.Path)
	if err != nil {
		return "", err
	}

	return baseURL.ResolveReference(pathOnly).String(), nil
}
