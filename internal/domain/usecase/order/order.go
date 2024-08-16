package order

import "fokoto/internal/domain/model/order"

type OrderUseCase interface {
	Save(order order.Order) (int64, error)
	Get() ([]order.Order, error)
}
