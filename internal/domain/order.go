package domain

import "github.com/lib/pq"

type Order struct {
	Id         int           `db:"id" json:"id" binding:"omitempty"`
	UserId     int           `db:"user_id" json:"userId" binding:"omitempty"`
	TotalPrice float64       `db:"totalPrice" json:"totalPrice" binding:"omitempty"`
	Items      pq.Int64Array `db:"items" json:"items" binding:"required"`
	Status     string        `db:"status" json:"status" binding:"required"`
}
