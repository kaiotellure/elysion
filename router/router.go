package router

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ikaio/tailmplx/pages"
	"github.com/ikaio/tailmplx/router/handlers"
)

var Router *chi.Mux = chi.NewRouter()

func Setup(public_folder_path string) {

	// A good base middleware stack
	Router.Use(middleware.RequestID)
	Router.Use(middleware.RealIP)
	Router.Use(middleware.Logger)
	Router.Use(middleware.Recoverer)
	Router.Use(SessionMiddleware)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	Router.Use(middleware.Timeout(60 * time.Second))

	FileServer(Router, "/", public_folder_path)
}

func SetupRoutes() {
	// Simple Pages with single method
	Router.Handle("/", &handlers.PageHandler{Title: "Home", Main: pages.Home})

	// Complex Handlers with services dependencies and multiple methods
	Router.Handle("/apps/create", &handlers.PageHandler{Title: "Create App", Main: pages.AppsCreate})

	Router.NotFound((&handlers.PageHandler{Title: "Not Found", Main: pages.NotFound}).ServeHTTP)
}
