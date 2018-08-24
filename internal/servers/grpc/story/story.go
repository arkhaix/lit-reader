package story

import (
	"database/sql"
	"time"

	context "golang.org/x/net/context"

	log "github.com/sirupsen/logrus"

	"github.com/arkhaix/lit-reader/internal/servers/grpc/common"

	apiscraper "github.com/arkhaix/lit-reader/api/scraper"
	apistory "github.com/arkhaix/lit-reader/api/story"
)

// Server implements api.StoryServiceServer
type Server struct {
	ScraperClient  apiscraper.ScraperServiceClient
	ScraperTimeout time.Duration
	DB             *sql.DB
}

// CreateStory returns the story id and metadata for the queried url, scraping it first if necessary.
func (s *Server) CreateStory(ctx context.Context, req *apistory.CreateStoryRequest) (*apistory.CreateStoryResponse, error) {
	url := req.GetUrl()
	log.Infof("In: CreateStory(%s)", url)

	response := &apistory.CreateStoryResponse{}

	id, err := s.queryStoryByURL(url)
	log.Infof("Query (%s) returned id %s", url, id)
	if len(id) > 0 && err == nil {
		getResponse, _ := s.GetStory(ctx, &apistory.GetStoryRequest{Id: id})
		response.Status = getResponse.Status
		response.Data = getResponse.Data
	} else {
		response = s.fetchStoryFromScraper(url)
		if response.Status.GetCode() == 200 && response.Data != nil {
			id, err = s.saveStoryToDb(response.Data)
			if err == nil {
				response.Data.Id = id
			}
		}
	}

	log.Infof("Out: CreateStory(%s): %s", url, response.GetData().GetId())
	return response, nil
}

// GetStory returns the metadata for a previously-scraped story.
func (s *Server) GetStory(_ context.Context, req *apistory.GetStoryRequest) (*apistory.GetStoryResponse, error) {

	log.Infof("In: GetStory(%s)", req.GetId())

	story, err := s.fetchStoryFromDb(req.GetId())
	status := common.StatusOk
	if err != nil {
		status = common.StatusNotFound
	}

	log.Infof("Out: GetStory(%s): %s", req.GetId(), story.Url)

	return &apistory.GetStoryResponse{
		Status: status,
		Data:   story,
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
			Url:         result.GetUrl(),
			Title:       result.GetTitle(),
			Author:      result.GetAuthor(),
			NumChapters: result.GetNumChapters(),
		},
	}
}
