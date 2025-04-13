package main

import (
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
)

func Test_Routes_Exists(t *testing.T) {
	testRoutes := testApp.routes()
	chiRoutes := testRoutes.(chi.Router)

	// this routes must exists
	routes := []string{
		"/users/login",
		"/users/logout",
		"/validate-token",
		"/admin/users/save",
		"/admin/users/delete",
		"/admin/users",
		"/admin/log-user-out/{id}",
		"/admin/users/get/{id}",
		"/books",
		"/books/{slug}",
		"/admin/authors/all",
		"/admin/books/save",
		"/admin/books/delete",
		"/admin/books/{id}",
	}

	for _, r := range routes {
		routesExists(t, chiRoutes, r)
	}

}

func routesExists(t *testing.T, routes chi.Router, route string) {
	// assume the route doesn not exist
	found := false

	// walk through all the registered routes
	_ = chi.Walk(routes, func(method, foundRoute string, handler http.Handler,
		middlewares ...func(http.Handler) http.Handler) error {

		// if we find the route we're looking for , set found to true
		if route == foundRoute {
			found = true
		}
		return nil
	})

	// fire an error if we did not find the route
	if !found {
		t.Errorf("did not find %s in registered routes", route)
	}
}
