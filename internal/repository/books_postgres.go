package repository

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
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

func (b BooksRepository) AddBookToCart(ID int, userID int) error {
	var bookID int
	err := b.db.Get(&bookID, `SELECT book_id FROM basket where book_id=$1`, ID)
	if err.Error() == "sql: no rows in result set" {
		query := `INSERT INTO basket(user_id, book_id) VALUES ($1,$2)`
		_, err = b.db.Exec(query, userID, ID)
	} else if err != nil {
		return err
	}

	return nil
}

func (b BooksRepository) GetAllBooks(filter *domain.BookFilter) ([]domain.Book, error) {
	selectStm := sq.Select("*").From("products")

	filters := make([]sq.Sqlizer, 0)
	if filter.Author != "" {

		filters = append(filters, sq.Expr("'"+filter.Author+"'=ANY(authors)"))
	}

	if filter.Publisher != "" {
		filters = append(filters, sq.Eq{"publisher": filter.Publisher})
	}

	if filter.Category != "" {
		filters = append(filters, sq.Eq{"category": filter.Category})
	}

	if filter.Name != "" {
		filters = append(filters, sq.ILike{"title": fmt.Sprintf("%%%s%%", filter.Name)})
	}

	if len(filters) > 0 {
		selectStm = selectStm.
			Where(sq.And(filters)).PlaceholderFormat(sq.Dollar)
	}

	if filter.Sort == "1" {
		selectStm = selectStm.OrderBy("price")
	} else if filter.Sort == "2" {
		selectStm = selectStm.OrderBy("price DESC")
	}
	sql, args, err := selectStm.ToSql()

	fmt.Println(sq.DebugSqlizer(selectStm))
	fmt.Println(sql)
	fmt.Println(args)
	if err != nil {
		return nil, err
	}

	books := []domain.Book{}
	err = b.db.Select(&books, sql, args...)
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

func (b BooksRepository) DeleteBookById(id int, userID int) error {
	query := `DELETE from basket where user_id = $2 AND book_id = $1;`
	_, err := b.db.Exec(query, id, userID)
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

func (b BooksRepository) GetBookFromCartByUserId(id int) ([]domain.Book, error) {
	selectStm := sq.Select("ID", "title", "authors", "year", "publisher", "amount", "price", "description", "category").
		From("products JOIN basket b on products.ID = b.book_ID").Where(sq.Eq{"user_id": id}).PlaceholderFormat(sq.Dollar)

	sql, args, err := selectStm.ToSql()

	if err != nil {
		return nil, err
	}

	books := []domain.Book{}
	err = b.db.Select(&books, sql, args...)
	if err != nil {
		return nil, err
	}

	return books, err
}

func NewBooksRepository(db *sqlx.DB) *BooksRepository {
	return &BooksRepository{db: db}
}
