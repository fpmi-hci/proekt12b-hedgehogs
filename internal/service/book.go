package service

import (
	"github.com/fpmi-hci/proekt12b-hedgehogs/internal/domain"
	"github.com/fpmi-hci/proekt12b-hedgehogs/internal/repository"
)

type BooksService struct {
	repo repository.Books
}

func (b BooksService) GetBookByAuthor(author string) (*domain.Book, error) {
	return b.repo.GetBookByAuthor(author)
}

func (b BooksService) GetBookByPublisher(publisher string) (*domain.Book, error) {
	return b.repo.GetBookByPublisher(publisher)
}

func (b BooksService) CreateBook(book *domain.Book, userId int) error {
	return b.repo.CreateBook(book)
}

func (b BooksService) GetAllBooks() ([]domain.Book, error) {
	return b.repo.GetAllBooks()
}

func (b BooksService) GetBookById(id int) (*domain.Book, error) {
	return b.repo.GetBookById(id)
}

func (b BooksService) UpdateBookById(book *domain.Book, id int) error {
	return b.repo.UpdateBookById(book, id)
}

func (b BooksService) DeleteBookById(id int) error {
	return b.repo.DeleteBookById(id)
}

func NewBooksService(repo repository.Books) *BooksService {
	return &BooksService{repo: repo}
}
