package main

import (
	"database/sql"
	"net"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"

	"github.com/arkhaix/lit-reader/common"
	server "github.com/arkhaix/lit-reader/internal/servers/grpc/story"

	apiscraper "github.com/arkhaix/lit-reader/api/scraper"
	apistory "github.com/arkhaix/lit-reader/api/story"
)

func main() {
	// Config
	listenPort := common.ConfigVar("3000", "STORY_GRPC_SERVICE_PORT", nil)
	scraperHost := common.ConfigVar("localhost", "SCRAPER_GRPC_SERVICE_HOSTNAME", nil)
	scraperPort := common.ConfigVar("3000", "SCRAPER_GRPC_SERVICE_PORT", nil)
	dbString := common.ConfigVar("postgresql://story_service@roach:26257/reader?sslmode=disable", "STORY_DB_STRING", nil)

	// Connect to scraper
	scraperAddress := scraperHost + ":" + scraperPort
	log.Infof("Connecting to scraper host at %s", scraperAddress)
	scraperConn, err := grpc.Dial(scraperAddress, grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_prometheus.UnaryClientInterceptor),
		grpc.WithStreamInterceptor(grpc_prometheus.StreamClientInterceptor))
	if err != nil {
		log.Fatalf("Failed to connect to scraper service: %v", err)
	}
	defer scraperConn.Close()

	// Connect to db
	log.Infof("Connecting to database at %s", dbString)
	db, err := sql.Open("postgres", dbString)
	if err != nil {
		log.Fatal("Error connecting to the database")
	}

	// Config server
	storyServer := server.Server{
		ScraperClient:  apiscraper.NewScraperServiceClient(scraperConn),
		ScraperTimeout: 10 * time.Second,
		DB:             db,
	}

	// Listen
	listenPort = ":" + listenPort
	lis, err := net.Listen("tcp", listenPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// gRPC middleware
	s := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_prometheus.StreamServerInterceptor,
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_prometheus.UnaryServerInterceptor,
		)),
	)

	// Serve prometheus metrics
	http.Handle("/metrics", promhttp.Handler())
	go func() { log.Debug(http.ListenAndServe(":8080", nil)) }()

	// Serve
	apistory.RegisterStoryServiceServer(s, &storyServer)
	reflection.Register(s)
	log.Info("Serving grpc on", lis.Addr().String())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
