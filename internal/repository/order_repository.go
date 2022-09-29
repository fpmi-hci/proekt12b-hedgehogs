package repository

import (
	"github.com/fpmi-hci/proekt12b-hedgehogs/internal/domain"
	"github.com/jmoiron/sqlx"
)

type OrderRepository struct {
	db *sqlx.DB
}

func (o OrderRepository) CreateOrder(order *domain.Order, userId int) error {
	query := `INSERT INTO orders (user_id, totalprice, items, status) VALUES ($1,$2,$3,$4)`
	_, err := o.db.Exec(query, userId, order.TotalPrice, order.Items, order.Status)
	return err
}

func (o OrderRepository) UpdateOrderById(order *domain.Order, id int) error {
	query := `UPDATE orders set totalprice=$2, items=$3, status=$4 WHERE id=$1`
	_, err := o.db.Exec(query, id, order.TotalPrice, order.Items, order.Status)
	return err
}

func NewOrderRepository(db *sqlx.DB) *OrderRepository {
	return &OrderRepository{db: db}
}
