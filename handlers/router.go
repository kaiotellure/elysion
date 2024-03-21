package handlers

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ikaio/tailmplx/components"
)

var Router *chi.Mux = chi.NewRouter()

func Setup(public_folder_path string) {

	Router.Use(middleware.RequestID)
	Router.Use(middleware.RealIP)

	Router.Use(middleware.Logger)
	Router.Use(middleware.Recoverer)

	Router.Use(middleware.Timeout(60 * time.Second))
	Router.Use(GoogleMiddleware)

	FileServer(Router, "/", public_folder_path)
}

func SetupRoutes() {
	Router.Handle("/", &PageHandler{Title: "Home", Page: components.PageHome})
	Router.NotFound((&PageHandler{Title: "Not Found", Page: components.NotFound}).ServeHTTP)

	Router.Route("/admin", routeAdmin)
	Router.Route("/production", routeProduction)
	Router.Route("/account/google", routeAccountGoogle)
}

func routeAccountGoogle(r chi.Router) {
	r.Get("/", handleAccountGoogle)
	r.Post("/logout", handleGoogleLogout)
	r.Post("/callback", handleGoogleCallback)
}
