package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	bookDbConfig "../config/book"

	"../models"
	"github.com/gorilla/mux"
)

type Controller struct{}

var books []models.Book

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		books = []models.Book{}

		bookStore := bookDbConfig.BookDbConfig{}

		books = bookStore.GetBooks(db, book, books)

		json.NewEncoder(w).Encode(books)
	}
}
func (c Controller) GetBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		params := mux.Vars(r)

		id, err := strconv.Atoi(params["id"])

		logFatal(err)

		bookStore := bookDbConfig.BookDbConfig{}

		book = bookStore.GetBook(db, book, id)

		json.NewEncoder(w).Encode(book)
	}
}

func (c Controller) AddBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var bookID int
		json.NewDecoder(r.Body).Decode(&book)
		
		bookStore := bookDbConfig.BookDbConfig{}

		bookID = bookStore.AddBook(db, book)

		json.NewEncoder(w).Encode(bookID)
	}
}

func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		json.NewDecoder(r.Body).Decode(&book)

		bookStore := bookDbConfig.BookDbConfig{}

		rowsUpdated := bookStore.AddBook(db, book)

		json.NewEncoder(w).Encode(rowsUpdated)

	}
}

func (c Controller) RemoveBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		// result, err := db.Exec("delete from books where id=$1", params["id"])
		// logFatal(err)

		// rowsDeleted, err := result.RowsAffected()
		// logFatal(err)
		id, err := strconv.Atoi(params["id"])

		logFatal(err)

		bookStore := bookDbConfig.BookDbConfig{}

		rowsDeleted := bookStore.RemoveBook(db, id)

		json.NewEncoder(w).Encode(rowsDeleted)
	}

}
