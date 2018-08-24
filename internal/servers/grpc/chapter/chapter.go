package chapter

import (
	"database/sql"
	"time"

	context "golang.org/x/net/context"

	log "github.com/sirupsen/logrus"

	"github.com/arkhaix/lit-reader/internal/servers/grpc/common"

	apichapter "github.com/arkhaix/lit-reader/api/chapter"
	apicommon "github.com/arkhaix/lit-reader/api/common"
	apiscraper "github.com/arkhaix/lit-reader/api/scraper"
	apistory "github.com/arkhaix/lit-reader/api/story"
)

// Server implements api.ChapterServiceServer
type Server struct {
	StoryClient  apistory.StoryServiceClient
	StoryTimeout time.Duration

	ScraperClient  apiscraper.ScraperServiceClient
	ScraperTimeout time.Duration

	DB *sql.DB
}

// GetChapter returns the chapter data for the requested chapter,
// scraping it first if necessary.
func (s *Server) GetChapter(_ context.Context, req *apichapter.GetChapterRequest) (*apichapter.GetChapterResponse, error) {
	storyID := req.GetStoryId()
	chapterID := int(req.GetChapterId())

	log.Infof("In: GetChapter(%s:%d)", storyID, chapterID)

	// Check db first
	chapter, err := s.fetchChapterFromDb(storyID, chapterID)
	if err != nil {

		// Chapter is not in the db, so we need to scrape it
		// But scraper needs the story url, not the id
		// So get the url from the story service first
		storyURL, status, err := s.fetchStoryURL(storyID)
		if err != nil || status.GetCode() != 200 {
			return &apichapter.GetChapterResponse{
				Status: status,
				Data:   nil,
			}, nil
		}

		// Now we can scrape the chapter
		chapter, status, err = s.fetchChapterFromScraper(storyURL, chapterID)
		if err != nil {
			log.Errorf("Scraper failed to fetch a chapter: %s", err.Error())
			return &apichapter.GetChapterResponse{
				Status: status,
				Data:   nil,
			}, nil
		}
		chapter.StoryId = storyID
		chapter.ChapterId = int32(chapterID)

		// And store it
		err = s.saveChapterToDb(chapter)
		if err != nil {
			log.Errorf("Failed to save chapter to db: %s", err.Error())
			return &apichapter.GetChapterResponse{
				Status: common.StatusInternalServerError,
				Data:   nil,
			}, nil
		}
	}

	log.Infof("Out: GetChapter(%s, %d): %s", storyID, chapterID, chapter.Title)

	return &apichapter.GetChapterResponse{
		Status: common.StatusOk,
		Data:   chapter,
	}, nil
}

func (s *Server) fetchStoryURL(storyID string) (string, *apicommon.Status, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.StoryTimeout)
	defer cancel()
	response, err := s.StoryClient.GetStory(ctx, &apistory.GetStoryRequest{Id: storyID})
	return response.GetData().GetUrl(), response.GetStatus(), err
}

func (s *Server) fetchChapterFromScraper(storyURL string, chapterID int) (*apichapter.Chapter, *apicommon.Status, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.ScraperTimeout)
	defer cancel()
	scraperResponse, err := s.ScraperClient.FetchChapter(ctx, &apiscraper.FetchChapterRequest{
		StoryUrl:     storyURL,
		ChapterIndex: int32(chapterID),
	})
	status := common.StatusOk
	if err != nil {
		status = common.StatusBadRequest
	}
	return &apichapter.Chapter{
		Url:   scraperResponse.GetUrl(),
		Title: scraperResponse.GetTitle(),
		Html:  scraperResponse.GetHtml(),
	}, status, err
}
