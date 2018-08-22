package main

import (
	"net"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	log "github.com/sirupsen/logrus"

	server "github.com/arkhaix/lit-reader/internal/servers/grpc/story"

	apiscraper "github.com/arkhaix/lit-reader/api/scraper"
	apistory "github.com/arkhaix/lit-reader/api/story"
)

var (
	listenPort = "3001"

	scraperHost = "localhost"
	scraperPort = "3000"
)

func main() {
	log.Info("=====")
	log.Info("Environment")
	envVars := os.Environ()
	for _, s := range envVars {
		log.Info(s)
	}
	log.Info("=====")

	// Environment config
	if envPort, ok := os.LookupEnv("STORY_GRPC_SERVICE_PORT"); ok {
		listenPort = envPort
	}
	if envScraperHost, ok := os.LookupEnv("SCRAPER_GRPC_SERVICE_HOSTNAME"); ok {
		// prefer host name
		scraperHost = envScraperHost
	} else if envScraperHost, ok := os.LookupEnv("SCRAPER_GRPC_SERVICE_HOST"); ok {
		// then host ip
		scraperHost = envScraperHost
	}
	if envScraperPort, ok := os.LookupEnv("SCRAPER_GRPC_SERVICE_PORT"); ok {
		scraperPort = envScraperPort
	}

	// Connect to scraper
	scraperAddress := scraperHost + ":" + scraperPort
	log.Infof("Connecting to scraper host at %s", scraperAddress)
	scraperConn, err := grpc.Dial(scraperAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to scraper: %v", err)
	}
	defer scraperConn.Close()

	// Config server
	storyServer := server.Server{
		ScraperClient:  apiscraper.NewScraperServiceClient(scraperConn),
		ScraperTimeout: 10 * time.Second,
	}

	// Listen
	listenPort = ":" + listenPort
	lis, err := net.Listen("tcp", listenPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Serve
	s := grpc.NewServer()
	apistory.RegisterStoryServiceServer(s, &storyServer)
	reflection.Register(s)
	log.Info("Serving grpc on", lis.Addr().String())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
