package server

import (
	"errors"
	"fmt"
	"github.com/a9u/we-mock-api/config"
	"github.com/a9u/we-mock-api/pkg/wlog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	conf   *config.Conf
	server *http.Server
}

func New(cfg *config.Conf) *Server {
	return &Server{
		conf: cfg,
		server: &http.Server{
			Addr:         cfg.Port,
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
		},
	}
}

func (s *Server) ListenAndServe() error {
	// catch
	defer func() {
		if r := recover(); r != nil {
			// something went wrong
			wlog.Error("server cannot be started", errors.New(fmt.Sprintf("%#v", r)))
		}
	}()

	go func() {
		defer func() {
			if r := recover(); r != nil {
				// something went wrong
				wlog.Error("listen and serve failed", errors.New(fmt.Sprintf("%#v", r)))
			}
		}()

		wlog.Info("starting server")
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			wlog.Error("failed: listenAndServe", err)
			panic(err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	wlog.Print("Shutting down server")

	return nil
}
