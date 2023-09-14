package usecase

import (
	"github.com/caiocp/clean-arch-go/internal/entity"
)

type FindOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewFindOrdersUseCase(orderRepository entity.OrderRepositoryInterface) *FindOrdersUseCase {
	return &FindOrdersUseCase{
		OrderRepository: orderRepository,
	}
}

func (c *FindOrdersUseCase) Execute() ([]entity.Order, error) {
	orders, err := c.OrderRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return orders, nil
}
