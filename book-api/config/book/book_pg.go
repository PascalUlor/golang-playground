package bookDbConfig

import (
	"database/sql"
	"log"

	"../../models"
)

type BookDbConfig struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (c BookDbConfig) GetBooks(db *sql.DB, book models.Book, books []models.Book) []models.Book {
	rows, err := db.Query("SELECT * FROM books")
	logFatal(err)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		logFatal(err)

		books = append(books, book)
	}
	return books
}


func (c BookDbConfig) GetBook(db *sql.DB, book models.Book, id int) models.Book {
	err := db.QueryRow("select * from books where id = $1", id).Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	logFatal(err)

	return book
}

func (c BookDbConfig) AddBook(db *sql.DB, book models.Book) int {
	err := db.QueryRow("insert into books (title, author, year) values($1, $2, $3) returning id;", book.Title, book.Author, book.Year).Scan(&book.ID)
		logFatal(err)
	return book.ID
}

func (c BookDbConfig) UpdateBook(db *sql.DB, book models.Book) int64 {
	result, err := db.Exec("update books set title=$1, author=$2, year=$3 where id=$4 returning id;", &book.Title, &book.Author, &book.Year, &book.ID)

		rowsUpdated, err := result.RowsAffected()
		logFatal(err)
	return rowsUpdated
}

func (c BookDbConfig) RemoveBook(db *sql.DB, id int) int64 {
	result, err := db.Exec("delete from books where id=$1", id)
		logFatal(err)

		rowsDeleted, err := result.RowsAffected()
		logFatal(err)

		return rowsDeleted
}
