package order

import (
	"database/sql"

	"github.com/peterest/go-basic-ecom/types"
)

type OrderRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (s *OrderRepository) CreateOrder(order types.Order) (int, error) {
	res, err := s.db.Exec(
		"INSERT INTO orders (user_id, total, status, address) VALUES (?, ?, ?, ?)",
		order.UserID,
		order.Total,
		order.Status,
		order.Address,
	)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (s *OrderRepository) CreateOrderItem(orderItem types.OrderItem) error {
	_, err := s.db.Exec(
		"INSERT INTO order_items (order_id, product_id, quantity, price) VALUES (?, ?, ?, ?)",
		orderItem.OrderID,
		orderItem.ProductID,
		orderItem.Quantity,
		orderItem.Price,
	)
	return err
}
