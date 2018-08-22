package main

import (
	"net/http"
	"os"
	"time"

	"google.golang.org/grpc"

	"github.com/go-chi/chi"
	"github.com/go-chi/docgen"
	"github.com/go-chi/render"

	log "github.com/sirupsen/logrus"

	// grpc
	api "github.com/arkhaix/lit-reader/api/scraper"

	// http handlers
	"github.com/arkhaix/lit-reader/internal/servers/http/storytest"
	"github.com/arkhaix/lit-reader/internal/servers/http/testpage"
)

var (
	scraperHostName = "localhost"
	scraperPort     = "3000"
)

func main() {
	log.Info("=====")
	log.Info("Environment")
	envVars := os.Environ()
	for _, s := range envVars {
		log.Info(s)
	}
	log.Info("=====")

	// Determine scraper service address
	if envScraperHostName, ok := os.LookupEnv("SCRAPER_GRPC_SERVICE_HOSTNAME"); ok {
		scraperHostName = envScraperHostName
	}
	if envScraperPort, ok := os.LookupEnv("SCRAPER_GRPC_SERVICE_PORT"); ok {
		scraperPort = envScraperPort
	}
	address := scraperHostName + ":" + scraperPort
	log.Infof("Connecting to scraper host at %s", address)

	// Set up gRPC client
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Set up handlers
	storytest.ScraperClient = api.NewScraperServiceClient(conn)
	storytest.ScraperTimeout = 10 * time.Second

	// Routes
	r := chi.NewRouter()
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/", func(r chi.Router) {
		r.Get("/", testpage.GetTestIndex)
	})

	r.Route("/story", func(r chi.Router) {
		r.Get("/{storyID}", storytest.GetStoryByID)
		r.Post("/", storytest.PostStoryByURL)

		r.Route("/{storyID}/chapter", func(r chi.Router) {
			r.Get("/{chapterID}", storytest.GetChapterByID)
		})
	})

	docgen.PrintRoutes(r)

	// Listen
	log.Info("Listening...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
