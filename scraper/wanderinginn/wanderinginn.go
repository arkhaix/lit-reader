package wanderinginn

import (
	"net/url"
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
var tocURL string

func init() {
	baseURL, _ = url.Parse("https://wanderinginn.com")
	tocURL = "https://wanderinginn.com/table-of-contents/"
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

	return true
}

// FetchStoryMetadata fetches the title, author, and chapter index of a story
func (scraper Scraper) FetchStoryMetadata(path string) (lit.Story, error) {
	story := lit.Story{}

	// validate
	if !scraper.IsSupportedStoryURL(path) {
		return story, lit.NewScraperErrorString("Invalid story URL: " + path)
	}

	// init
	story.URL = baseURL.String()
	story.Author = "pirateaba"
	story.Title = "The Wandering Inn"

	c := colly.NewCollector(
		colly.AllowedDomains(baseURL.Hostname()),
	)

	// chapter index
	var callbackError error
	c.OnHTML("article .entry-content * > a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		// The table of contents links to wanderinginn.wordpress.com, so forceBaseURL
		// is needed here to rewrite the links to wanderinginn.com
		fixedLink, err := forceBaseURL(link)
		if err != nil {
			callbackError = err
		}
		linkURL, err := url.Parse(fixedLink)
		if err != nil {
			callbackError = err
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

	c.Visit(tocURL)

	if callbackError != nil {
		return story, lit.ScraperError{
			Err: callbackError,
		}
	}

	return story, nil
}

// FetchChapter fetches the text of one chapter of a story, inserting it into the Story
func (Scraper) FetchChapter(story *lit.Story, index int) error {
	// validate
	if story == nil {
		return lit.NewScraperErrorString("Story must not be nil")
	}
	if index < 0 || index >= len(story.Chapters) {
		return lit.NewScraperErrorString("Chapter index out of bounds")
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
	var callbackError error
	c.OnHTML("div.entry-content", func(e *colly.HTMLElement) {
		story.Chapters[index].Text = strings.TrimSpace(e.Text)
		story.Chapters[index].HTML, err = e.DOM.Html()
		if err != nil {
			callbackError = err
		}

		// Insert the chapter title at the beginning of the chapter
		story.Chapters[index].Text =
			story.Chapters[index].Title + "\n" + story.Chapters[index].Text
		story.Chapters[index].HTML =
			"<h1>" + story.Chapters[index].Title + "</h1>" + story.Chapters[index].HTML

		// Remove the "Previous Chapter" and "Next Chapter" links
		story.Chapters[index].Text = strings.Replace(story.Chapters[index].Text,
			"Previous Chapter", "", 1)
		story.Chapters[index].Text = strings.Replace(story.Chapters[index].Text,
			"Next Chapter", "", 1)
		story.Chapters[index].HTML = strings.Replace(story.Chapters[index].HTML,
			"Previous Chapter", "", 1)
		story.Chapters[index].HTML = strings.Replace(story.Chapters[index].HTML,
			"Next Chapter", "", 1)
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
		return lit.NewScraperError(callbackError)
	}
	return nil
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
