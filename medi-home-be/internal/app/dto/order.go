package dto

import "time"

type OrderDTO struct {
	OrderID       int64   `json:"order_id"`
	UserID        int64   `json:"user_id"`
	VoucherCode   string  `json:"voucher_code"`
	ShippingID    int64   `json:"shipping_id"`
	OrderStatus   string  `json:"order_status"`
	PaymentMethod string  `json:"payment_method"`
	PaymentStatus string  `json:"payment_status"`
	TotalAmount   float64 `json:"total_amount"`
	FinalAmount   float64 `json:"final_amount"`
}

type OrderRequestDTO struct {
	UserID        int64  `json:"user_id"`
	VoucherCode   string `json:"voucher_code"`
	ShippingID    int64  `json:"shipping_id"`
	PaymentMethod string `json:"payment_method"`
}

type OrderDetailItemDTO struct {
	Name     string  `json:"name"`
	Quantity int64   `json:"quantity"`
	Price    float64 `json:"price"`
}

type AllOrderResponse struct {
	OrderID       int64     `json:"order_id"`
	Date          time.Time `json:"date"`
	FullName      string    `json:"full_name"`
	OrderItem     int64     `json:"order_item"`
	OrderStatus   string    `json:"order_status"`
	PaymentMethod string    `json:"payment_method"`
	PaymentStatus string    `json:"payment_status"`
	FinalAmount   float64   `json:"final_amount"`
}

type OrderResponse struct {
	OrderID       int64                 `json:"order_id"`
	FullName      string                `json:"full_name"`
	Phone         string                `json:"phone"`
	Address       string                `json:"address"`
	VoucherCode   string                `json:"voucher_code"`
	ShippingName  string                `json:"shipping_name"`
	OrderStatus   string                `json:"order_status"`
	PaymentMethod string                `json:"payment_method"`
	PaymentStatus string                `json:"payment_status"`
	TotalAmount   float64               `json:"total_amount"`
	FinalAmount   float64               `json:"final_amount"`
	OrderDetail   []OrderDetailResponse `json:"order_detail"`
}

type OrderDetailResponse struct {
	MedicineName string  `json:"medicine_name"`
	UnitPrice    float64 `json:"unit_price"`
	Type         string  `json:"type"`
	Quantity     int64   `json:"quantity"`
}
