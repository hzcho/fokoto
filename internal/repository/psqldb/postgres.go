package psqldb

import (
	"fmt"
	"fokoto/internal/config"
	"github.com/jmoiron/sqlx"
)

const (
	OrdersTable = "orders"
	ItemsTable  = "items"
)

func New(cfg config.DB) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode))

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
