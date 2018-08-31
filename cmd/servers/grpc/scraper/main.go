package main

import (
	"net"
	"os"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	log "github.com/sirupsen/logrus"

	api "github.com/arkhaix/lit-reader/api/scraper"
	server "github.com/arkhaix/lit-reader/internal/servers/grpc/scraper"
)

var (
	port = "3000"
)

func main() {
	log.Info("=====")
	log.Info("Environment")
	envVars := os.Environ()
	for _, s := range envVars {
		log.Info(s)
	}
	log.Info("=====")

	if envPort, ok := os.LookupEnv("SCRAPER_GRPC_SERVICE_PORT"); ok {
		port = envPort
	}
	port = ":" + port

	lis, err := net.Listen("tcp", port)
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

	api.RegisterScraperServiceServer(s, &server.Server{})
	reflection.Register(s)
	log.Info("Serving grpc on", lis.Addr().String())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
