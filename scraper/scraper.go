package scraper

import (
	"errors"

	lit "github.com/arkhaix/lit-reader/common"
	"github.com/arkhaix/lit-reader/scraper/archiveofourown"
	"github.com/arkhaix/lit-reader/scraper/fictionpress"
	"github.com/arkhaix/lit-reader/scraper/royalroad"
)

var scrapers []lit.Scraper

func init() {
	scrapers = append(scrapers, archiveofourown.NewScraper())
	scrapers = append(scrapers, fictionpress.NewScraper())
	scrapers = append(scrapers, royalroad.NewScraper())
}

// IsSupportedStoryURL returns true if the specified URL matches the expected
// pattern of a story supported by any scraper
func IsSupportedStoryURL(url string) bool {
	if getScraper(url) == nil {
		return false
	}
	return true
}

// FetchStoryMetadata fetches the title, author, and chapter index of a story,
// but not the actual chapter text
func FetchStoryMetadata(url string) (lit.Story, error) {
	scraper := getScraper(url)
	if scraper == nil {
		return lit.Story{}, errors.New("Unsupported URL")
	}
	return scraper.FetchStoryMetadata(url)
}

// FetchChapter fetches one chapter of a story
func FetchChapter(story *lit.Story, index int) error {
	if story == nil {
		return errors.New("story must not be nil")
	}

	scraper := getScraper(story.URL)
	if scraper == nil {
		return errors.New("Unsupported URL")
	}

	return scraper.FetchChapter(story, index)
}

func getScraper(url string) lit.Scraper {
	for _, scraper := range scrapers {
		if scraper.IsSupportedStoryURL(url) == true {
			return scraper
		}
	}
	return nil
}
