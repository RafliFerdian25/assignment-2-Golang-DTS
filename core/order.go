package core

import "time"

type Order struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	CustomerName string    `json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	Items        []Item
}

type OrderResponse struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	CustomerName string         `json:"customerName"`
	OrderedAt    time.Time      `json:"orderedAt"`
	Items        []ItemResponse `json:"items"`
}

type OrderImpl interface {
	GetAllOrders() ([]Order, error)
	CreateOrder() ([]Order, error)
	UpdateOrder() ([]Order, error)
	DeleteOrder() ([]Order, error)
}
