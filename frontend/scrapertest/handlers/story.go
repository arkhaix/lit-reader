package handlers

import (
	"fmt"
	"net/http"

	"golang.org/x/net/context"

	api "github.com/arkhaix/lit-reader/api/scraper"
)

// StorySearch proxies scraper.FetchStoryMetadata
func StorySearch(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "invalid method", http.StatusBadRequest)
		return
	}

	storyURL := r.URL.Query().Get("url")

	ctx, cancel := context.WithTimeout(context.Background(), ScraperTimeout)
	defer cancel()

	// result, err := client.CheckStoryURL(ctx, &api.CheckStoryURLRequest{Url: storyURL})
	result, err := ScraperClient.FetchStoryMetadata(ctx, &api.FetchStoryMetadataRequest{
		Url: storyURL,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// w.Write([]byte(fmt.Sprintf("{\"allowed\":%v}", result.GetAllowed())))
	// w.Write([]byte(spew.Sdump(result)))
	w.Write([]byte(fmt.Sprintf("url: '%s'\ntitle: '%s'\nauthor: '%s'\nchapters: %d",
		result.GetUrl(), result.GetTitle(), result.GetAuthor(), result.GetNumChapters())))
}
