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
	api "github.com/arkhaix/lit-reader/api/story"

	// http handlers
	"github.com/arkhaix/lit-reader/internal/servers/http/story"
)

const (
	envVarListenPort = "STORY_HTTP_SERVICE_PORT"

	envVarGRPCHostName = "STORY_GRPC_SERVICE_HOSTNAME"
	envVarGRPCPort     = "STORY_GRPC_SERVICE_PORT"
)

var (
	listenPort = "8080"

	grpcHostName = "localhost"
	grpcPort     = "3000"
)

func main() {
	log.Info("=====")
	log.Info("Environment")
	envVars := os.Environ()
	for _, s := range envVars {
		log.Info(s)
	}
	log.Info("=====")

	// Read environment config
	if envGRPCHostName, ok := os.LookupEnv(envVarGRPCHostName); ok {
		grpcHostName = envGRPCHostName
	}
	if envGRPCPort, ok := os.LookupEnv(envVarGRPCPort); ok {
		grpcPort = envGRPCPort
	}
	if envListenPort, ok := os.LookupEnv(envVarListenPort); ok {
		listenPort = envListenPort
	}

	// Set up gRPC client
	grpcAddress := grpcHostName + ":" + grpcPort
	log.Infof("Connecting to grpc host at %s", grpcAddress)
	conn, err := grpc.Dial(grpcAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Set up handlers
	story.Client = api.NewStoryServiceClient(conn)
	story.Timeout = 10 * time.Second

	// Routes
	r := chi.NewRouter()
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/story", func(r chi.Router) {
		r.Get("/{storyID}", story.GetStory)
		r.Post("/", story.PostStory)
	})

	docgen.PrintRoutes(r)

	// Listen
	listenPort = ":" + listenPort
	log.Infof("Listening for http on %s", listenPort)
	log.Fatal(http.ListenAndServe(listenPort, r))
}
