package handler

import (
	"fokoto/internal/usecase/order"
	"github.com/labstack/echo/v4"
	"log/slog"
)

type OrderHandler struct {
	orderUseCase *order.OrderUseCase
	log          *slog.Logger
}

func NewOrderHandler(e *echo.Echo, orderUseCase *order.OrderUseCase, log *slog.Logger) *OrderHandler {
	handler := OrderHandler{
		orderUseCase: orderUseCase,
		log:          log,
	}

	orders := e.Group("api/v1/orders")
	orders.GET("/", handler.Get)
	orders.POST("/", handler.Save)

	return &handler
}
