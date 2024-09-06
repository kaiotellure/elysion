package handlers

import (
	"context"
	"net/http"

	"github.com/ikaio/tailmplx/components"
	"github.com/ikaio/tailmplx/services/google"
)

func getCredential(r *http.Request) *google.GoogleCredential {
	if c, ok := r.Context().Value("credential").(*google.GoogleCredential); ok {
		return c
	}
	return nil
}

func GoogleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if credential_cookie, err := r.Cookie("g_credential"); err == nil {
			if c, err := google.ParseJWTCredential(credential_cookie.Value); err == nil {
				next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "credential", c)))
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

func handleConta(w http.ResponseWriter, r *http.Request) {
	credentials := getCredential(r)

	components.Document(
		components.PageProps{
			Request: r, Auth: credentials,
			Title: "Entrar com o Google",
		},
		components.PageEntrar(credentials),
	).Render(r.Context(), w)
}

func handleGoogleLogout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{Name: "g_credential", MaxAge: -1, Path: "/"})
	w.Header().Set("hx-redirect", "/") // redirect htmx buttons
	http.Redirect(w, r, r.URL.Query().Get("resume"), http.StatusSeeOther)
}

func handleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("credential")
	c, err := google.ParseJWTCredential(token)

	if err != nil {
		components.Document(
			components.PageProps{
				Auth: c, Request: r,
				Title: "Google Login Failed",
			},
			components.GoogleError(err.Error()),
		).Render(r.Context(), w)
	}

	http.SetCookie(w, &http.Cookie{Name: "g_credential", Value: token, Path: "/"})
	http.Redirect(w, r, r.URL.Query().Get("resume"), http.StatusSeeOther)
}
