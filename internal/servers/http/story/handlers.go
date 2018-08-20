package story

import (
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/context"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	api "github.com/arkhaix/lit-reader/api/scraper"
	"github.com/arkhaix/lit-reader/common"
	httpcommon "github.com/arkhaix/lit-reader/internal/servers/http/common"
)

var (
	// ScraperClient is the gRPC client for communicating with the scraper service.
	// Set this before using the handlers
	ScraperClient api.ScraperClient

	// ScraperTimeout is the gRPC timeout.
	// Set this before using the handlers
	ScraperTimeout time.Duration
)

type stories struct {
	m          sync.RWMutex
	stories    []common.Story
	storyIndex map[string]int
}

var data stories

func init() {
	data = stories{}
	data.stories = make([]common.Story, 0)
	data.storyIndex = make(map[string]int)
}

// GetStoryByID returns a story's metadata by its ID
func GetStoryByID(w http.ResponseWriter, r *http.Request) {
	// Parse params
	idStr := chi.URLParam(r, "storyID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		render.Render(w, r, httpcommon.ErrInvalidRequest(err))
		return
	}

	data.m.RLock()
	if id < 0 || id >= len(data.stories) {
		data.m.RUnlock()
		render.Render(w, r, httpcommon.ErrNotFound())
		return
	}
	story := data.stories[id]
	data.m.RUnlock()

	render.Render(w, r, storyResponseFromStory(id, story))
}

// PostStoryByURL retrieves a story's metadata from the provided URL, stores it,
// and returns it
func PostStoryByURL(w http.ResponseWriter, r *http.Request) {
	// Parse params
	requestData := &storyRequest{}
	if err := render.Bind(r, requestData); err != nil {
		render.Render(w, r, httpcommon.ErrInvalidRequest(err))
		return
	}
	storyURL := requestData.URL
	cacheKey := cacheKeyFromURL(storyURL)

	// Check index
	response := storyResponse{}
	data.m.RLock()
	storyIndex, ok := data.storyIndex[cacheKey]
	if ok {
		response = storyResponseFromStory(storyIndex, data.stories[storyIndex])
	}
	data.m.RUnlock()

	if !ok {
		// RPC
		ctx, cancel := context.WithTimeout(context.Background(), ScraperTimeout)
		defer cancel()
		result, err := ScraperClient.FetchStoryMetadata(ctx, &api.FetchStoryMetadataRequest{
			Url: storyURL,
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Store result
		data.m.Lock()
		storyIndex, ok = data.storyIndex[cacheKey]
		if !ok {
			data.stories = append(data.stories, storyFromProto(result))
			storyIndex = len(data.stories) - 1
			data.storyIndex[cacheKey] = storyIndex
		}
		response = storyResponseFromStory(storyIndex, data.stories[storyIndex])
		data.m.Unlock()
	}

	render.Render(w, r, response)
}

// GetChapterByID returns a story's metadata by its ID
func GetChapterByID(w http.ResponseWriter, r *http.Request) {
	// Parse params
	storyIDStr := chi.URLParam(r, "storyID")
	storyID, err := strconv.Atoi(storyIDStr)
	if err != nil {
		render.Render(w, r, httpcommon.ErrInvalidRequest(err))
		return
	}
	chapterIDStr := chi.URLParam(r, "chapterID")
	chapterID, err := strconv.Atoi(chapterIDStr)
	if err != nil {
		render.Render(w, r, httpcommon.ErrInvalidRequest(err))
		return
	}

	data.m.RLock()
	if storyID < 0 || storyID >= len(data.stories) {
		data.m.RUnlock()
		render.Render(w, r, httpcommon.ErrNotFound())
		return
	}
	story := data.stories[storyID]

	if chapterID < 0 || chapterID >= len(story.Chapters) {
		data.m.RUnlock()
		render.Render(w, r, httpcommon.ErrNotFound())
	}
	chapter := story.Chapters[chapterID]
	data.m.RUnlock()

	// RPC
	if len(chapter.HTML) == 0 {
		ctx, cancel := context.WithTimeout(context.Background(), ScraperTimeout)
		defer cancel()
		result, err := ScraperClient.FetchChapter(ctx, &api.FetchChapterRequest{
			StoryUrl:     story.URL,
			ChapterIndex: int32(chapterID),
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Store result
		chapter.URL = result.GetUrl()
		chapter.Title = result.GetTitle()
		chapter.HTML = result.GetHtml()
		data.m.Lock()
		if len(data.stories[storyID].Chapters[chapterID].HTML) == 0 {
			data.stories[storyID].Chapters[chapterID] = chapter
		}
		data.m.Unlock()
	}

	render.Render(w, r, chapterResponseFromChapter(chapterID, chapter))
}

type storyRequest struct {
	URL string
}

func (sr *storyRequest) Bind(r *http.Request) error {
	return nil
}

type storyResponse struct {
	ID          int
	URL         string
	Title       string
	Author      string
	NumChapters int
}

func (storyResponse) Render(http.ResponseWriter, *http.Request) error {
	return nil
}

type chapterResponse struct {
	ID    int
	URL   string
	Title string
	HTML  string
}

func (chapterResponse) Render(http.ResponseWriter, *http.Request) error {
	return nil
}

func cacheKeyFromURL(url string) string {
	protocolIndex := strings.Index(url, "://")
	if protocolIndex < 0 {
		return url
	}
	return url[protocolIndex+3:]
}

func storyFromProto(pb *api.FetchStoryMetadataResponse) common.Story {
	return common.Story{
		URL:      pb.GetUrl(),
		Title:    pb.GetTitle(),
		Author:   pb.GetAuthor(),
		Chapters: make([]common.Chapter, pb.GetNumChapters()),
	}
}

func storyResponseFromStory(id int, s common.Story) storyResponse {
	return storyResponse{
		ID:          id,
		URL:         s.URL,
		Title:       s.Title,
		Author:      s.Author,
		NumChapters: len(s.Chapters),
	}
}

func chapterResponseFromChapter(id int, c common.Chapter) chapterResponse {
	return chapterResponse{
		ID:    id,
		URL:   c.URL,
		Title: c.Title,
		HTML:  c.HTML,
	}
}
