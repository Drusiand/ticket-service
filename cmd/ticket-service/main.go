package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Drusiand/ticket-service/internal/config"
	"github.com/Drusiand/ticket-service/internal/httpserver"
	"github.com/Drusiand/ticket-service/internal/logger"
)

func main() {
	os.Exit(run())
}

func run() int {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	baseLog := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	cfg, err := config.New()
	if err != nil {
		baseLog.Error("failed to load config", slog.Any("error", err))
		return 1
	}

	log, err := logger.New(cfg.AppLogLevel, cfg.AppLogHandler, os.Stdout)
	if err != nil {
		baseLog.Error("failed to create logger", slog.Any("error", err))
		return 1
	}
	log = log.With("service", "ticket-service")
	slog.SetDefault(log)

	log.Info("starting HTTP server", slog.Int("port", cfg.AppHTTPPort))
	server := httpserver.New(cfg, log.With("package", "httpserver"))

	errCh := make(chan error, 1)
	go func() {
		err := server.Run()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
	}()

	select {
	case err := <-errCh:
		log.Error("unexpectedly stopping HTTP server", slog.Any("error", err))
		return 1
	case <-ctx.Done():
	}

	log.Info("stopping HTTP server")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), cfg.AppShutdownTimeout)
	defer shutdownCancel()

	err = server.Stop(shutdownCtx)
	if err != nil {
		log.Error("error during server shutdown", slog.Any("error", err))
		return 1
	}
	log.Info("stopped HTTP server")

	return 0
}
