package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/ajay-code/go-bookstore/pkg/models"
	"github.com/ajay-code/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var Book models.Book

func CreateBook(w http.ResponseWriter, r *http.Request) {
	newBook := &models.Book{}

	utils.ParseBody(r, newBook)
	newBook.CreateBook()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newBook)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	
	w.Header().Set("Content-Type", "application/json")
	
	json.NewEncoder(w).Encode(books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID, err := strconv.ParseInt(vars["bookId"],0,0)
	if err != nil {
		log.Fatalln(err)
	}
	book, _ := models.GetBookById(bookID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID, err := strconv.ParseInt(vars["bookId"], 0, 0)
	if err != nil {
		log.Fatalln(err)
	}

	book, _ := models.GetBookById(bookID)
	utils.ParseBody(r, book)
	book.Save()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars:= mux.Vars(r)
	bookID, err := strconv.ParseInt(vars["bookId"],0,0)
	if err != nil {
		log.Fatalln(err)
	}
	models.DeleteBookById(bookID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}