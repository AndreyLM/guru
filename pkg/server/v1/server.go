package v1

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"

	"github.com/gorilla/handlers"
)

// Server - app server
type Server struct {
	port   string
	router *mux.Router
}

// NewServer - creates server
func NewServer(port string) Server {
	return Server{
		port:   port,
		router: mux.NewRouter().StrictSlash(true),
	}
}

// Start - start server
func (s *Server) Start() {
	log.Println("Starting server on port", s.port)
	s.init()

	handler := handlers.LoggingHandler(os.Stdout, handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Origin", "Cache-Control", "X-App-Token"}),
		handlers.ExposedHeaders([]string{"*"}),
		handlers.MaxAge(1000),
		handlers.AllowCredentials(),
	)(s.router))
	handler = handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))(handler)

	newServer := &http.Server{
		Handler:      handler,
		Addr:         "0.0.0.0:" + s.port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if err := newServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func (s *Server) init() {
	rootRoutes := getRootRoutes()

	for _, route := range rootRoutes {
		s.router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
}
