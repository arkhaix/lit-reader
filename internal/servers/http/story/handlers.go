package story

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	log "github.com/sirupsen/logrus"

	api "github.com/arkhaix/lit-reader/api/story"
	httpcommon "github.com/arkhaix/lit-reader/internal/servers/http/common"
)

var (
	// Client is the gRPC client for communicating with the story service.
	// Set this before using the handlers
	Client api.StoryServiceClient

	// Timeout is the gRPC timeout.
	// Set this before using the handlers
	Timeout time.Duration
)

// PostStory returns the story id and metadata for the queried url, scraping it first if necessary.
func PostStory(w http.ResponseWriter, r *http.Request) {
	// Parse params
	requestData := &postStoryRequest{}
	if err := render.Bind(r, requestData); err != nil {
		render.Render(w, r, httpcommon.ErrInvalidRequest(err))
		return
	}
	storyURL := requestData.URL
	log.Debugf("In: PostStory(%s)", storyURL)

	// RPC
	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()
	log.Debugf("RpcOut: PostStory(%s)", storyURL)
	result, err := Client.CreateStory(ctx, &api.CreateStoryRequest{
		Url: storyURL,
	})

	if err != nil {
		log.Errorf("RpcFail: PostStory(%s): %s", storyURL, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Debugf("RpcGood: PostStory(%s): %d", storyURL, result.GetStatus().GetStatusCode())

	// Output
	response := newPostStoryResponse(result)
	log.Debugf("Out: PostStory(%s): %d", storyURL, response.Status.Code)
	render.Render(w, r, response)
}

// GetStory returns the metadata for a previously-scraped story.
func GetStory(w http.ResponseWriter, r *http.Request) {
	// Parse params
	id := chi.URLParam(r, "storyID")
	log.Debugf("In: GetStory(%s)", id)

	// RPC
	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()
	log.Debugf("RpcOut: GetStory(%s)", id)
	result, err := Client.GetStory(ctx, &api.GetStoryRequest{
		Id: id,
	})

	if err != nil {
		log.Errorf("RpcFail: GetStory(%s): %s", id, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Debugf("RpcGood: GetStory(%s): %d", id, result.GetStatus().GetStatusCode())

	// Output
	response := newGetStoryResponse(result)
	log.Debugf("Out: GetStory(%s): %d", id, response.Status.Code)
	render.Render(w, r, response)
}

type postStoryRequest struct {
	URL string
}

func (sr *postStoryRequest) Bind(r *http.Request) error {
	return nil
}

type storyResponse struct {
	Status      httpcommon.Status
	ID          string
	URL         string
	Title       string
	Author      string
	NumChapters int
}

func (storyResponse) Render(http.ResponseWriter, *http.Request) error {
	return nil
}

func newPostStoryResponse(pb *api.CreateStoryResponse) *storyResponse {
	return &storyResponse{
		Status:      httpcommon.NewStatusFromProto(pb.GetStatus()),
		ID:          pb.GetData().GetId(),
		URL:         pb.GetData().GetUrl(),
		Title:       pb.GetData().GetTitle(),
		Author:      pb.GetData().GetAuthor(),
		NumChapters: int(pb.GetData().GetNumChapters()),
	}
}

func newGetStoryResponse(pb *api.GetStoryResponse) *storyResponse {
	return &storyResponse{
		Status:      httpcommon.NewStatusFromProto(pb.GetStatus()),
		ID:          pb.GetData().GetId(),
		URL:         pb.GetData().GetUrl(),
		Title:       pb.GetData().GetTitle(),
		Author:      pb.GetData().GetAuthor(),
		NumChapters: int(pb.GetData().GetNumChapters()),
	}
}
