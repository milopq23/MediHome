package model

type Cart struct {
	CartID     int64   `json:"cart_id" gorm:"primaryKey;autoIncrement;column:cart_id"`
	UserID     int64   `json:"user_id" gorm:"column:user_id"`
}

func (Cart) TableName() string { 
	return "cart"
}

type CartItem struct {
	CartItemID int64 `json:"cartitem_id" gorm:"primaryKey;autoIncrement;column:cartitem_id"`	
	CartID     int64 `json:"cart_id" gorm:"column:cart_id"`
	MedicineID int64 `json:"medicine_id" gorm:"column:medicine_id"`
	Quantity   int64 `json:"quantity" gorm:"column:quantity"`
}
func (CartItem) TableName() string { 
	return "cartitems"
}
