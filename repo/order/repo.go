package order

import (
	"assignment-2/core"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) GetAllOrders() ([]core.OrderResponse, error) {
	orders := []core.OrderResponse{}

	err := r.db.Model(&core.Order{}).Preload("Items").Find(&orders).Error

	return orders, err
}
