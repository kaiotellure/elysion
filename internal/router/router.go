package router

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ikaio/tailmplx/internal/production"
	"github.com/ikaio/tailmplx/pages"
)

var Router *chi.Mux = chi.NewRouter()

func Setup(public_folder_path string) {

	Router.Use(middleware.RequestID)
	Router.Use(middleware.RealIP)

	Router.Use(middleware.Logger)
	Router.Use(middleware.Recoverer)

	Router.Use(SessionMiddleware)
	Router.Use(middleware.Timeout(60 * time.Second))

	FileServer(Router, "/", public_folder_path)
}

func SetupRoutes() {
	Router.Handle("/", &PageHandler{Title: "Home", Page: pages.Home})
	Router.Handle("/production/new", &PageHandler{Title: "New Production", Page: pages.ProductionNew, Put: production.HandlePut})
	Router.Handle("/production/{id}/edit", &PageHandler{Title: "New Production", Page: pages.ProductionSlugEdit, Put: production.HandlePut})
	Router.NotFound((&PageHandler{Title: "Not Found", Page: pages.NotFound}).ServeHTTP)
}
