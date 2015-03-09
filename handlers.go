package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// Render is the global renderer for all views.
var Render *render.Render

// Then you can initialize it or add this to a setup method elsewhere.
func init() {
	Render = render.New(render.Options{
		IndentJSON: true,
	})
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Welcome to the api hander page!")
	Render.JSON(w, http.StatusOK, "Welcome to the api hander page!")
}

func apiKeyHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	fmt.Fprintf(w, "Welcome to the api Key hander page! %s", id)
}

func apiKeyDetailsHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	fmt.Fprintf(w, "Welcome to the api Key Details hander page! %s", id)
}

func webKeyHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	fmt.Fprintf(w, "Welcome to the web Key hander page! %s", id)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
