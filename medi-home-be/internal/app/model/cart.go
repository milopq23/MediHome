package model

type Cart struct {
	CartID     int64   `json:"cart_id" gorm:"primaryKey;autoIncrement;column:cart_id"`
	UserID     int64   `json:"user_id" gorm:"column:user_id"`
}

func (Cart) TableName() string { 
	return "cart"
}

type CartItems struct {
	CartItemID int64 `json:"cart_item_id" gorm:"primaryKey;autoIncrement;column:cart_item_id"`	
	CartID     int64 `json:"cart_id" gorm:"column:cart_id"`
	MedicineID int64 `json:"medicine_id" gorm:"column:medicine_id"`
	Quantity   int64 `json:"quantity" gorm:"column:quantity"`
}
func (CartItems) TableName() string { 
	return "cart_items"
}


type CartItemDetail struct {
	Name       string  `json:"name"`
	Quantity   int     `json:"quantity"`
	PriceStrip float64 `json:"price_strip"`
	PriceBox   float64 `json:"price_box"`
}