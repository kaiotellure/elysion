package handlers

import (
	"net/http"
	"runtime"

	"github.com/go-chi/chi/v5"
	"github.com/ikaio/tailmplx/components"
)

func routeAdmin(r chi.Router) {
	r.Get("/", handleAdmin)
}

func handleAdmin(w http.ResponseWriter, r *http.Request) {
	c := getCredential(r)
	if c == nil || c.Email != "ikaiodev@gmail.com" {
		components.Document(
			components.PageProps{Request: r, Auth: c},
			components.Warn("You do not have access to this. Please, check if you are logged in and you have the right permission."),
			"No Permission",
		).Render(r.Context(), w)
		return
	}

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	components.Document(
		components.PageProps{Request: r, Auth: c},
		components.AdminDashboard(m),
		"Admin Dashboard",
	).Render(r.Context(), w)
}
