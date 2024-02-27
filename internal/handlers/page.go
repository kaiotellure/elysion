package handlers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/ikaio/tailmplx/ui"
)

type PageHandler struct {
	Title string
	Main  func(r *http.Request) templ.Component
}

func (h *PageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.Get(w, r)
	default:
		http.Error(w, "Method not implemented.", http.StatusNotImplemented)
	}
}

func (h *PageHandler) Get(w http.ResponseWriter, r *http.Request) {
	ui.Page(h.Main(r), h.Title).Render(r.Context(), w)
}
