# Build stage
FROM golang:1.11 AS builder

ENV GOPATH /go
ADD . /go/src/github.com/arkhaix/lit-reader
WORKDIR /go/src/github.com/arkhaix/lit-reader

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./scraper-grpc ./cmd/servers/grpc/scraper

# Result stage
FROM alpine
RUN apk add --no-cache ca-certificates 
WORKDIR /app
COPY --from=builder /go/src/github.com/arkhaix/lit-reader/scraper-grpc /app/
EXPOSE 3000
ENTRYPOINT ["/app/scraper-grpc"]