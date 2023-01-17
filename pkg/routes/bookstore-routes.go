package routes

import (
	"bookManageGO/pkg/controllers"

	"github.com/gorilla/mux"
)

//here we just handle all the methods


var RegisterBookStore = func(r *mux.Router) {
	r.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	r.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	r.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
	r.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/book/", controllers.DeleteAll).Methods("DELETE")
}