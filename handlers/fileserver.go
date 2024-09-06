package handlers

import (
	"errors"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/ikaio/tailmplx/helpers"
)

func FileServer(router *chi.Mux, path, root string) {
	rootfs := http.Dir(root)

	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		router.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
		path += "/"
	}
	path += "*"

	router.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")

		_, err := os.Stat(filepath.Join(root, r.URL.Path))
		if errors.Is(err, fs.ErrNotExist) {
			router.NotFoundHandler().ServeHTTP(w, r)
			return
		}

		switch helpers.Env(helpers.MODE) {
		case "production":
			// 1 week cache, catch weekly updates
			w.Header().Set("Cache-Control", "public, max-age=604800")
		case "development":
			// allow css realtime update while developing
			w.Header().Set("Cache-Control", "no-store")
		}

		fs := http.StripPrefix(pathPrefix, http.FileServer(rootfs))
		fs.ServeHTTP(w, r)
	})
}
