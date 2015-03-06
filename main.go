package main

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the api hander page!")
}

func apiKeyHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	fmt.Fprintf(w, "Welcome to the api Key hander page! %s", id)
}

func apiKeyDetailsHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	fmt.Fprintf(w, "Welcome to the api Key hander page! %s", id)
}

func webKeyHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	fmt.Fprintf(w, "Welcome to the web Key hander page! %s", id)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/api/", apiHandler)
	router.HandleFunc("/api/{id}", apiKeyHandler)
	router.HandleFunc("/api/{id}/details", apiKeyDetailsHandler)
	router.HandleFunc("/web/{id}", webKeyHandler)

	n := negroni.New()
	n.UseHandler(router)
	n.Run(":3010")
}
