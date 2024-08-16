package item

import "fokoto/internal/domain/model/item"

type ItemRepository interface {
	SaveAll(orderId uint64, items []item.Item) error
	Get(orderId uint64) ([]item.Item, error)
}
