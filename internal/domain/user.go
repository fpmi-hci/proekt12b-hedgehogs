package domain

import _ "github.com/lib/pq"

type User struct {
	ID           int     `db:"id" json:"id" binding:"omitempty"`
	Name         string  `db:"name" json:"name" binding:"omitempty"`
	LastName     string  `db:"last_name" json:"lastName" binding:"omitempty"`
	Login        string  `db:"login" json:"login" binding:"required"`
	PasswordHash string  `db:"password_hash" json:"passwordHash" binding:"required,gte=8"`
	Email        *string `db:"email" json:"email" binding:"omitempty"`
	Phone        *string `db:"phone" json:"phone" binding:"omitempty,e164"`
}
