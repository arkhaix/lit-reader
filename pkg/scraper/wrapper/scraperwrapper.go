package wrapper

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/arkhaix/lit-reader/common"
	"github.com/arkhaix/lit-reader/internal/cache/local/lru"
)

// ScraperWrapper provides decorators for scrapers.  It is used internally by scrapers.
type ScraperWrapper struct {
	storyCache common.Cache
}

// NewScraperWrapper returns a new ScraperWrapper
func NewScraperWrapper() *ScraperWrapper {
	c, _ := lru.NewCache(500)
	res := ScraperWrapper{
		storyCache: c,
	}
	return &res
}

var storyMetadataTTL time.Duration

func init() {
	storyMetadataTTL = 12 * time.Hour
}

// FetchStoryMetadataFunc is here for readability
type FetchStoryMetadataFunc func(string) (common.Story, error)

// FetchStoryMetadata wraps a FetchStoryMetadata call with decorators
func (s *ScraperWrapper) FetchStoryMetadata(url string, next FetchStoryMetadataFunc) (common.Story, error) {
	story := common.Story{}

	storyString, ok := s.storyCache.Get(cacheKeyFromURL(url))
	if ok {
		err := json.Unmarshal([]byte(storyString), &story)
		if err == nil {
			return story, nil
		}
		// if it can't be unmarshalled, then kill the cache entry
		// this should only be possible via external interference
		s.storyCache.Delete(url)
	}

	story, err := next(url)

	if err == nil {
		storyBytes, err := json.Marshal(story)
		if err == nil {
			url = cacheKeyFromURL(url)
			s.storyCache.Put(url, string(storyBytes), storyMetadataTTL)
		}
	}

	return story, err
}

func cacheKeyFromURL(url string) string {
	protocolIndex := strings.Index(url, "://")
	if protocolIndex < 0 {
		return url
	}
	return url[protocolIndex+3:]
}
