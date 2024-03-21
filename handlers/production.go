package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ikaio/tailmplx/components"
	"github.com/ikaio/tailmplx/services/database"
	"github.com/ikaio/tailmplx/services/production"
)

func routeProduction(r chi.Router) {
	r.Get("/new", handleProductionNew)
	r.Put("/new", production.HandlePut)

	r.Post("/import", handleProductionImport)

	r.Get("/{id}", handleProduction)
	r.Handle("/{id}/edit", &PageHandler{Title: "Edit Production", Page: components.PageProductionSlugEdit, Put: production.HandlePut})
	r.Post("/{id}/rate", handleProductionRate)
	r.Post("/{id}/comment", handleProductionComment)
}

func handleProduction(w http.ResponseWriter, r *http.Request) {
	c := getCredential(r)
	p, err := production.GetById(chi.URLParam(r, "id"))

	if err != nil {
		components.Warn(err.Error()).Render(r.Context(), w)
		return
	}

	var rate string = "none"
	if c != nil {
		// only fetch user rating if logged in, was getting nil pointer deference
		production.RetrieveRating("none", p.ID, c.Sub)
	}

	love, like := production.CountProductionRating(p.ID)
	data := components.ProductionRatingData{rate, love, like}

	components.Document(
		components.PageProps{Request: r, Auth: c},
		components.ProductionLanding(*p, c, data),
		p.Title,
	).Render(r.Context(), w)
}

func handleProductionNew(w http.ResponseWriter, r *http.Request) {
	components.Document(
		components.PageProps{Request: r, Auth: getCredential(r)},
		components.ProductionEditor(production.Production{ID: database.NewUUID()}),
		"New Production",
	).Render(r.Context(), w)
}
