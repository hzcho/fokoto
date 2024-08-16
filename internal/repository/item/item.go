package item

import (
	"fmt"
	"fokoto/internal/domain/model/item"
	"fokoto/internal/repository/psqldb"
	"github.com/jmoiron/sqlx"
)

type ItemRepo struct {
	db *sqlx.DB
}

func NewItemRepo(db *sqlx.DB) *ItemRepo {
	return &ItemRepo{
		db: db,
	}
}

func (r *ItemRepo) SaveAll(orderId uint64, items []item.Item) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	defer func() {
		if rErr := recover(); rErr != nil || err != nil {
			tx.Rollback()
		}
	}()

	for _, item := range items {
		query := fmt.Sprintf("INSERT INTO %s (order_id, amount, discounted_amount) VALUES ($1, $2, $3)", psqldb.ItemsTable)
		_, err = tx.Exec(query, orderId, item.Amount, item.DiscountedAmount)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (r *ItemRepo) Get(orderId uint64) ([]item.Item, error) {
	var items []item.Item

	query := fmt.Sprintf("SELECT * FROM %s WHERE order_id = $1", psqldb.ItemsTable)
	err := r.db.Select(&items, query, orderId)
	if err != nil {
		return nil, err
	}

	return items, nil
}
