package chapter

import (
	"net/http"
	"strconv"
	"time"

	"golang.org/x/net/context"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	log "github.com/sirupsen/logrus"

	api "github.com/arkhaix/lit-reader/api/chapter"
	httpcommon "github.com/arkhaix/lit-reader/internal/servers/http/common"
)

var (
	// Client is the gRPC client for communicating with the story service.
	// Set this before using the handlers
	Client api.ChapterServiceClient

	// Timeout is the gRPC timeout.
	// Set this before using the handlers
	Timeout time.Duration
)

// GetChapter returns the requested chapter, scraping it first if necessary.
func GetChapter(w http.ResponseWriter, r *http.Request) {
	// Parse params
	storyID := chi.URLParam(r, "storyId")
	chapterIDStr := chi.URLParam(r, "chapterId")
	chapterID, err := strconv.Atoi(chapterIDStr)
	if err != nil {
		http.Error(w, "Invalid chapter id", http.StatusBadRequest)
	}
	log.Debugf("In: GetChapter(%s, %d)", storyID, chapterID)

	// RPC
	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()
	log.Debugf("RpcOut: GetChapter(%s, %d)", storyID, chapterID)
	result, err := Client.GetChapter(ctx, &api.GetChapterRequest{
		StoryId:   storyID,
		ChapterId: int32(chapterID),
	})

	if err != nil {
		log.Errorf("RpcFail: GetChapter(%s, %d): %s", storyID, chapterID, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Debugf("RpcGood: GetChapter(%s): %d", storyID, chapterID, result.GetStatus().GetCode())

	// Output
	response := newGetChapterResponse(result)
	log.Debugf("Out: GetChapter(%s, %d): %d", storyID, chapterID, response.Status.Code)
	render.Render(w, r, response)
}

type chapterResponse struct {
	Status    httpcommon.Status `json:"Status"`
	StoryID   string            `json:"StoryId"`
	ChapterID int               `json:"ChapterId"`
	URL       string            `json:"Url"`
	Title     string            `json:"Title"`
	HTML      string            `json:"Html"`
}

func (chapterResponse) Render(http.ResponseWriter, *http.Request) error {
	return nil
}

func newGetChapterResponse(pb *api.GetChapterResponse) *chapterResponse {
	return &chapterResponse{
		Status:    httpcommon.NewStatusFromProto(pb.GetStatus()),
		StoryID:   pb.GetData().GetStoryId(),
		ChapterID: int(pb.GetData().GetChapterId()),
		URL:       pb.GetData().GetUrl(),
		Title:     pb.GetData().GetTitle(),
		HTML:      pb.GetData().GetHtml(),
	}
}
