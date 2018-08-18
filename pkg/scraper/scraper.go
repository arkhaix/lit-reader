package scraper

import (
	"errors"

	"github.com/arkhaix/lit-reader/common"
	"github.com/arkhaix/lit-reader/pkg/scraper/sites/archiveofourown"
	"github.com/arkhaix/lit-reader/pkg/scraper/sites/fictionpress"
	"github.com/arkhaix/lit-reader/pkg/scraper/sites/royalroad"
	"github.com/arkhaix/lit-reader/pkg/scraper/sites/wanderinginn"
	"github.com/arkhaix/lit-reader/pkg/scraper/wrapper"
)

var scrapers []common.Scraper
var storyCache common.Cache

func init() {
	sw := wrapper.NewScraperWrapper()
	scrapers = append(scrapers, archiveofourown.NewScraper(sw))
	scrapers = append(scrapers, fictionpress.NewScraper(sw))
	scrapers = append(scrapers, royalroad.NewScraper(sw))
	scrapers = append(scrapers, wanderinginn.NewScraper(sw))
}

// CheckStoryURL returns true if the specified URL matches the expected
// pattern of a story supported by any scraper
func CheckStoryURL(url string) bool {
	if getScraper(url) == nil {
		return false
	}
	return true
}

// FetchStoryMetadata fetches the title, author, and chapter index of a story,
// but not the actual chapter text
func FetchStoryMetadata(url string) (common.Story, error) {
	scraper := getScraper(url)
	if scraper == nil {
		return common.Story{}, errors.New("Unsupported URL")
	}
	return scraper.FetchStoryMetadata(url)
}

// FetchChapter fetches one chapter of a story
func FetchChapter(storyURL string, index int) (common.Chapter, error) {
	scraper := getScraper(storyURL)
	if scraper == nil {
		return common.Chapter{}, errors.New("Unsupported URL")
	}

	return scraper.FetchChapter(storyURL, index)
}

func getScraper(url string) common.Scraper {
	for _, scraper := range scrapers {
		if scraper.CheckStoryURL(url) == true {
			return scraper
		}
	}
	return nil
}
