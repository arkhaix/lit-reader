# Build stage
FROM golang:1.10 AS builder

ENV GOPATH /go
ADD . /go/src/github.com/arkhaix/lit-reader
WORKDIR /go/src/github.com/arkhaix/lit-reader

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./chapter-http ./cmd/servers/http/chapter

# Result stage
FROM alpine
RUN apk add --no-cache ca-certificates 
WORKDIR /app
COPY --from=builder /go/src/github.com/arkhaix/lit-reader/chapter-http /app/
EXPOSE 8082
ENTRYPOINT ["/app/chapter-http"]