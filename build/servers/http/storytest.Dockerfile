# Build stage
FROM golang:1.10 AS builder

ENV GOPATH /go
ADD . /go/src/github.com/arkhaix/lit-reader
WORKDIR /go/src/github.com/arkhaix/lit-reader

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./storytest_http ./cmd/servers/http/storytest

# Result stage
FROM alpine
RUN apk add --no-cache ca-certificates 
WORKDIR /app
COPY --from=builder /go/src/github.com/arkhaix/lit-reader/storytest_http /app/
EXPOSE 8080
ENTRYPOINT ["/app/storytest_http"]