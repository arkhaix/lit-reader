package story

import (
	context "golang.org/x/net/context"

	log "github.com/sirupsen/logrus"

	pbcommon "github.com/arkhaix/lit-reader/api/common"
	pb "github.com/arkhaix/lit-reader/api/story"
)

// Server implements api.StoryServiceServer
type Server struct{}

// CreateStory returns the story id and metadata for the queried url, scraping it first if necessary.
func (*Server) CreateStory(_ context.Context, req *pb.CreateStoryRequest) (*pb.CreateStoryResponse, error) {
	log.Infof("In: CreateStory(%s)", req.GetUrl())

	story := &pb.Story{
		Id:          "1234",
		Url:         req.GetUrl(),
		Title:       "Title",
		Author:      "Author",
		NumChapters: 1,
	}

	log.Infof("Out: CreateStory(%s): %s", req.GetUrl(), story.Id)

	return &pb.CreateStoryResponse{
		Status: &pbcommon.Status{
			StatusCode: 200,
			StatusText: "ok",
		},
		Data: story,
	}, nil
}

// GetStory returns the metadata for a previously-scraped story.
func (*Server) GetStory(_ context.Context, req *pb.GetStoryRequest) (*pb.GetStoryResponse, error) {
	log.Infof("In: GetStory(%s)", req.GetId())

	story := &pb.Story{
		Id:          req.GetId(),
		Url:         "example.com",
		Title:       "Title",
		Author:      "Author",
		NumChapters: 1,
	}

	log.Infof("Out: GetStory(%s): %s", req.GetId(), story.Url)

	return &pb.GetStoryResponse{
		Status: &pbcommon.Status{
			StatusCode: 200,
			StatusText: "ok",
		},
		Data: story,
	}, nil
}
