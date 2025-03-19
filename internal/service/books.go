package service

import (
	"library/internal/repository"
	"library/models"
)

type BooksService struct {
	repo repository.Books
}

func (b BooksService) GetAllInfoForBook() ([]models.Book, error) {
	return b.repo.GetAllInfoForBook()
}

func NewBooksService(repo repository.Books) *BooksService {
	return &BooksService{repo: repo}
}
