package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/ikaio/tailmplx/components"
)

func main() {
	http.Handle("/", templ.Handler(components.Index()))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":3000", nil)
}
