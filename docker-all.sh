#! /bin/bash

# Do the first one by itself to build common layers
docker build -t chapter-grpc -f ./build/chapter-grpc.Dockerfile .

# Then the rest in parallel since the only real difference should be the build target
docker build -t chapter-http -f ./build/chapter-http.Dockerfile . &
docker build -t scraper-grpc -f ./build/scraper-grpc.Dockerfile . &
docker build -t story-grpc -f ./build/story-grpc.Dockerfile . &
docker build -t story-http -f ./build/story-http.Dockerfile . &
docker build -t reader -f ./build/reader.Dockerfile . &
docker build -t cockroach-init -f ./build/cockroach-init.Dockerfile . &
wait $(jobs -p)

while getopts "g" o; do
  case "${o}" in 
    g)
      # Tag and push for gke
      echo 'tagging for gke'
      docker tag chapter-grpc gcr.io/arkhaix-lit-reader/chapter-grpc &
      docker tag chapter-http gcr.io/arkhaix-lit-reader/chapter-http &
      docker tag scraper-grpc gcr.io/arkhaix-lit-reader/scraper-grpc &
      docker tag story-grpc gcr.io/arkhaix-lit-reader/story-grpc &
      docker tag story-http gcr.io/arkhaix-lit-reader/story-http &
      docker tag reader gcr.io/arkhaix-lit-reader/reader &
      docker tag cockroach-init gcr.io/arkhaix-lit-reader/cockroach-init &
      wait $(jobs -p)
      echo 'finished tagging for gke'

      echo 'pushing to gke'
      docker push gcr.io/arkhaix-lit-reader/chapter-grpc
      docker push gcr.io/arkhaix-lit-reader/chapter-http &
      docker push gcr.io/arkhaix-lit-reader/scraper-grpc &
      docker push gcr.io/arkhaix-lit-reader/story-grpc &
      docker push gcr.io/arkhaix-lit-reader/story-http &
      docker push gcr.io/arkhaix-lit-reader/reader &
      docker push gcr.io/arkhaix-lit-reader/cockroach-init &
      wait $(jobs -p)
      echo 'finished pushing to gke'
      ;;
  esac
done
