package router

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ikaio/tailmplx/ui"
	"github.com/ikaio/tailmplx/internal/handlers"
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
	Router.Handle("/", &handlers.PageHandler{Title: "Home", Main: ui.Home})
	Router.Handle("/publish", &handlers.PageHandler{Title: "Home", Main: ui.Publish})
	Router.NotFound((&handlers.PageHandler{Title: "Not Found", Main: ui.NotFound}).ServeHTTP)
}
