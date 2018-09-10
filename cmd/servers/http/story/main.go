package main

import (
	"time"

	"google.golang.org/grpc"

	"github.com/go-chi/chi"

	api "github.com/arkhaix/lit-reader/api/story"

	"github.com/arkhaix/lit-reader/internal/servers/http/story"
	"github.com/arkhaix/lit-reader/internal/servers/httphost"
)

// hostApp implements httphost.HostApp
type hostApp struct{}

func (*hostApp) GetName() string {
	return "story"
}

func (*hostApp) GetParams() *httphost.Params {
	return &httphost.Params{
		EnvVarListenPort:  "STORY_HTTP_SERVICE_PORT",
		DefaultListenPort: "8080",

		EnvVarGRPCHostName:  "STORY_GRPC_SERVICE_HOSTNAME",
		DefaultGRPCHostName: "localhost",

		EnvVarGRPCPort:  "STORY_GRPC_SERVICE_PORT",
		DefaultGRPCPort: "3000",
	}
}

func (*hostApp) DefineRoutes(r *chi.Mux) {
	r.Route("/story", func(r chi.Router) {
		r.Get("/{storyId}", story.GetStory)
		r.Post("/", story.PostStory)
	})
}

func (*hostApp) ConfigureHandlers(conn *grpc.ClientConn) {
	story.Client = api.NewStoryServiceClient(conn)
	story.Timeout = 10 * time.Second
}

func main() {
	app := &hostApp{}
	httphost.Host(app)
}
