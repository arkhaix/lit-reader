package main

import (
	"net/http"
	"time"

	"google.golang.org/grpc"

	"github.com/go-chi/chi"
	"github.com/go-chi/docgen"
	"github.com/go-chi/render"

	log "github.com/sirupsen/logrus"

	// grpc
	api "github.com/arkhaix/lit-reader/api/scraper"

	// http handlers
	"github.com/arkhaix/lit-reader/internal/servers/http/story"
	"github.com/arkhaix/lit-reader/internal/servers/http/testpage"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up gRPC client
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Set up handlers
	story.ScraperClient = api.NewScraperClient(conn)
	story.ScraperTimeout = 10 * time.Second

	// Routes
	r := chi.NewRouter()
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/", func(r chi.Router) {
		r.Get("/", testpage.GetTestIndex)
	})

	r.Route("/story", func(r chi.Router) {
		r.Get("/{storyID}", story.GetStoryByID)
		r.Post("/", story.PostStoryByURL)

		r.Route("/{storyID}/chapter", func(r chi.Router) {
			r.Get("/{chapterID}", story.GetChapterByID)
		})
	})

	docgen.PrintRoutes(r)

	// Listen
	log.Info("Listening...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
