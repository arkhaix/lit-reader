package story

import (
	"time"

	context "golang.org/x/net/context"

	log "github.com/sirupsen/logrus"

	"github.com/arkhaix/lit-reader/internal/servers/grpc/common"

	apicommon "github.com/arkhaix/lit-reader/api/common"
	apiscraper "github.com/arkhaix/lit-reader/api/scraper"
	apistory "github.com/arkhaix/lit-reader/api/story"
)

// Server implements api.StoryServiceServer
type Server struct {
	ScraperClient  apiscraper.ScraperServiceClient
	ScraperTimeout time.Duration
}

// CreateStory returns the story id and metadata for the queried url, scraping it first if necessary.
func (s *Server) CreateStory(_ context.Context, req *apistory.CreateStoryRequest) (*apistory.CreateStoryResponse, error) {
	log.Infof("In: CreateStory(%s)", req.GetUrl())

	// TODO: db read

	response := s.fetchStoryFromScraper(req.GetUrl())

	// TODO: db store

	log.Infof("Out: CreateStory(%s): %s", req.GetUrl(), response.GetData().GetId())

	return response, nil
}

// GetStory returns the metadata for a previously-scraped story.
func (*Server) GetStory(_ context.Context, req *apistory.GetStoryRequest) (*apistory.GetStoryResponse, error) {
	log.Infof("In: GetStory(%s)", req.GetId())

	// Todo: db read

	story := &apistory.Story{
		Id:          req.GetId(),
		Url:         "example.com",
		Title:       "Title",
		Author:      "Author",
		NumChapters: 1,
	}

	log.Infof("Out: GetStory(%s): %s", req.GetId(), story.Url)

	return &apistory.GetStoryResponse{
		Status: &apicommon.Status{
			StatusCode: 200,
			StatusText: "ok",
		},
		Data: story,
	}, nil
}

func (s *Server) fetchStoryFromScraper(url string) *apistory.CreateStoryResponse {
	ctx, cancel := context.WithTimeout(context.Background(), s.ScraperTimeout)
	defer cancel()

	// Make sure the URL is a valid pattern
	checkReq := &apiscraper.CheckStoryURLRequest{
		Url: url,
	}
	checkResult, err := s.ScraperClient.CheckStoryURL(ctx, checkReq)
	if err != nil {
		log.Errorf("Rpc failed: CheckStoryUrl(%s)", url)
		return &apistory.CreateStoryResponse{
			Status: common.StatusInternalServerError,
			Data:   nil,
		}
	}
	if checkResult.GetAllowed() == false {
		return &apistory.CreateStoryResponse{
			Status: common.StatusBadRequest,
			Data:   nil,
		}
	}

	// Fetch the story metadata
	fetchReq := &apiscraper.FetchStoryMetadataRequest{
		Url: url,
	}
	result, err := s.ScraperClient.FetchStoryMetadata(ctx, fetchReq)
	if err != nil {
		log.Errorf("Rpc failed: FetchStoryMetadata(%s)", url)
		return &apistory.CreateStoryResponse{
			Status: common.StatusInternalServerError,
			Data:   nil,
		}
	}

	return &apistory.CreateStoryResponse{
		Status: common.StatusOk,
		Data: &apistory.Story{
			Id:          "1234",
			Url:         result.GetUrl(),
			Title:       result.GetTitle(),
			Author:      result.GetAuthor(),
			NumChapters: result.GetNumChapters(),
		},
	}
}
