package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// NoSurf adds SCRF protection to all POST requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",   // for all site
		Secure:   false, // no https, may change in production
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}
func (app *application) AuthTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := app.models.Token.AuthenticateToken(r)
		if err != nil {
			payload := jsonResponse{
				Error:   true,
				Message: "invalid authentication credentials",
			}

			_ = app.writeJSON(w, http.StatusUnauthorized, payload)
			return
		}

		next.ServeHTTP(w, r)

	})
}
