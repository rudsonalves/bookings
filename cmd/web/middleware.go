package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

/*
NoSurf start the nosurf package to helps Go to prevent Cross-Site
Request Gorgery attacks. See more information in
https://github.com/justinas/nosurf
*/
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

/*
SessionLoad run a LoadAndSave session method to provides middleware
which automatically loads and saves session data for the current
request, and communicates the session token to and from the client
in a cookie.
*/
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
