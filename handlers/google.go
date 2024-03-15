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

func handleAccountGoogle(w http.ResponseWriter, r *http.Request) {
	c := getCredential(r)
	if c == nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	components.Document(
		components.PageProps{Request: r, Auth: c},
		components.GoogleAccountDashboard(c),
		"Account Dashboard",
	).Render(r.Context(), w)
}

func handleGoogleLogout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{Name: "g_credential", MaxAge: -1, Path: "/"})
	w.Header().Set("hx-redirect", "/") // redirect htmx buttons
}

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
	http.Redirect(w, r, r.URL.Query().Get("resume"), http.StatusSeeOther)
}
