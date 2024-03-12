package core

import "time"

type Order struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	CustomerName string    `json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	Items        []Item    `json:"items" gorm:"foreignkey:order_id;references:id"`
}

type OrderResponse struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	CustomerName string         `json:"customerName"`
	OrderedAt    time.Time      `json:"orderedAt"`
	Items        []ItemResponse `json:"items" gorm:"foreignkey:order_id;references:id"`
}

type OrderRequest struct {
	CustomerName string        `json:"customerName"`
	OrderedAt    time.Time     `json:"orderedAt"`
	Items        []ItemRequest `json:"items"`
}
