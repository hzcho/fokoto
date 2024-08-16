package handler

import (
	"fokoto/internal/domain/model/item"
	"fokoto/internal/domain/model/order"
	"github.com/labstack/echo/v4"
	"net/http"
)

type SaveOrderReq struct {
	UserId      uint64 `json:"user_id" validate:"required"`
	PaymentType string `json:"payment_type" validate:"required"`
	Items       []Item `json:"items" validate:"required,dive"`
}

type Item struct {
	ID       uint64 `json:"id" validate:"required"`
	Amount   uint64 `json:"amount" validate:"required,gte=0"`
	Discount uint64 `json:"discount" validate:"gte=0"`
}

type PaymentType uint8

const (
	UndefinedType PaymentType = iota
	Card
	Wallet
)

var paymentTypes = map[string]PaymentType{
	"card":   Card,
	"wallet": Wallet,
}

func (in SaveOrderReq) OrderFromDTO() order.Order {
	items := []item.Item{}
	for _, it := range in.Items {
		items = append(items, item.Item{
			ID:               it.ID,
			Amount:           it.Amount,
			DiscountedAmount: it.Discount,
		})
	}
	return order.Order{
		Status:      order.CreatedStatus,
		UserID:      in.UserId,
		PaymentType: order.PaymentType(paymentTypes[in.PaymentType]),
		Items:       items,
	}
}

func (h OrderHandler) Get(ctx echo.Context) error {
	orders, err := h.orderUseCase.Get()
	if err != nil {
		h.log.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, "internal server error")
	}

	return ctx.JSON(http.StatusOK, orders)
}

func (h OrderHandler) Save(ctx echo.Context) error {
	var req SaveOrderReq

	if err := ctx.Bind(&req); err != nil {
		h.log.Error(err.Error())
		return ctx.JSON(http.StatusBadRequest, "unsupported structure")
	}

	or := req.OrderFromDTO()

	id, err := h.orderUseCase.Save(or)
	if err != nil {
		h.log.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, id)
}
