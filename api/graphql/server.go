package graphql

import (
	"net"
	"net/http"

	rice "github.com/GeertJohan/go.rice"

	"github.com/dfuse-io/solana-go/api/graphql/static"

	"github.com/gorilla/handlers"

	"github.com/dfuse-io/derr"
	"github.com/dfuse-io/solana-go/api/graphql/apollo"
	"github.com/dfuse-io/solana-go/api/graphql/resolvers"
	"github.com/gorilla/mux"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"go.uber.org/zap"
)

type Server struct {
	schemaFile  string
	servingAddr string
}

func NewServer(schemaFile string, servingAddr string) *Server {
	return &Server{
		schemaFile:  schemaFile,
		servingAddr: servingAddr,
	}
}

func (s *Server) Launch() error {
	// initialize GraphQL
	box := rice.MustFindBox("build")

	resolver := resolvers.NewRoot()
	schema, err := graphql.ParseSchema(
		string(box.MustBytes("schema.graphql")),
		resolver,
		graphql.UseFieldResolvers(),
		graphql.UseStringDescriptions(),
	)
	derr.Check("parse schema", err)
	return StartHTTPServer(s.servingAddr, schema)

}

func StartHTTPServer(listenAddr string, schema *graphql.Schema) error {
	router := mux.NewRouter()
	router.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		if derr.IsShuttingDown() {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		w.Write([]byte("ok"))
	})

	staticRouter := router.PathPrefix("/").Subrouter()
	static.RegisterStaticRoutes(staticRouter)

	restRouter := router.PathPrefix("/").Subrouter()
	restRouter.Use(apollo.NewMiddleware(schema).Handler)
	restRouter.Handle("/graphql", &relay.Handler{Schema: schema})

	// http
	httpListener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		zlog.Panic("http listen failed", zap.String("http_addr", listenAddr), zap.Error(err))
	}

	corsMiddleware := NewCORSMiddleware()
	httpServer := http.Server{
		Handler: corsMiddleware(router),
	}

	zlog.Info("serving HTTP", zap.String("http_addr", listenAddr))
	return httpServer.Serve(httpListener)
}

func NewCORSMiddleware() mux.MiddlewareFunc {
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "X-Eos-Push-Guarantee"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "OPTIONS"})
	maxAge := handlers.MaxAge(86400) // 24 hours - hard capped by Firefox / Chrome is max 10 minutes

	return handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods, maxAge)
}
