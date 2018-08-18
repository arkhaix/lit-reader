package scraper

import (
	"errors"

	"github.com/arkhaix/lit-reader/common"
	"github.com/arkhaix/lit-reader/scraper/sites/archiveofourown"
	"github.com/arkhaix/lit-reader/scraper/sites/fictionpress"
	"github.com/arkhaix/lit-reader/scraper/sites/royalroad"
	"github.com/arkhaix/lit-reader/scraper/sites/wanderinginn"
)

var scrapers []common.Scraper

func init() {
	scrapers = append(scrapers, archiveofourown.NewScraper())
	scrapers = append(scrapers, fictionpress.NewScraper())
	scrapers = append(scrapers, royalroad.NewScraper())
	scrapers = append(scrapers, wanderinginn.NewScraper())
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
