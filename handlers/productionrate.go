package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ikaio/tailmplx/components"
	"github.com/ikaio/tailmplx/services/production"
)

func handleProductionRate(w http.ResponseWriter, r *http.Request) {
	c := getCredential(r)

	// user is not logged in, show google login prompt
	if c == nil {
		components.GoogleOneTapPrompt(r.Header.Get("referer"), true).Render(r.Context(), w)
		return
	}

	p, err := production.GetById(chi.URLParam(r, "id"))

	// production not found, display error
	if err != nil {
		components.GoogleError(err.Error()).Render(r.Context(), w)
		return
	}

	rate := r.URL.Query().Get("type")
	err = production.StoreRating(rate, p.ID, c.Sub)

	// could not save rating of this user to this production on db
	if err != nil {
		components.GoogleError(err.Error()).Render(r.Context(), w)
		return
	}

	love, like := production.CountProductionRating(p.ID)

	components.ProductionRating(
		*p, components.ProductionRatingData{rate, love, like},
	).Render(r.Context(), w)
}
