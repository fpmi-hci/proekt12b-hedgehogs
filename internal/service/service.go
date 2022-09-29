package service

import (
	"github.com/fpmi-hci/proekt12b-hedgehogs/internal/domain"
	"github.com/fpmi-hci/proekt12b-hedgehogs/internal/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	CreateUser(user *domain.User) (domain.User, error)
	GenerateToken(username, password string) (string, error)
	GetUserByUsername(username string) (*domain.User, error)
	ParseToken(token string) (int, error)
}

type Books interface {
	CreateBook(book *domain.Book, userId int) error
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
type Service struct {
	Authorization
	Books
	Order
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(r.Authorization),
		Books:         NewBooksService(r.Books),
		Order:         NewOrderService(r.Order),
	}
}
