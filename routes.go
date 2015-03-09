package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"index", "GET", "/", index},
	Route{"apiHandler", "GET", "/api", apiHandler},
	Route{"apiKeyHandler", "GET", "/api/{id}", apiKeyHandler},
	Route{"apiKeyDetailsHandler", "GET", "/api/{id}/details", apiKeyDetailsHandler},
	Route{"webKeyHandler", "GET", "/web/{id}", webKeyHandler},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
	}
	router.NotFoundHandler = http.HandlerFunc(notFound)

	return router
}
