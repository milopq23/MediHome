package model

import "time"

type Order struct{
	OrderID int64
	UserID int64
	AddressID int64
	VoucherID int64
	ShippingID int64
	OrderStatus string
	PaymentMethod string
	PaymentStatus string
	TotalAmount int64
	FinalAmount int64
	CreatedAt time.Time
	UpdatedAt time.Time
}