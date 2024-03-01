package router

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/ikaio/tailmplx/ui"
)

type PageHandler struct {
	Title string
	Page  func(r *http.Request) templ.Component
	Put   func(w http.ResponseWriter, r *http.Request)
}

func (h *PageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.Get(w, r)
	case http.MethodPut:
		h.Put(w, r)
	default:
		http.Error(w, "Method not implemented.", http.StatusNotImplemented)
	}
}

func (h *PageHandler) Get(w http.ResponseWriter, r *http.Request) {
	ui.Document(h.Page(r), h.Title).Render(r.Context(), w)
}
