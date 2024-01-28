package router

import (
	"net/http"

	"github.com/ikaio/tailmplx/database"
)

func SessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session_cookie, _ := r.Cookie("session")

		if session_cookie == nil {
			session_cookie = &http.Cookie{
				Name: "session", HttpOnly: true,
				Value: database.SF.Generate().String(),
			}
			http.SetCookie(w, session_cookie)
		}

		r.Header.Set("session", session_cookie.Value)
		next.ServeHTTP(w, r)
	})
}
