package order

import (
	"fmt"
	orderEntity "fokoto/internal/domain/model/order"
	itemRepo "fokoto/internal/domain/repository/item"
	orderRepo "fokoto/internal/domain/repository/order"
)

type OrderUseCase struct {
	orderRepo orderRepo.OrderRepository
	itemRepo  itemRepo.ItemRepository
}

func NewOrderUseCase(orderRepo orderRepo.OrderRepository, itemRepo itemRepo.ItemRepository) *OrderUseCase {
	return &OrderUseCase{
		orderRepo: orderRepo,
		itemRepo:  itemRepo,
	}
}

func (u *OrderUseCase) Save(order orderEntity.Order) (uint64, error) {
	orderId, err := u.orderRepo.Save(order)
	if err != nil {
		return 0, err
	}

	if err = u.itemRepo.SaveAll(orderId, order.Items); err != nil {
		return 0, err
	}

	return orderId, nil
}

func (u *OrderUseCase) Get() ([]orderEntity.Order, error) {
	orders, err := u.orderRepo.Get()
	if err != nil {
		return nil, err
	}

	for i, order := range orders {
		orders[i].Items, err = u.itemRepo.Get(order.ID)
		if err != nil {
			return nil, err
		}

		for _, item := range order.Items {
			order.OriginalAmount = order.OriginalAmount + item.Amount
			order.DiscountedAmount = order.DiscountedAmount + item.DiscountedAmount
		}
	}

	for _, order := range orders {
		fmt.Println(order)
	}

	return orders, nil
}
