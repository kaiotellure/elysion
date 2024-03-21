package handlers

import (
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/ikaio/tailmplx/components"
	"github.com/ikaio/tailmplx/services/database"
	"github.com/ikaio/tailmplx/services/production"
)

var debounce = make(map[string]time.Time)

func handleProductionComment(w http.ResponseWriter, r *http.Request) {

	content := strings.TrimSpace(r.FormValue("content"))
	withoutspaces := strings.ReplaceAll(content, " ", "")

	if len(withoutspaces) < 10 {
		components.Warn("Type something longer.").Render(r.Context(), w)
		return
	}

	if len(withoutspaces) > 500 {
		components.Warn("Type something shorter (max: 500).").Render(r.Context(), w)
		return
	}

	c := getCredential(r)
	if c == nil {
		components.GoogleOneTapPrompt(r.Header.Get("referer"), true).Render(r.Context(), w)
		return
	}

	last_comment_stamp, ok := debounce[c.Sub]
	if ok && time.Since(last_comment_stamp) < 10*time.Second {
		components.Warn("Please, wait a while before sending another comment.").Render(r.Context(), w)
		return
	}

	debounce[c.Sub] = time.Now()

	p, err := production.GetById(chi.URLParam(r, "id"))
	if err != nil {
		components.Warn(err.Error()).Render(r.Context(), w)
		return
	}

	comment := production.Comment{
		ID:            database.NewUUID(),
		AuthorName:    c.Name,
		AuthorPicture: c.Picture,
		Content:       content,
	}

	err = production.StoreComment(comment, p.ID, c.Sub, comment.ID)
	if err != nil {
		components.Warn(err.Error()).Render(r.Context(), w)
		return
	}

	components.Comment(comment, components.GREENISH).Render(r.Context(), w)
}
