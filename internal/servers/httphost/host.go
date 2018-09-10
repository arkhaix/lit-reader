package httphost

import (
	"net/http"
	"os"

	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/docgen"
	"github.com/go-chi/render"

	log "github.com/sirupsen/logrus"

	chiprometheus "github.com/766b/chi-prometheus"
	"github.com/prometheus/client_golang/prometheus"
)

// HostApp is the interface each /cmd/ app must implement
type HostApp interface {
	GetName() string

	// GetParams must return the httphost.Params configuration
	GetParams() *Params

	// DefineRoutes should define all desired http routes
	DefineRoutes(r *chi.Mux)

	// ConfigureHandlers should perform any post-connect setup needed by the handlers
	ConfigureHandlers(conn *grpc.ClientConn)
}

// Params defines the configurable host app parameters
type Params struct {
	EnvVarListenPort  string
	DefaultListenPort string

	EnvVarGRPCHostName  string
	DefaultGRPCHostName string

	EnvVarGRPCPort  string
	DefaultGRPCPort string
}

// Host runs the grpc client and http listener
func Host(app HostApp) {
	debugPrintEnv()

	// Set up gRPC client
	grpcAddress, listenPort := getGRPCConfig(app)
	log.Infof("Connecting to grpc host at %s", grpcAddress)
	conn, err := grpc.Dial(grpcAddress, grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_prometheus.UnaryClientInterceptor),
		grpc.WithStreamInterceptor(grpc_prometheus.StreamClientInterceptor))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Set up handlers
	app.ConfigureHandlers(conn)

	// Set up router
	r := setUpRouter(app)

	// Serve
	listenPort = ":" + listenPort
	log.Infof("Listening for http on %s", listenPort)
	log.Fatal(http.ListenAndServe(listenPort, r))
}

func debugPrintEnv() {
	log.Debug("=====")
	log.Debug("Environment")
	envVars := os.Environ()
	for _, s := range envVars {
		log.Debug(s)
	}
	log.Debug("=====")
}

func getGRPCConfig(app HostApp) (grpcAddress string, listenPort string) {
	params := app.GetParams()

	listenPort = params.DefaultListenPort
	grpcHostName := params.DefaultGRPCHostName
	grpcPort := params.DefaultGRPCPort

	// Read environment config
	if envListenPort, ok := os.LookupEnv(params.EnvVarListenPort); ok {
		listenPort = envListenPort
	}
	if envGRPCHostName, ok := os.LookupEnv(params.EnvVarGRPCHostName); ok {
		grpcHostName = envGRPCHostName
	}
	if envGRPCPort, ok := os.LookupEnv(params.EnvVarGRPCPort); ok {
		grpcPort = envGRPCPort
	}

	grpcAddress = grpcHostName + ":" + grpcPort

	return /*grpcAddress, listenPort*/
}

func setUpRouter(app HostApp) *chi.Mux {
	r := chi.NewRouter()

	// Prometheus
	m := chiprometheus.NewMiddleware(app.GetName())
	r.Use(m)

	// JSON
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// CORS
	cors := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		// alternatively, AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

	// Routes must be defined after all middlewares
	r.Handle("/metrics", prometheus.Handler())
	app.DefineRoutes(r)
	docgen.PrintRoutes(r)

	return r
}
