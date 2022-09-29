package domain

import (
	"github.com/lib/pq"
)

type Book struct {
	ID        int            `db:"id" json:"id" binding:"omitempty"`
	Title     string         `db:"title" json:"title" binding:"required,gte=1"`
	Authors   pq.StringArray `db:"authors" json:"authors" binding:"required,gte=1"`
	Year      string         `db:"year" json:"year" binding:"required,datetime=2006-01-02"`
	Price     float32        `db:"price" json:"price" binding:"required"`
	Amount    int            `db:"amount" json:"amount" binding:"omitempty"`
	Publisher string         `db:"publisher" json:"publisher" binding:"omitempty"`
}
