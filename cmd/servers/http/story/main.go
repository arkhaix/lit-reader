package main

import (
	"time"

	"google.golang.org/grpc"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	api "github.com/arkhaix/lit-reader/api/story"

	"github.com/arkhaix/lit-reader/internal/servers/http/story"
	"github.com/arkhaix/lit-reader/internal/servers/httphost"
)

func configHandlers(conn *grpc.ClientConn) {
	story.Client = api.NewStoryServiceClient(conn)
	story.Timeout = 10 * time.Second
}

func main() {
	// Config
	params := httphost.Params{
		EnvVarListenPort:  "STORY_HTTP_SERVICE_PORT",
		DefaultListenPort: "8080",

		EnvVarGRPCHostName:  "STORY_GRPC_SERVICE_HOSTNAME",
		DefaultGRPCHostName: "localhost",

		EnvVarGRPCPort:  "STORY_GRPC_SERVICE_PORT",
		DefaultGRPCPort: "3000",
	}

	// Routes
	r := chi.NewRouter()
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/story", func(r chi.Router) {
		r.Get("/{storyID}", story.GetStory)
		r.Post("/", story.PostStory)
	})

	// Run
	httphost.Host(params, r, configHandlers)
}
