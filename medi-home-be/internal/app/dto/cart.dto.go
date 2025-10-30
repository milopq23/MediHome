package dto

type CartItemDetail struct {
    Name       string  `json:"name"`
    Quantity   int     `json:"quantity"`
    PriceStrip float64 `json:"price_strip"`
    PriceBox   float64 `json:"price_box"`
}