package service

import (
	"library/internal/repository"
	"library/models"
)

type BookService struct {
	repo *repository.Book
}

func NewBookService(repo *repository.Book) *BookService {
	return &BookService{repo: repo}
}

// Метод добавление книги
func (b *BookService) CreateBook(book models.Book) (id int, err error) {
	return 0, nil
}
