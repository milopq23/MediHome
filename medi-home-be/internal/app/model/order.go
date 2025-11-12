package model

type Order struct {
	OrderID       int64   `json:"order_id" gorm:"primaryKey; autoIncrement; column:order_id"`
	UserID        int64   `json:"user_id" gorm:"column:user_id"`
	AddressID     int64   `json:"address_id" gorm:"column:address_id"`
	VoucherID     int64   `json:"voucher_id" gorm:"column:voucher_id; default:null"`
	ShippingID    int64   `json:"shipping_id" gorm:"column:shipping_id"`
	OrderStatus   string  `json:"order_status" gorm:"order_status"`
	PaymentMethod string  `json:"payment_method" gorm:"payment_method"`
	PaymentStatus string  `json:"payment_status" gorm:"payment_status"`
	TotalAmount   float64 `json:"total_amount" gorm:"total_amount"`
	FinalAmount   float64 `json:"final_amount" gorm:"final_amount"`
}

func (Order) TableName() string {
	return "orders"
}
