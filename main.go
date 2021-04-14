package main

import (
	"log"
	"net/http"

	"github.com/charleyvibez/kitties/app/bundles/kittiesbundle"
	"github.com/gorilla/mux"
)

func main() {
	// Controllers declaration
	kc := &kittiesbundle.KittiesController{}

	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1/").Subrouter()

	// Routes handling
	s.HandleFunc("/kitties", kc.Index).Methods("GET")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
