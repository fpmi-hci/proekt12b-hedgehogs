package repository

import (
	"github.com/fpmi-hci/proekt12b-hedgehogs/internal/domain"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user *domain.User) (domain.User, error)
	GetUser(username string) (*domain.User, error)
}

type Books interface {
	CreateBook(book *domain.Book) error
	GetAllBooks() ([]domain.Book, error)
	UpdateBookById(book *domain.Book, id int) error
	GetBookById(id int) (*domain.Book, error)
	DeleteBookById(id int) error
	GetBookByAuthor(author string) (*domain.Book, error)
	GetBookByPublisher(publisher string) (*domain.Book, error)
}

type Order interface {
	CreateOrder(order *domain.Order, userId int) error
	UpdateOrderById(order *domain.Order, id int) error
}

type Repository struct {
	Authorization
	Books
	Order
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Books:         NewBooksRepository(db),
		Order:         NewOrderRepository(db),
	}
}
