package order

import (
	"fokoto/internal/domain/model/item"
)

type Order struct {
	ID               uint64
	Status           Status
	UserID           uint64      `db:"user_id"`
	PaymentType      PaymentType `db:"payment_type"`
	OriginalAmount   uint64
	DiscountedAmount uint64
	Items            []item.Item
}

type Status uint8

const (
	UnknownStatus Status = iota
	CreatedStatus
	ProcessedStatus
	CanceledStatus
)

type PaymentType uint8

const (
	UnknownType PaymentType = iota
	Card
	Wallet
)
