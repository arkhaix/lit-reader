#! /bin/bash

docker build -t scraper-grpc -f ./build/servers/grpc/scraper.Dockerfile . &
docker build -t story-grpc -f ./build/servers/grpc/story.Dockerfile . &
docker build -t story-http -f ./build/servers/http/story.Dockerfile . &
docker build -t storytest-http -f ./build/servers/http/storytest.Dockerfile . &
wait $(jobs -p)

docker tag scraper-grpc gcr.io/arkhaix-lit-reader/scraper-grpc
docker tag story-grpc gcr.io/arkhaix-lit-reader/story-grpc
docker tag story-http gcr.io/arkhaix-lit-reader/story-http
docker tag storytest-http gcr.io/arkhaix-lit-reader/storytest-http

docker push gcr.io/arkhaix-lit-reader/scraper-grpc
docker push gcr.io/arkhaix-lit-reader/story-grpc
docker push gcr.io/arkhaix-lit-reader/story-http
docker push gcr.io/arkhaix-lit-reader/storytest-http