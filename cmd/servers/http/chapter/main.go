package main

import (
	"time"

	"google.golang.org/grpc"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	api "github.com/arkhaix/lit-reader/api/chapter"

	"github.com/arkhaix/lit-reader/internal/servers/http/chapter"
	"github.com/arkhaix/lit-reader/internal/servers/httphost"
)

func configHandlers(conn *grpc.ClientConn) {
	chapter.Client = api.NewChapterServiceClient(conn)
	chapter.Timeout = 10 * time.Second
}

func main() {
	// Config
	params := httphost.Params{
		EnvVarListenPort:  "CHAPTER_HTTP_SERVICE_PORT",
		DefaultListenPort: "8080",

		EnvVarGRPCHostName:  "CHAPTER_GRPC_SERVICE_HOSTNAME",
		DefaultGRPCHostName: "localhost",

		EnvVarGRPCPort:  "CHAPTER_GRPC_SERVICE_PORT",
		DefaultGRPCPort: "3000",
	}

	// Routes
	r := chi.NewRouter()
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/story/{storyID}/chapter", func(r chi.Router) {
		r.Get("/{chapterID}", chapter.GetChapter)
	})

	// Run
	httphost.Host(params, r, configHandlers)
}
