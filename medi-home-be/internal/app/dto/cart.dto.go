package dto

type CartItemDetailDTO struct {
	Name       string  `json:"name"`
	Thumbnail  string  `json:"thumbnail"`
	Quantity   int     `json:"quantity"`
	PriceStrip float64 `json:"price_strip"`
	PriceBox   float64 `json:"price_box"`
}

type CartItemResponseDTO struct {
	Name       string  `json:"name"`
	Thumbnail  string  `json:"thumbnail"`
	Quantity   int     `json:"quantity"`
	SelectType string  `json:"select_type"`
	Price      float64 `json:"price"`
}

// type UpdateCartItemRequestDTO struct{
// 	CartID
// }

type CartRequestDTO struct {
	CartID     int64  `json:"cart_id"`
	CartItemID int64  `json:"cartitem_id"`
	MedicineID int64  `json:"medicine_id"`
	Quantity   int64  `json:"quantity"`
	SelectType string `json:"select_type"`
}

type SelectTypeMedicineDTO struct {
	MedicineID int64   `json:"medicine_id"`
	SelectType string  `json:"select_type"`
	Price      float64 `json:"price"`
}
