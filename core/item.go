package core

type Item struct {
	ID          uint   `gorm:"primaryKey"`
	Code        string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	Order_id    uint
}

type ItemResponse struct {
	Code        string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
}
