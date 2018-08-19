package scrapertest

import (
	"log"
	"net/http"
	"time"

	"google.golang.org/grpc"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	api "github.com/arkhaix/lit-reader/api/scraper"
	"github.com/arkhaix/lit-reader/frontend/scrapertest/handlers"
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
	handlers.ScraperClient = api.NewScraperClient(conn)
	handlers.ScraperTimeout = 10 * time.Second

	// Route http
	r := chi.NewRouter()
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/story", func(r chi.Router) {
		r.Get("/search", handlers.StorySearch)
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}
