package services

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ikaio/tailmplx/ui"
)

type GoogleAuthObject struct {
	jwt.Claims
	Sub     string `json:"sub"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

const ERR_CREDENTIAL_PARSE = "The Google login went well, but we had an issue parsing the Google response."

func HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	credential_token := r.FormValue("credential")

	token, _, err := jwt.NewParser().ParseUnverified(credential_token, &GoogleAuthObject{})
	if err != nil {
		ui.Document(GoogleError(ERR_CREDENTIAL_PARSE), "Google Login Failed").Render(r.Context(), w)
		return
	}

	if auth, ok := token.Claims.(*GoogleAuthObject); ok {
		http.SetCookie(w, &http.Cookie{Name: "g_credential", Value: credential_token})
		ui.Document(GoogleLoginSuccess(auth), "Google Login Successful").Render(r.Context(), w)
	} else {
		ui.Document(GoogleError(ERR_CREDENTIAL_PARSE), "Google Login Failed").Render(r.Context(), w)
		return
	}

}
