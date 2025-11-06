package dto

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


type TotalOrderDTO struct {

}