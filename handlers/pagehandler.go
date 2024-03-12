package handlers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/ikaio/tailmplx/components"
	"github.com/ikaio/tailmplx/services/google"
)

type PageHandler struct {
	Title string
	Page  func(props components.PageProps) templ.Component
	Put   func(w http.ResponseWriter, r *http.Request)
}

func (h *PageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.Get(w, r)
	case http.MethodPut:
		h.Put(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
	}
}

func (h *PageHandler) Get(w http.ResponseWriter, r *http.Request) {
	props := components.PageProps{
		Request: r,
	}

	token, err := r.Cookie("g_credential")
	if err == nil {
		credential, _ := google.ParseJWTCredential(token.Value)
		props.Auth = credential
	}

	components.Document(props, h.Page(props), h.Title).Render(r.Context(), w)
}
