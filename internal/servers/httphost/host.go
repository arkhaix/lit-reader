package httphost

import (
	"net/http"
	"os"
	"time"

	"google.golang.org/grpc"

	"github.com/go-chi/chi"
	"github.com/go-chi/docgen"

	log "github.com/sirupsen/logrus"

	// grpc
	api "github.com/arkhaix/lit-reader/api/story"

	// http handlers
	"github.com/arkhaix/lit-reader/internal/servers/http/story"
)

// Params stuff
type Params struct {
	EnvVarListenPort  string
	DefaultListenPort string

	EnvVarGRPCHostName  string
	DefaultGRPCHostName string

	EnvVarGRPCPort  string
	DefaultGRPCPort string
}

// OnConnectFunc is invoked after a grpc connection is established
type OnConnectFunc func(*grpc.ClientConn)

// Host runs the grpc client and http listener
func Host(params Params, router *chi.Mux, onConnect OnConnectFunc) {
	// Debug
	log.Debug("=====")
	log.Debug("Environment")
	envVars := os.Environ()
	for _, s := range envVars {
		log.Debug(s)
	}
	log.Debug("=====")

	// Set defaults
	listenPort := params.DefaultListenPort
	grpcHostName := params.DefaultGRPCHostName
	grpcPort := params.DefaultGRPCPort

	// Read environment config
	if envListenPort, ok := os.LookupEnv(params.EnvVarListenPort); ok {
		listenPort = envListenPort
	}
	if envGRPCHostName, ok := os.LookupEnv(params.EnvVarGRPCHostName); ok {
		grpcHostName = envGRPCHostName
	}
	if envGRPCPort, ok := os.LookupEnv(params.EnvVarGRPCPort); ok {
		grpcPort = envGRPCPort
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
	docgen.PrintRoutes(router)

	// Listen
	listenPort = ":" + listenPort
	log.Infof("Listening for http on %s", listenPort)
	log.Fatal(http.ListenAndServe(listenPort, router))
}
