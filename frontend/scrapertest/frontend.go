package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	_ "github.com/davecgh/go-spew/spew"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	api "github.com/arkhaix/lit-reader/api/scraper"
)

const (
	address    = "localhost:50051"
	defaultURL = "wanderinginn.com"
	timeout    = 10 * time.Second
)

var (
	client api.ScraperClient
)

func main() {
	// Set up gRPC client
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client = api.NewScraperClient(conn)

	// Route http
	r := chi.NewRouter()
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/story", func(r chi.Router) {
		r.Get("/search", storySearch)
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}

func storySearch(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "invalid method", http.StatusBadRequest)
		return
	}

	storyURL := r.URL.Query().Get("url")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// result, err := client.CheckStoryURL(ctx, &api.CheckStoryURLRequest{Url: storyURL})
	result, err := client.FetchStoryMetadata(ctx, &api.FetchStoryMetadataRequest{
		Url: storyURL,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// w.Write([]byte(fmt.Sprintf("{\"allowed\":%v}", result.GetAllowed())))
	// w.Write([]byte(spew.Sdump(result)))
	w.Write([]byte(fmt.Sprintf("url: '%s'\ntitle: '%s'\nauthor: '%s'\nchapters: %d",
		result.GetUrl(), result.GetTitle(), result.GetAuthor(), result.GetNumChapters())))
}
