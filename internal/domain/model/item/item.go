package item

type Item struct {
	ID               uint64 `db:"item_id"`
	OrderID          uint64 `db:"order_id"`
	Amount           uint64 `db:"amount"`
	DiscountedAmount uint64 `db:"discounted_amount"`
}
