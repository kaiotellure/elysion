package handlers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/ikaio/tailmplx/pages"
)

func NewPage() *PageHandler {
	return &PageHandler{}
}

type PageHandler struct{
	Title string
	Main func(r *http.Request) templ.Component
}

func (h *PageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.Post(w, r)
	case http.MethodGet:
		h.Get(w, r)
	}
}

func (h *PageHandler) Post(w http.ResponseWriter, r *http.Request) {

}

func (h *PageHandler) Get(w http.ResponseWriter, r *http.Request) {
	pages.Page(h.Main(r), h.Title, pages.DEFAULT_PROPS).Render(r.Context(), w)
}
