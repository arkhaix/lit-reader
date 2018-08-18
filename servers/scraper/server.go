package scraper

import (
	context "golang.org/x/net/context"

	api "github.com/arkhaix/lit-reader/api/scraper"
	"github.com/arkhaix/lit-reader/pkg/scraper"
)

// Server implements api.ScraperServer
type Server struct{}

// CheckStoryURL returns true if the specified URL matches the expected
// pattern of a story supported by this service
func (s *Server) CheckStoryURL(ctx context.Context, in *api.CheckStoryURLRequest) (*api.CheckStoryURLResponse, error) {
	allowed := scraper.CheckStoryURL(in.GetUrl())

	return &api.CheckStoryURLResponse{Allowed: allowed}, nil
}

// FetchStoryMetadata fetches the title, author, and chapter index of a story,
// but not the actual chapter text
func (s *Server) FetchStoryMetadata(ctx context.Context, in *api.FetchStoryMetadataRequest) (*api.FetchStoryMetadataResponse, error) {
	story, err := scraper.FetchStoryMetadata(in.GetUrl())

	return &api.FetchStoryMetadataResponse{
		Url:         story.URL,
		Title:       story.Title,
		Author:      story.Author,
		NumChapters: int32(len(story.Chapters)),
	}, err
}

// FetchChapter fetches one chapter of a story
func (s *Server) FetchChapter(ctx context.Context, in *api.FetchChapterRequest) (*api.FetchChapterResponse, error) {
	storyURL := in.GetStoryUrl()
	chapterIndex := int(in.GetChapterIndex())

	chapter, err := scraper.FetchChapter(storyURL, chapterIndex)

	return &api.FetchChapterResponse{
		Url:   chapter.URL,
		Title: chapter.Title,
		Text:  chapter.Text,
		Html:  chapter.HTML,
	}, err
}
