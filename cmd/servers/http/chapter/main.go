package main

import (
	"time"

	"google.golang.org/grpc"

	"github.com/go-chi/chi"

	api "github.com/arkhaix/lit-reader/api/chapter"

	"github.com/arkhaix/lit-reader/internal/servers/http/chapter"
	"github.com/arkhaix/lit-reader/internal/servers/httphost"
)

// hostApp implements httphost.HostApp
type hostApp struct{}

func (*hostApp) GetParams() *httphost.Params {
	return &httphost.Params{
		EnvVarListenPort:  "CHAPTER_HTTP_SERVICE_PORT",
		DefaultListenPort: "8080",

		EnvVarGRPCHostName:  "CHAPTER_GRPC_SERVICE_HOSTNAME",
		DefaultGRPCHostName: "localhost",

		EnvVarGRPCPort:  "CHAPTER_GRPC_SERVICE_PORT",
		DefaultGRPCPort: "3000",
	}
}

func (*hostApp) DefineRoutes(r *chi.Mux) {
	r.Route("/story/{storyID}/chapter", func(r chi.Router) {
		r.Get("/{chapterID}", chapter.GetChapter)
	})
}

func (*hostApp) ConfigureHandlers(conn *grpc.ClientConn) {
	chapter.Client = api.NewChapterServiceClient(conn)
	chapter.Timeout = 10 * time.Second
}

func main() {
	app := &hostApp{}
	httphost.Host(app)
}
