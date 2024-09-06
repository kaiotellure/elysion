package handlers

import (
	"net/http"
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
	SetupRoutes()
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	components.Document(
		components.PageProps{
			Request: r, Auth: getCredential(r),
			Title: "Página Não Encontrada",
		},
		components.NotFound(r),
	).Render(r.Context(), w)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	components.Document(
		components.PageProps{
			Request: r, Auth: getCredential(r),
			Title: "Elysion Bistro Restaurante",
		},
		components.PageHome(),
	).Render(r.Context(), w)
}

func SetupRoutes() {
	Router.NotFound(notFoundHandler)
	Router.Get("/", handleHome)
	Router.Route("/conta", routeAccount)
}

func routeAccount(r chi.Router) {
	r.Get("/", handleConta)
	r.Get("/sair", handleGoogleLogout)
	r.Post("/callback", handleGoogleCallback)
}
