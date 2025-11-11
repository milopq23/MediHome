package repository

import (
	"medi-home-be/config"
	"medi-home-be/internal/app/model"
)

type OrderRepository interface {
	CreateOrder(order model.Order) (model.Order, error)
	CreateOrderDetail(orderdetail model.OrderDetail) (model.OrderDetail, error)
	GetAllOrder() ([]OrderList, error)
	GetStatusTypeOrder(status string) ([]OrderList, error)
}

type orderRepository struct{}

func NewOrderRepository() OrderRepository {
	return &orderRepository{}
}

func (r *orderRepository) GetOrder() (model.Order, error) {
	var order model.Order
	err := config.DB.Model(&order).Error
	return order, err
}

func (r *orderRepository) CreateOrder(order model.Order) (model.Order, error) {
	err := config.DB.Create(&order).Error
	return order, err
}

func (r *orderRepository) CreateOrderDetail(orderdetail model.OrderDetail) (model.OrderDetail, error) {
	err := config.DB.Create(&orderdetail).Error
	return orderdetail, err
}

type OrderList struct {
	OrderID       int64   `json:"order_id"`
	FullName      string  `json:"full_name"`
	Phone         string  `json:"phone"`
	Address       string  `json:"address"`
	VoucherCode   string  `json:"voucher_code"`
	ShippingName  string  `json:"shipping_name"`
	OrderStatus   string  `json:"order_status"`
	PaymentMethod string  `json:"payment_method"`
	PaymentStatus string  `json:"payment_status"`
	TotalAmount   float64 `json:"total_amount"`
	FinalAmount   float64 `json:"final_amount"`
}

func (r *orderRepository) GetAllOrder() ([]OrderList, error) {
	var order []OrderList
	query := `
		select o.order_id, a.full_name, a.phone, a.address, v.code as voucher_code, s.name as shipping_name,
		o.order_status, o.payment_method, o.payment_status,o.total_amount, o.final_amount
		from orders o
		join addresses a on a.address_id = o.address_id
		join vouchers v on v.voucher_id = o.voucher_id
		join shippings s on s.shipping_id = o.shipping_id
		order by o.order_id desc`
	err := config.DB.Raw(query).Scan(&order).Error
	return order, err
}

func (r *orderRepository) GetStatusTypeOrder(status string) ([]OrderList, error) {
	var order []OrderList
	query := `
		select o.order_id, a.full_name, a.phone, a.address, v.code as voucher_code, s.name as shipping_name,
		o.order_status, o.payment_method, o.payment_status,o.total_amount, o.final_amount
		from orders o
		join addresses a on a.address_id = o.address_id
		join vouchers v on v.voucher_id = o.voucher_id
		join shippings s on s.shipping_id = o.shipping_id
		where o.order_status = ?
		order by o.order_id asc`
	err := config.DB.Raw(query, status).Scan(&order).Error
	return order, err
}


// func (r *orderRepository) ApproveStatus()