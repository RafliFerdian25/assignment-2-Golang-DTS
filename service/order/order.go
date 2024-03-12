package order

import (
	"assignment-2/core"
	"assignment-2/repo/order"
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
	res, err := s.orderRepo.GetAllOrders()

	return res, err
}
