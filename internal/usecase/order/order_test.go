package order

import (
	"fokoto/internal/domain/model/item"
	"fokoto/internal/domain/model/order"
	itemmock "fokoto/internal/domain/repository/item/mocks"
	ordermock "fokoto/internal/domain/repository/order/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestOrderUseCase_Save(t *testing.T) {
	orderRepoMock := ordermock.NewOrderRepository(t)
	itemRepoMock := itemmock.NewItemRepository(t)

	orderUseCase := NewOrderUseCase(orderRepoMock, itemRepoMock)

	ordr := order.Order{
		ID:          0,
		Status:      1,
		UserID:      0,
		PaymentType: 1,
		Items: []item.Item{
			{ID: 0, OrderID: 0, Amount: 100, DiscountedAmount: 0},
		},
	}

	orderRepoMock.
		On("Save", ordr).
		Once().
		Return(ordr.ID, nil)

	itemRepoMock.
		On("SaveAll", ordr.ID, ordr.Items).
		Once().
		Return(nil)

	orderId, err := orderUseCase.Save(ordr)

	assert.NoError(t, err)
	assert.Equal(t, ordr.ID, orderId)

	orderRepoMock.AssertExpectations(t)
	itemRepoMock.AssertExpectations(t)
}

func TestOrderUseCase_Get(t *testing.T) {
	orderRepoMock := new(ordermock.OrderRepository)
	itemRepoMock := new(itemmock.ItemRepository)

	orderUseCase := NewOrderUseCase(orderRepoMock, itemRepoMock)

	orderRepoMock.On("Get").Return([]order.Order{
		{
			ID:          0,
			Status:      1,
			UserID:      0,
			PaymentType: 1,
		},
	}, nil)
	itemRepoMock.On("Get", mock.AnythingOfType("uint64")).Return([]item.Item{}, nil)

	_, err := orderUseCase.Get()

	assert.NoError(t, err)

	orderRepoMock.AssertExpectations(t)
	itemRepoMock.AssertExpectations(t)
}
