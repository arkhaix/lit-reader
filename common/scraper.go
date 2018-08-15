package common

// Story contains all the text and metadata for one story
type Story struct {
	URL      string
	Title    string
	Author   string
	Chapters []Chapter
}

// Chapter contains all the text and metadata for one chapter of a story
type Chapter struct {
	Title string
	URL   string
	HTML  string
	Text  string
}

// Scraper is the interface implemented by each site scraper for fetching stories
type Scraper interface {
	// IsSupportedStoryURL returns true if the specified URL matches the expected
	// pattern of a story supported by this parser
	IsSupportedStoryURL(path string) bool

	// FetchStoryMetadata fetches the title, author, and chapter index of a story,
	// but not the actual chapter text
	FetchStoryMetadata(path string) (Story, error)

	// FetchChapter fetches one chapter of a story
	FetchChapter(story *Story, index int) error
}
