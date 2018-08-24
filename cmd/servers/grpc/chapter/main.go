package main

import (
	"database/sql"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"

	"github.com/arkhaix/lit-reader/common"
	server "github.com/arkhaix/lit-reader/internal/servers/grpc/chapter"

	apichapter "github.com/arkhaix/lit-reader/api/chapter"
	apiscraper "github.com/arkhaix/lit-reader/api/scraper"
	apistory "github.com/arkhaix/lit-reader/api/story"
)

func main() {
	// Config
	listenPort := common.ConfigVar("3002", "CHAPTER_GRPC_SERVICE_PORT", nil)
	scraperHost := common.ConfigVar("localhost", "SCRAPER_GRPC_SERVICE_HOSTNAME", nil)
	scraperPort := common.ConfigVar("3000", "SCRAPER_GRPC_SERVICE_PORT", nil)
	storyHost := common.ConfigVar("localhost", "STORY_GRPC_SERVICE_HOSTNAME", nil)
	storyPort := common.ConfigVar("3001", "STORY_GRPC_SERVICE_PORT", nil)
	dbString := common.ConfigVar("postgresql://chapter_service@roach:26257/reader?sslmode=disable", "CHAPTER_DB_STRING", nil)

	// Connect to scraper
	scraperAddress := scraperHost + ":" + scraperPort
	log.Infof("Connecting to scraper host at %s", scraperAddress)
	scraperConn, err := grpc.Dial(scraperAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to scraper service: %v", err)
	}
	defer scraperConn.Close()

	// Connect to story
	storyAddress := storyHost + ":" + storyPort
	log.Infof("Connecting to story host at %s", storyAddress)
	storyConn, err := grpc.Dial(storyAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to story service: %v", err)
	}
	defer storyConn.Close()

	// Connect to db
	log.Infof("Connecting to database at %s", dbString)
	db, err := sql.Open("postgres", dbString)
	if err != nil {
		log.Fatal("Error connecting to the database")
	}

	// Config server
	chapterServer := server.Server{
		ScraperClient:  apiscraper.NewScraperServiceClient(scraperConn),
		ScraperTimeout: 10 * time.Second,
		StoryClient:    apistory.NewStoryServiceClient(storyConn),
		StoryTimeout:   10 * time.Second,
		DB:             db,
	}

	// Listen
	listenPort = ":" + listenPort
	lis, err := net.Listen("tcp", listenPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Serve
	s := grpc.NewServer()
	apichapter.RegisterChapterServiceServer(s, &chapterServer)
	reflection.Register(s)
	log.Info("Serving grpc on", lis.Addr().String())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
