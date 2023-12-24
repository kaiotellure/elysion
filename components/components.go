package components

import (
	"context"
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
)

func page(r chi.Router, name string, compexe func(cookies string) templ.Component, title string, props PageProps) {
	r.Get(name, func(w http.ResponseWriter, r *http.Request) {
		cookies := r.Header.Get("cookie")
		Page(compexe(cookies), "Nalvok® / "+title, props).Render(context.TODO(), w)
	})
}

func dynamic(r chi.Router, name string, compexe func(cookies string) templ.Component) {
	r.Get("/dc/"+name, func(w http.ResponseWriter, r *http.Request) {
		cookies := r.Header.Get("cookie")
		compexe(cookies).Render(context.TODO(), w)
	})
}

func render(component templ.Component) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := component.Render(context.TODO(), w)
		if err != nil {
			fmt.Println(fmt.Errorf("Error while rendering: %w", err))
		}
	}
}

func Init(r chi.Router) {

	page(r, "/", Home, "Home", DEFAULT_PROPS)
	page(r, "/about", About, "About", DEFAULT_PROPS)
	page(r, "/account", Account, "Account", DEFAULT_PROPS)

	dynamic(r, "products", Products)

	r.NotFound(templ.Handler(Page(NotFound(), "Nalvok® / 404", DEFAULT_PROPS)).ServeHTTP)
}
