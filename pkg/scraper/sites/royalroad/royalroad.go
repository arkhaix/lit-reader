package royalroad

import (
	"net/url"
	"regexp"
	"strings"

	"github.com/gocolly/colly"

	"github.com/arkhaix/lit-reader/common"
	"github.com/arkhaix/lit-reader/pkg/scraper/wrapper"
)

// Scraper implements common.Scraper
type Scraper struct {
	wrapper *wrapper.ScraperWrapper
}

// NewScraper returns an initialized Scraper
func NewScraper(wrapper *wrapper.ScraperWrapper) Scraper {
	return Scraper{
		wrapper: wrapper,
	}
}

var baseURL *url.URL
var storyPattern *regexp.Regexp

func init() {
	baseURL, _ = url.Parse("https://www.royalroad.com")
	storyPattern = regexp.MustCompile("/fiction/[0-9]+/[^/]+$")
}

// CheckStoryURL returns true if the specified URL matches the expected
// pattern of a story supported by this parser
func (Scraper) CheckStoryURL(path string) bool {
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
func (scraper Scraper) FetchStoryMetadata(path string) (common.Story, error) {
	return scraper.wrapper.FetchStoryMetadata(path, scraper.fetchStoryMetadata)
}

func (scraper Scraper) fetchStoryMetadata(path string) (common.Story, error) {
	story := common.Story{}

	// validate
	path, err := forceBaseURL(path)
	if err != nil {
		return story, common.NewScraperErrorString("Invalid story URL: " + path)
	}
	if scraper.CheckStoryURL(path) == false {
		return story, common.NewScraperErrorString("Invalid story URL: " + path)
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
	var callbackError error
	c.OnHTML("#chapters tbody tr td a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		linkURL, err := url.Parse(link)
		if err != nil {
			callbackError = err
		}
		absoluteLink := baseURL.ResolveReference(linkURL)
		linkText := strings.TrimSpace(e.Text)
		story.Chapters = append(story.Chapters, common.Chapter{
			Title: linkText,
			URL:   absoluteLink.String(),
			HTML:  "",
		})
	})

	// errors
	c.OnError(func(r *colly.Response, err error) {
		if err != nil {
			callbackError = err
		}
	})

	c.Visit(path)

	if callbackError != nil {
		return story, common.ScraperError{
			Err: callbackError,
		}
	}

	return story, nil
}

// FetchChapter fetches the text of one chapter of a story, inserting it into the Story
func (scraper Scraper) FetchChapter(storyURL string, index int) (common.Chapter, error) {
	chapter := common.Chapter{}

	story, err := scraper.FetchStoryMetadata(storyURL)

	// validate
	if err != nil {
		return chapter, common.NewScraperError(err)
	}
	if index < 0 || index >= len(story.Chapters) {
		return chapter, common.NewScraperErrorString("Chapter index out of bounds")
	}
	chapterURL, err := forceBaseURL(story.Chapters[index].URL)
	if err != nil {
		return chapter, err
	}

	// init
	c := colly.NewCollector(
		colly.AllowedDomains(baseURL.Hostname()),
	)

	// parse
	var callbackError error
	c.OnHTML(".chapter-content", func(e *colly.HTMLElement) {
		story.Chapters[index].HTML, err = e.DOM.Html()
		if err != nil {
			callbackError = err
		}
		chapter = story.Chapters[index]
	})

	// errors
	c.OnError(func(r *colly.Response, err error) {
		if err != nil {
			callbackError = err
		}
	})

	// fetch
	c.Visit(chapterURL)

	if callbackError != nil {
		return chapter, common.NewScraperError(callbackError)
	}
	return chapter, nil
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
