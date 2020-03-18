package main

import (
	// "fmt"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"database/sql"
	"github.com/lib/pq"
	"github.com/subosito/gotenv"
	"github.com/gorilla/mux"
)

type Book struct {
	ID     int    `json:id`
	Title  string `json:title`
	Author string `json:author`
	Year   string `json:year`
}

var books []Book

func main() {
	port := ":8080"
	router := mux.NewRouter()
	books = append(books, Book{ID: 1, Title: "Night Has a Thousand Eyes", Author: "Peder", Year: "6/6/2019"},
		Book{ID: 2, Title: "Open Season", Author: "Hobart", Year: "2/2/2020"},
		Book{ID: 3, Title: "Hercules in New York", Author: "Malinda", Year: "9/4/2019"},
		Book{ID: 4, Title: "Patrick", Author: "Arni", Year: "2/29/2020"},
		Book{ID: 5, Title: "Twilight Saga: Breaking Dawn - Part 1, The", Author: "Colleen", Year: "12/8/2019"},
		Book{ID: 6, Title: "Above the Street, Below the Water (Over gaden under vandet)", Author: "Christine", Year: "5/16/2019"},
		Book{ID: 7, Title: "Dream Boy", Author: "Larisa", Year: "1/26/2020"},
		Book{ID: 8, Title: "Winnie the Pooh and a Day for Eeyore", Author: "Aksel", Year: "12/14/2019"},
		Book{ID: 9, Title: "Vai~E~Vem", Author: "Erskine", Year: "9/7/2019"},
		Book{ID: 10, Title: "Santa Fe", Author: "Cullie", Year: "12/21/2019"},
		Book{ID: 11, Title: "Princess Protection Program", Author: "Garth", Year: "5/18/2019"},
		Book{ID: 12, Title: "Planet B-Boy", Author: "Griffie", Year: "11/22/2019"},
		Book{ID: 13, Title: "RaidDogs (idapped) (Cani arrabbiati)", Author: "Gloriana", Year: "4/30/2019"},
		Book{ID: 14, Title: "Adventures of Arsène Lupin, The (Aventures d'Arsène Lupin, Les)", Author: "Kermit", Year: "7/23/2019"},
		Book{ID: 15, Title: "Zotz!", Author: "Lou", Year: "2/18/2020"},
		Book{ID: 16, Title: "Amelia", Author: "Lonnard", Year: "1/10/2020"},
		Book{ID: 17, Title: "Dirty Money (Un flic)", Author: "Turner", Year: "9/16/2019"},
		Book{ID: 18, Title: "Land Before Time, The", Author: "Gus", Year: "4/29/2019"},
		Book{ID: 19, Title: "Pirates of Penzance, The", Author: "Gerta", Year: "7/4/2019"},
		Book{ID: 20, Title: "Strange Days", Author: "Flint", Year: "7/2/2019"})
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	done := make(chan bool)
	go http.ListenAndServe(port, router)
	log.Printf("Server started at port %v", port)
	<-done
	log.Fatal(http.ListenAndServe(port, router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}
func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	i, _ := strconv.Atoi(params["id"])

	for _, book := range books {
		if book.ID == i {
			json.NewEncoder(w).Encode(&book)
		}
	}
}
func addBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	books = append(books, book)
	json.NewEncoder(w).Encode(books)
}
func updateBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	for i, item := range books {
		if book.ID == item.ID {
			books[i] = book
		}
	}
	json.NewEncoder(w).Encode(&book)
}
func removeBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for i, item := range books {
		if item.ID == id {
			books = append(books[:i], books[i+1:]...)
		}
	}
	json.NewEncoder(w).Encode(books)
}
