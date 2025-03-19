package service

import (
	repository2 "library/internal/repository"
	"library/models"
)

// Интерфейс для описания методов бизнес логики книг
type Book interface {
	CreateBook(book models.Book) (id int, err error) //Добавление книг
	RemoveBook(book models.Book) (err error)         //Удаление книг
	EditBook(id int) (err error)                     //Редактирование книги по инвентарному номеру(id)
}

type Books interface {
	GetAllInfoForBook() ([]models.Book, error) //Вывод списка книг
	//GetBookName() ([]string, error)
}

type SearchBook interface {
	SearchBookPerId(id int) (err error)       //Редактирование книги по инвентарному номеру(id)
	SearchBookPerAuthor(id int) (err error)   //Редактирование книги по имени автора(Author)
	SearchBookPerBookName(id int) (err error) //Редактирование книги по названию книги(BookName)
}

// Создаем конструктор
type Service struct {
	Book
	Books
	SearchBook
}

// Создаем сервис,в который заложим репозиторий
func NewService(repo *repository2.Repository) *Service {
	return &Service{
		Book:       nil,
		Books:      NewBooksService(repo.Books),
		SearchBook: nil,
	}
}
