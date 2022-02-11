package routes

import (
	"github.com/a9u/we-mock-api/config"
	"github.com/a9u/we-mock-api/pkg/wlog"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"net/http"
	"time"
)

type APIRouter struct {
	mux  *chi.Mux
	conf *config.Conf
}

func NewAPIRouter(cfg *config.Conf) (*APIRouter, error) {
	wlog.Info("initialising router")
	router := chi.NewRouter()

	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.RequestID,
		middleware.RealIP,
		middleware.Timeout(60*time.Second),
		middleware.Recoverer,
	)

	wlog.Info("adding basic handlers")
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	return &APIRouter{
		conf: cfg,
		mux:  router,
	}, nil
}

func (apiRouter *APIRouter) Mux() *chi.Mux {
	return apiRouter.mux
}
