package server

import (
	"context"
	"electronic-document-service/internal/config"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

const (
	shutdownTimeout     = 5 * time.Second
	readerHeaderTimeout = 5 * time.Second
)

type Server struct {
	*http.Server
	log *slog.Logger
}

func (s *Server) ConfigureRoutes() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("pong"))
	})

	s.Handler = mux
}

func New(log *slog.Logger, cfg config.Server) *Server {
	addr := net.JoinHostPort(cfg.Host, strconv.Itoa(cfg.Port))
	//nolint: exhaustruct // default options in http server is good
	srv := &http.Server{
		ReadHeaderTimeout: readerHeaderTimeout,
		Addr:              addr,
		ErrorLog:          slog.NewLogLogger(log.Handler(), slog.LevelError),
	}

	return &Server{
		Server: srv,
		log:    log,
	}
}

func (s *Server) Run(ctx context.Context) error {
	idleConnClosed := make(chan struct{})

	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-ctx.Done():
		case <-signals:
		}

		if err := s.Shutdown(ctx); err != nil {
			s.log.Error("graceful shutdown http server", slog.Any("error", err))
		}

		close(idleConnClosed)
	}()

	if err := s.ListenAndServe(); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			s.log.Info("http server stopped")
		} else {
			return fmt.Errorf("server run: %w", err)
		}
	}

	<-idleConnClosed

	return nil
}
