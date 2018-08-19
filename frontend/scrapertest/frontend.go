package main

import (
	"log"
	"os"
	"time"

	api "github.com/arkhaix/lit-reader/api/scraper"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address    = "localhost:50051"
	defaultURL = "wanderinginn.com"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := api.NewScraperClient(conn)

	// Contact the server and print out its response.
	storyURL := defaultURL
	if len(os.Args) > 1 {
		storyURL = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.CheckStoryURL(ctx, &api.CheckStoryURLRequest{Url: storyURL})
	if err != nil {
		log.Fatalf("could not check: %v", err)
	}
	log.Printf("Result: %v", r.Allowed)
}
