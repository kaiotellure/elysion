package handlers

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/ikaio/tailmplx/components"
	"github.com/ikaio/tailmplx/database"
	"github.com/ikaio/tailmplx/services/production"
)

func handleProduction(w http.ResponseWriter, r *http.Request) {
	c := getCredential(r)
	p, err := production.GetById(chi.URLParam(r, "id"))

	if err != nil {
		components.GoogleError(err.Error()).Render(r.Context(), w)
		return
	}

	rate := "none"
	if c != nil {
		rate = production.RetrieveRating("none", p.ID, c.Sub)
	}

	love, like := production.CountProductionRating(p.ID)

	components.Document(
		components.PageProps{Request: r, Auth: c},
		components.ProductionLanding(
			*p, c, components.ProductionRatingData{rate, love, like},
		),
		p.Title,
	).Render(r.Context(), w)
}

func handleProductionEdit(w http.ResponseWriter, r *http.Request) {
	components.Document(
		components.PageProps{Request: r, Auth: getCredential(r)},
		components.ProductionEditor(
			production.Production{ID: database.SF.Generate().String()},
		),
		"New Production",
	).Render(r.Context(), w)
}

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

var CommentsDebounce = make(map[string]time.Time)

func handleProductionComment(w http.ResponseWriter, r *http.Request) {
	c := getCredential(r)
	if c == nil {
		components.GoogleOneTapPrompt(r.Header.Get("referer"), true).Render(r.Context(), w)
		return
	}

	last_comment_stamp, ok := CommentsDebounce[c.Sub]
	if ok && time.Since(last_comment_stamp) < 10*time.Second {
		components.GoogleError("Hell, you type too fast bro, we can't catch you up. Please, wait a while before sending another comment.").Render(r.Context(), w)
		return
	}

	CommentsDebounce[c.Sub] = time.Now()

	p, err := production.GetById(chi.URLParam(r, "id"))

	// production not found, display error
	if err != nil {
		components.GoogleError(err.Error()).Render(r.Context(), w)
		return
	}

	comment := production.Comment{
		ID:            database.SF.Generate().String(),
		AuthorName:    c.Name,
		AuthorPicture: c.Picture,
		Content:       r.FormValue("content"),
	}

	err = production.StoreComment(comment, p.ID, c.Sub, comment.ID)
	if err != nil {
		components.GoogleError(err.Error()).Render(r.Context(), w)
		return
	}

	components.Comment(comment).Render(r.Context(), w)
}
