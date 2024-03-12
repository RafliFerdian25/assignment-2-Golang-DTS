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

func (r *OrderRepository) GetAllOrders() ([]core.Order, error) {
	orders := []core.Order{}

	err := r.db.Preload("Items").Find(&orders).Error
	if err != nil {
		return nil, err
	}

	return orders, err
}

func (r *OrderRepository) CreateOrder(order core.Order) (*core.Order, error) {
	// create order
	tx := r.db.Begin()
	err := r.db.Model(&core.Order{}).Create(&order).Error
	if err != nil {
		tx.Rollback()
		return &core.Order{}, err
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &order, err
}

func (r *OrderRepository) DeleteOrder(orderId uint) error {
	// delete order
	tx := r.db.Begin()
	err := r.db.Delete(&core.Order{}, orderId)
	if err.Error != nil {
		tx.Rollback()
		return err.Error
	}

	if err.RowsAffected <= 0 {
		tx.Rollback()
		return gorm.ErrRecordNotFound
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// update
func (r *OrderRepository) UpdateOrder(order core.Order) (*core.Order, error) {
	tx := r.db.Begin()
	// delete order items
	err := r.db.Where("order_id = ?", order.ID).Delete(&core.Item{})
	if err.Error != nil {
		return nil, err.Error
	}

	// update order
	err = r.db.Save(&order)
	if err.Error != nil {
		tx.Rollback()
		return &core.Order{}, err.Error
	}
	if err.RowsAffected <= 0 {
		tx.Rollback()
		return &core.Order{}, gorm.ErrRecordNotFound
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &order, nil
}
