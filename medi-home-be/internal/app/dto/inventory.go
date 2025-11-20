package dto

import "time"

type ListInventoryDTO struct {
	Code           string   `json:"code"`
	Name           string   `json:"name"`
	Thumbnail      string   `json:"thumbnail"`
	BatchNumber    string   `json:"batch_number"`
	ImportPrice    float64  `json:"import_price"`
	MarkUp         float64  `json:"mark_up"`
	SellingPrice   float64  `json:"selling_price"`
	ExpirationDate time.Time `json:"expiration_date"`
	Quantity       int64    `json:"quantity"`
	Status         string   `json:"status"`
}
