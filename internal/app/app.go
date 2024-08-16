package app

import (
	"fokoto/internal/app/server"
	"fokoto/internal/config"
	"fokoto/internal/handler"
	"fokoto/internal/repository/item"
	"fokoto/internal/repository/order"
	"fokoto/internal/repository/psqldb"
	usecase "fokoto/internal/usecase/order"
	"github.com/labstack/echo/v4"
	"log/slog"
)

type App struct {
	Server *server.Server
}

func New(log *slog.Logger, cfg *config.Config) *App {
	storage, err := psqldb.New(cfg.DB)
	if err != nil {
		panic(err)
	}

	e := echo.New()

	orderRepo := order.NewOrderRepo(storage)
	itemRepo := item.NewItemRepo(storage)
	orderUseCase := usecase.NewOrderUseCase(orderRepo, itemRepo)
	handler.NewOrderHandler(e, orderUseCase, log)

	srv := server.NewServer(cfg.Server, e)

	return &App{
		Server: srv,
	}
}
