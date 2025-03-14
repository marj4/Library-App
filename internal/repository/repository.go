package repository

import (
	"database/sql"
	"library/models"
)

// / Интерфейс для описания методов бизнес логики книг
type Book interface {
	CreateBook(book models.Book) (id int, err error) //Добавление книг
	RemoveBook(book models.Book) (err error)         //Удаление книг
	EditBook(id int) (err error)                     //Редактирование книги по инвентарному номеру(id)
}

type Books interface {
	GetAllInfoForBook() ([]models.Book, error) //Вывод списка книг
}

type SearchBook interface {
	SearchBookPerId(id int) (err error)       //Редактирование книги по инвентарному номеру(id)
	SearchBookPerAuthor(id int) (err error)   //Редактирование книги по имени автора(Author)
	SearchBookPerBookName(id int) (err error) //Редактирование книги по названию книги(BookName)
}

type Repository struct {
	Book
	Books
	SearchBook
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Book:       nil,
		Books:      nil,
		SearchBook: nil,
	}
}
