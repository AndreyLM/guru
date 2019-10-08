package v1

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/andreylm/guru/pkg/cache"
	"github.com/andreylm/guru/pkg/db"

	"github.com/gorilla/mux"

	"github.com/gorilla/handlers"
)

// Server - app server
type Server struct {
	port   string
	router *mux.Router
	db     db.Storage
}

// NewServer - creates server
func NewServer(port string) Server {
	return Server{
		port:   port,
		router: mux.NewRouter().StrictSlash(true),
		db:     db.NewMockStorage(),
	}
}

// Start - start server
func (s *Server) Start() {
	log.Println("Starting server on port", s.port)
	s.init()

	handler := handlers.LoggingHandler(os.Stdout, handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Origin", "Cache-Control"}),
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

	ticker := time.NewTicker(10 * time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				updateUsers(s.db)
			}
		}
	}()

	if err := newServer.ListenAndServe(); err != nil {
		cancel()
		log.Println(err)
	}
}

func (s *Server) init() {
	rootRoutes := getRootRoutes(s.db)

	for _, route := range rootRoutes {
		s.router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
}

func updateUsers(dbStorage db.Storage) {
	for _, user := range cache.Storage.GetModifiedUsers() {
		log.Println("Saving user", user.ID)
		dbStorage.SaveUser(user)
	}

	cache.Storage.ClearModifiedUserCollection()
}
