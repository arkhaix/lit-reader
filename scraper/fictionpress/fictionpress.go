package fictionpress

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly"

	lit "github.com/arkhaix/lit-reader/common"
)

// FictionPressScraper implements Scraper
type FictionPressScraper struct {
}

// NewScraper returns an empty FictionPressScraper
func NewScraper() FictionPressScraper {
	return FictionPressScraper{}
}

var baseURL *url.URL
var storyPattern []*regexp.Regexp
var chapterPattern *regexp.Regexp
var chapterSelectPattern *regexp.Regexp

func init() {
	baseURL, _ = url.Parse("https://www.fictionpress.com")
	storyPattern = []*regexp.Regexp{
		// This first format is not supported yet because it doesn't provide an easy way to get the
		// URLs of the chapters.  The suffix name is needed, and I don't see a good way to get it yet.
		// regexp.MustCompile("/s/[0-9]+/[0-9]+$"),       // https://www.fictionpress.com/s/2961893
		regexp.MustCompile("/s/[0-9]+/[0-9]+/[^/]+$"), // https://www.fictionpress.com/s/2961893/1/Mother-of-Learning
	}
	chapterPattern = regexp.MustCompile("/s/([0-9]+)/([0-9]+)/([^/]+)$")
	chapterSelectPattern = regexp.MustCompile("[0-9]+\\. (.*)$")
}

// IsSupportedStoryURL returns true if the specified URL matches the expected
// pattern of a story supported by this parser
func (FictionPressScraper) IsSupportedStoryURL(path string) bool {
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

	validPattern := false
	for _, re := range storyPattern {
		if re.FindStringIndex(path) != nil {
			validPattern = true
		}
	}
	if validPattern == false {
		return false
	}

	return true
}

// FetchStoryMetadata fetches the title, author, and chapter index of a story
func (scraper FictionPressScraper) FetchStoryMetadata(path string) (lit.Story, error) {
	story := lit.Story{}

	// validate
	path, err := forceBaseURL(path)
	if err != nil {
		return story, err
	}
	if scraper.IsSupportedStoryURL(path) == false {
		return story, errors.New("Invalid story URL: " + path)
	}

	// Parse the story and chapter parts from the path
	pathSuffix, err := url.Parse(path)
	if err != nil {
		return story, errors.New("Invalid chapter URL: " + path)
	}
	matches := chapterPattern.FindStringSubmatch(pathSuffix.Path)
	if matches == nil || len(matches) < 4 {
		return story, errors.New("Invalid chapter URL: " + path)
	}
	storyID := matches[1]
	// chapterID := matches[2]
	storySuffix := matches[3]

	// init
	story.URL = buildChapterURL(storyID, storySuffix, 1)

	c := colly.NewCollector(
		colly.AllowedDomains(baseURL.Hostname()),
	)

	// title
	c.OnHTML("#profile_top b", func(e *colly.HTMLElement) {
		story.Title = strings.TrimSpace(e.Text)
	})

	// author
	c.OnHTML("#profile_top a", func(e *colly.HTMLElement) {
		// the ":first" selector doesn't seem to work in go-colly, so do it manually
		if len(story.Author) == 0 {
			story.Author = strings.TrimSpace(e.Text)
		}
	})

	// chapter index
	var parseError error
	c.OnHTML("#chap_select > option", func(e *colly.HTMLElement) {
		chapterIndex, err := strconv.Atoi(e.Attr("value"))
		if err != nil {
			parseError = err
		}

		chapterTitleMatches := chapterSelectPattern.FindStringSubmatch(e.Text)
		if chapterTitleMatches == nil || len(chapterTitleMatches) < 2 {
			parseError = errors.New("Failed to parse the chapter titles")
		}
		chapterTitle := chapterTitleMatches[1]

		link := buildChapterURL(storyID, storySuffix, chapterIndex)
		linkURL, err := url.Parse(link)
		if err != nil {
			fmt.Println("Error parsing link", link)
			fmt.Println(err)
		}
		absoluteLink := baseURL.ResolveReference(linkURL)
		story.Chapters = append(story.Chapters, lit.Chapter{
			Title: chapterTitle,
			URL:   absoluteLink.String(),
			HTML:  "",
			Text:  "",
		})
	})

	c.Visit(path)

	return story, parseError
}

// FetchChapter fetches the text of one chapter of a story, inserting it into the Story
func (FictionPressScraper) FetchChapter(story *lit.Story, index int) error {
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
	c.OnHTML("#storytext", func(e *colly.HTMLElement) {
		story.Chapters[index].Text = strings.TrimSpace(e.Text)
		story.Chapters[index].HTML, parseError = e.DOM.Html()
	})

	// fetch
	c.Visit(chapterURL)

	return parseError
}

func buildChapterURL(storyID string, storySuffix string, chapter int) string {
	return fmt.Sprintf("%s/s/%s/%d/%s", baseURL, storyID, chapter, storySuffix)
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
