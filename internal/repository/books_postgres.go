package repository

import (
	"github.com/fpmi-hci/proekt12b-hedgehogs/internal/domain"
	"github.com/jmoiron/sqlx"
)

type BooksRepository struct {
	db *sqlx.DB
}

func (b BooksRepository) GetBookByAuthor(author string) (*domain.Book, error) {
	query := `SELECT * FROM products WHERE $1 = ANY (authors)`
	book := domain.Book{}
	err := b.db.Get(&book, query, author)
	if err != nil {
		return nil, err
	}

	return &book, err
}

func (b BooksRepository) GetBookByPublisher(publisher string) (*domain.Book, error) {
	query := `SELECT * FROM products WHERE publisher = $1`
	book := domain.Book{}
	err := b.db.Get(&book, query, publisher)
	if err != nil {
		return nil, err
	}

	return &book, err
}

func (b BooksRepository) CreateBook(book *domain.Book) error {
	query := `INSERT INTO products(title, authors, year, price, amount, publisher) VALUES ($1,$2,$3,$4,$5,$6)`
	_, err := b.db.Exec(query, book.Title, book.Authors, book.Year, book.Price, book.Amount, book.Publisher)
	return err
}

func (b BooksRepository) GetAllBooks() ([]domain.Book, error) {
	books := []domain.Book{}
	err := b.db.Select(&books, `SELECT * FROM products`)
	if err != nil {
		return nil, err
	}

	return books, err
}

func (b BooksRepository) GetBookById(id int) (*domain.Book, error) {
	book := domain.Book{}
	err := b.db.Get(&book, `SELECT * FROM products where id=$1`, id)
	if err != nil {
		return nil, err
	}

	return &book, err
}

func (b BooksRepository) DeleteBookById(id int) error {
	query := `DELETE from products where id = $1;`
	_, err := b.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (b BooksRepository) UpdateBookById(book *domain.Book, id int) error {
	query := `UPDATE products set title = $2, authors =$3, year = $4, price=$5, amount=$6 where id = $1;`
	_, err := b.db.Exec(query, id, book.Title, book.Authors, book.Year, book.Price, book.Amount)
	if err != nil {
		return err
	}
	return nil
}

func NewBooksRepository(db *sqlx.DB) *BooksRepository {
	return &BooksRepository{db: db}
}
