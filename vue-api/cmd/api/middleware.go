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
