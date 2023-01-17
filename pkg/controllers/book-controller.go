package controllers

import (
	"bookManageGO/pkg/models"
	"bookManageGO/pkg/utils"
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

// here we transform all data to json and sent to client

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request){
	res, _ := json.Marshal(models.GetAllBooks())
	w.Header().Set("Content-type", "application/json")
	w.Write(res)
}
func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars  := mux.Vars(r)
	book := vars["bookId"]
	id, _ := strconv.ParseInt(book, 0, 0)
	toJson, _ := models.GetBookById(id)
	res, _ := json.Marshal(toJson)
	w.Header().Set("Content-type", "application/json")
	w.Write(res)
}
func CreateBook(w http.ResponseWriter, r *http.Request) {
	vars := &models.Book{}
	utils.ParseBody(r, vars)
	b := vars.CreateBook()
	res, _ := json.Marshal(b)
	w.Write(res)
}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	book := vars["bookId"]
	id, _ := strconv.ParseInt(book, 0, 0)
	toJson := models.DeleteBook(id)
	res, _ := json.Marshal(toJson)
	w.Write(res)
}
func DeleteAll(w http.ResponseWriter, r *http.Request) {
	res := models.Delete()
	json.Marshal(res)
}
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	book := vars["bookId"]
	toJson, _ := strconv.ParseInt(book, 0, 0)
	res, db := models.GetBookById(toJson)
	if updateBook.Name != "" {
		res.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		res.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		res.Publication = updateBook.Publication
	}
	db.Save(&res)
	marsh, _ := json.Marshal(res)
	w.Write(marsh)
}
