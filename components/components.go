package components

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type LazyHandler struct {
	Path     string
	Callback func(w http.ResponseWriter, r *http.Request)
}

// components register themselves into during init()
// then, when the router is ready it attaches the handlers to himself
var LazyHandlerPool []*LazyHandler

func RegisterHandler(handler *LazyHandler) {
	LazyHandlerPool = append(LazyHandlerPool, handler)
}

func AttachHandlersFromPoolTo(router *chi.Mux) {
	for _, handler := range LazyHandlerPool {
		router.HandleFunc(handler.Path, handler.Callback)
	}
}
