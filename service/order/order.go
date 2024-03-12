package order

import (
	"assignment-2/core"
	"assignment-2/repo/order"

	"github.com/jinzhu/copier"
)

type Service struct {
	orderRepo *order.OrderRepository
}

func NewService(orderRepo *order.OrderRepository) *Service {
	return &Service{
		orderRepo: orderRepo,
	}
}

func (s *Service) GetAllOrders() ([]core.OrderResponse, error) {
	orders, err := s.orderRepo.GetAllOrders()
	if err != nil {
		return nil, err
	}

	var getOrders []core.OrderResponse
	err = copier.Copy(&getOrders, &orders)
	if err != nil {
		return nil, err
	}

	return getOrders, err
}

func (s *Service) CreateOrder(OrderRequest core.OrderRequest) (*core.OrderResponse, error) {
	var order core.Order
	err := copier.Copy(&order, &OrderRequest)
	if err != nil {
		return nil, err
	}

	createdOrder, err := s.orderRepo.CreateOrder(order)
	if err != nil {
		return nil, err
	}

	var getOrders core.OrderResponse
	err = copier.Copy(&getOrders, createdOrder)
	if err != nil {
		return nil, err
	}

	return &getOrders, nil
}
