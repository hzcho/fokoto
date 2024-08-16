package order

import (
	"fmt"
	"fokoto/internal/domain/model/order"
	"fokoto/internal/repository/psqldb"
	"github.com/jmoiron/sqlx"
)

type OrderRepo struct {
	db *sqlx.DB
}

func NewOrderRepo(db *sqlx.DB) *OrderRepo {
	return &OrderRepo{
		db: db,
	}
}

func (r *OrderRepo) Save(order order.Order) (uint64, error) {
	var id uint64

	query := fmt.Sprintf(
		"insert into %s (status, user_id, payment_type) values($1, $2, $3) returning id",
		psqldb.OrdersTable,
	)
	if err := r.db.QueryRow(query, order.Status, order.UserID, order.PaymentType).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *OrderRepo) Get() ([]order.Order, error) {
	var orders []order.Order

	query := fmt.Sprintf("select * from %s", psqldb.OrdersTable)
	if err := r.db.Select(&orders, query); err != nil {
		return nil, err
	}

	return orders, nil
}
