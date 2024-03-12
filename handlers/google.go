package handlers

import (
	"net/http"

	"github.com/ikaio/tailmplx/components"
	"github.com/ikaio/tailmplx/services/google"
)

func handleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("credential")
	c, err := google.ParseJWTCredential(token)

	if err != nil {
		components.Document(
			components.PageProps{Auth: c, Request: r},
			components.GoogleError(err.Error()),
			"Google Login Failed",
		).Render(r.Context(), w)
	}

	http.SetCookie(w, &http.Cookie{Name: "g_credential", Value: token, Path: "/"})

	components.Document(
		components.PageProps{Auth: c, Request: r},
		components.GoogleLoginSuccess(c, r.URL.Query().Get("resume")),
		"Google Login Successful",
	).Render(r.Context(), w)
}
