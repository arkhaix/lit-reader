package main

import (
	"bytes"
	"database/sql"
	"net"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags/logrus"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"

	"github.com/arkhaix/lit-reader/common"
	server "github.com/arkhaix/lit-reader/internal/servers/grpc/chapter"

	apichapter "github.com/arkhaix/lit-reader/api/chapter"
	apiscraper "github.com/arkhaix/lit-reader/api/scraper"
	apistory "github.com/arkhaix/lit-reader/api/story"
)

// JSONPbMarshaller is the marshaller used for serializing protobuf messages.
var JSONPbMarshaller = &jsonpb.Marshaler{}

var logLevels = map[string]log.Level{
	"debug": log.DebugLevel,
	"info":  log.InfoLevel,
	"warn":  log.WarnLevel,
	"error": log.ErrorLevel,
	"panic": log.PanicLevel,
}

func main() {
	// Service config
	listenPort := common.ConfigVar("3000", "CHAPTER_GRPC_SERVICE_PORT", nil)
	scraperHost := common.ConfigVar("localhost", "SCRAPER_GRPC_SERVICE_HOSTNAME", nil)
	scraperPort := common.ConfigVar("3000", "SCRAPER_GRPC_SERVICE_PORT", nil)
	storyHost := common.ConfigVar("localhost", "STORY_GRPC_SERVICE_HOSTNAME", nil)
	storyPort := common.ConfigVar("3000", "STORY_GRPC_SERVICE_PORT", nil)
	dbString := common.ConfigVar("postgresql://chapter_service@roach:26257/reader?sslmode=disable", "CHAPTER_DB_STRING", nil)

	// Logging config
	logFormat := common.ConfigVar("text", "LOG_FORMAT", nil)
	logLevel := common.ConfigVar("info", "LOG_LEVEL", nil)

	// Logger setup
	logEntry := log.NewEntry(log.StandardLogger())
	grpc_logrus.ReplaceGrpcLogger(logEntry)
	if logFormat == "json" {
		log.SetFormatter(&log.JSONFormatter{})
	}
	log.SetLevel(logLevels[logLevel])

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

	// gRPC middleware
	s := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			// grpc_prometheus.StreamServerInterceptor,
			grpc_logrus.StreamServerInterceptor(logEntry),
			streamLogPayloadOnError(logEntry, decider),
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			// grpc_prometheus.UnaryServerInterceptor,
			grpc_logrus.UnaryServerInterceptor(logEntry),
			unaryLogPayloadOnError(logEntry, decider),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)

	// Serve
	apichapter.RegisterChapterServiceServer(s, &chapterServer)
	reflection.Register(s)
	log.Info("Serving grpc on", lis.Addr().String())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func unaryLogPayloadOnError(entry *log.Entry, decider grpc_logging.ServerPayloadLoggingDecider) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		result, err := handler(ctx, req)

		statusOk := true
		if response, ok := result.(*apichapter.GetChapterResponse); ok {
			if response.GetStatus().GetCode() != 200 {
				statusOk = false
			}
		} else {
			log.Errorf("Failed to cast result to GetChapterResponse: %v", result)
		}

		if decider(ctx, info.FullMethod, info.Server) && (err != nil || statusOk == false) {
			// Use the provided logrus.Entry for logging but use the fields from context.
			logEntry := entry.WithFields(ctx_logrus.Extract(ctx).Data)
			logProtoMessageAsJSON(logEntry, req, "grpc.request.content", "server request payload logged as grpc.request.content field")
			if err == nil {
				logProtoMessageAsJSON(logEntry, result, "grpc.response.content", "server response payload logged as grpc.request.content field")
			}
		}

		return result, err
	}
}

type loggingServerStream struct {
	grpc.ServerStream
	entry *log.Entry
}

func (l *loggingServerStream) SendMsg(m interface{}) error {
	err := l.ServerStream.SendMsg(m)
	if err == nil {
		logProtoMessageAsJSON(l.entry, m, "grpc.response.content", "server response payload logged as grpc.response.content field")
	}
	return err
}

func (l *loggingServerStream) RecvMsg(m interface{}) error {
	err := l.ServerStream.RecvMsg(m)
	if err == nil {
		logProtoMessageAsJSON(l.entry, m, "grpc.request.content", "server request payload logged as grpc.request.content field")
	}
	return err
}

func streamLogPayloadOnError(entry *log.Entry, decider grpc_logging.ServerPayloadLoggingDecider) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		if !decider(stream.Context(), info.FullMethod, srv) {
			return handler(srv, stream)
		}
		// Use the provided logrus.Entry for logging but use the fields from context.
		logEntry := entry.WithFields(grpc_logrus.Extract(stream.Context()).Data)
		newStream := &loggingServerStream{ServerStream: stream, entry: logEntry}
		return handler(srv, newStream)
	}
}

func logProtoMessageAsJSON(entry *log.Entry, pbMsg interface{}, key string, msg string) {
	if p, ok := pbMsg.(proto.Message); ok {
		entry.WithField(key, &jsonpbMarshallable{p}).Info(msg)
	}
}

func decider(ctx context.Context, fullMethodName string, servingObject interface{}) bool {
	return true
}

type jsonpbMarshallable struct {
	proto.Message
}

func (j *jsonpbMarshallable) MarshalJSON() ([]byte, error) {
	b := &bytes.Buffer{}
	if err := JSONPbMarshaller.Marshal(b, j.Message); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
