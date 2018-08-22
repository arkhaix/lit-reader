package scraper

import (
	context "golang.org/x/net/context"

	log "github.com/sirupsen/logrus"

	api "github.com/arkhaix/lit-reader/api/scraper"
	"github.com/arkhaix/lit-reader/pkg/scraper"
)

// Server implements api.ScraperServiceServer
type Server struct{}

// CheckStoryURL returns true if the specified URL matches the expected
// pattern of a story supported by this service
func (s *Server) CheckStoryURL(ctx context.Context, in *api.CheckStoryURLRequest) (*api.CheckStoryURLResponse, error) {
	log.Infof("In: CheckStoryURL(%s)", in.GetUrl())
	allowed := scraper.CheckStoryURL(in.GetUrl())
	log.Infof("Out: CheckStoryURL(%s): %v", in.GetUrl(), allowed)

	return &api.CheckStoryURLResponse{Allowed: allowed}, nil
}

// FetchStoryMetadata fetches the title, author, and chapter index of a story,
// but not the actual chapter text
func (s *Server) FetchStoryMetadata(ctx context.Context, in *api.FetchStoryMetadataRequest) (*api.FetchStoryMetadataResponse, error) {
	log.Infof("In: FetchStoryMetadata(%s)", in.GetUrl())
	story, err := scraper.FetchStoryMetadata(in.GetUrl())
	log.Infof("Out: FetchStoryMetadata(%s): %s", in.GetUrl(), story.Title)

	return &api.FetchStoryMetadataResponse{
		Url:         story.URL,
		Title:       story.Title,
		Author:      story.Author,
		NumChapters: int32(len(story.Chapters)),
	}, err
}

// FetchChapter fetches one chapter of a story
func (s *Server) FetchChapter(ctx context.Context, in *api.FetchChapterRequest) (*api.FetchChapterResponse, error) {
	log.Infof("In: FetchChapter(%s, %d)", in.GetStoryUrl(), in.GetChapterIndex())
	storyURL := in.GetStoryUrl()
	chapterIndex := int(in.GetChapterIndex())

	chapter, err := scraper.FetchChapter(storyURL, chapterIndex)
	log.Infof("Out: FetchChapter(%s, %d): %s", in.GetStoryUrl(), in.GetChapterIndex(), chapter.Title)

	return &api.FetchChapterResponse{
		Url:   chapter.URL,
		Title: chapter.Title,
		Html:  chapter.HTML,
	}, err
}
