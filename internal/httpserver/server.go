package httpserver

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/Drusiand/ticket-service/internal/config"
	"github.com/gorilla/mux"
)

type Server struct {
	server *http.Server
	log    *slog.Logger
	cfg    *config.Config
}

func New(cfg *config.Config, logger *slog.Logger) *Server {
	mux := mux.NewRouter()
	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "OK")
	})

	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.AppHTTPPort),
		Handler:           mux,
		ReadTimeout:       cfg.AppReadTimeout,
		ReadHeaderTimeout: cfg.AppReadHeaderTimeout,
		WriteTimeout:      cfg.AppWriteTimeout,
	}

	return &Server{
		server: server,
		log:    logger,
		cfg:    cfg,
	}
}

func (srv *Server) Run() error {
	if err := srv.server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func (srv *Server) Stop(ctx context.Context) error {
	if err := srv.server.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}
