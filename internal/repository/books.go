package repository

import (
	"database/sql"
	"library/error2"
	"library/models"
)

type BooksPostgres struct {
	db *sql.DB
}

func NewBooksPostgres(db *sql.DB) *BooksPostgres {
	return &BooksPostgres{db: db}
}

// Method for get all book list
func (b *BooksPostgres) GetAllInfoForBook() ([]models.Book, error) {
	//Create query
	query := "SELECT * FROM bookinfo"

	//Create query for get rows from table
	rows, err := b.db.Query(query)
	if err != nil {
		return nil, error2.ErrorWrap("Cant get rows from table", err)
	}

	defer rows.Close()

	//Create empty list for save books
	var books []models.Book
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.Id, &book.BookName, &book.Author, &book.Genre, &book.YearRealease); err != nil {
			return nil, error2.ErrorWrap("Cant scan row", err)
		}
		books = append(books, book)
	}
	return books, nil
}
