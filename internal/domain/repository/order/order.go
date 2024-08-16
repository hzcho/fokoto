package order

import "fokoto/internal/domain/model/order"

type OrderRepository interface {
	Save(order order.Order) (uint64, error)
	Get() ([]order.Order, error)
}
