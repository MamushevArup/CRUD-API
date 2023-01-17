package main

import (
	"bookManageGO/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// in main we only start a server work
func main() {
	r := mux.NewRouter()
	routes.RegisterBookStore(r)
	http.Handle("/", r)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("Error", err)
	}
}