package repository

import (
	"log"
	"medi-home-be/config"
	"medi-home-be/internal/app/model"
	"time"

	"gorm.io/gorm"
)

type OrderRepository interface {
	//View for admin
	GetViewAllOrder() ([]OrderList, error)
	GetViewOrder(order_id int64) (Order, error)
	GetViewOrderStatus(status string) ([]OrderList, error)
	GetViewOrderDetail(order_id int64) ([]OrderDetail, error)
	GetOrderID(order_id int64) (model.Order, error)

	//View for user
	GetViewAllOrderUser(user_id int64) ([]OrderList, error)
	GetViewOrderStatusByUser(user_id int64, status string) ([]OrderList, error)

	//CRUD
	CreateOrder(tx *gorm.DB, order model.Order) (model.Order, error)
	CreateOrderDetail(tx *gorm.DB, orderdetail model.OrderDetail) (model.OrderDetail, error)
	UpdateStatus(order_id int64, status string) (model.Order, error)
}

type orderRepository struct{}

func NewOrderRepository() OrderRepository {
	return &orderRepository{}
}

// View
type OrderList struct {
	OrderID       int64     `json:"order_id"`
	Date          time.Time `json:"date"`
	FullName      string    `json:"full_name"`
	OrderItem     int64     `json:"order_item"`
	OrderStatus   string    `json:"order_status"`
	PaymentMethod string    `json:"payment_method"`
	PaymentStatus string    `json:"payment_status"`
	FinalAmount   float64   `json:"final_amount"`
}

func (r *orderRepository) GetViewAllOrder() ([]OrderList, error) {
	var order []OrderList
	query := `
			select o.order_id,o.created_at,a.full_name,count(od.orderdetail_id) as order_item,
			o.order_status,o.payment_method,o.payment_status,o.final_amount
			from orders o
			join addresses a on a.address_id = o.address_id
			left join vouchers v on v.voucher_id = o.voucher_id
			join shippings s on s.shipping_id = o.shipping_id
			join orderdetails od on od.order_id = o.order_id
			group by 
				o.order_id,o.created_at,a.full_name,o.order_status,
				o.payment_method,o.payment_status,o.total_amount,o.final_amount
			order by o.order_id desc
		`
	err := config.DB.Raw(query).Scan(&order).Error
	return order, err
}

func (r *orderRepository) GetViewOrderStatus(status string) ([]OrderList, error) {
	var order []OrderList
	query := `
			select o.order_id,o.created_at,a.full_name,count(od.orderdetail_id) as order_item,
			o.order_status,o.payment_method,o.payment_status,o.final_amount
			from orders o
			join addresses a on a.address_id = o.address_id
			left join vouchers v on v.voucher_id = o.voucher_id
			join shippings s on s.shipping_id = o.shipping_id
			join orderdetails od on od.order_id = o.order_id
			where o.order_status = ? 
			group by 
				o.order_id,o.created_at,a.full_name,o.order_status,
				o.payment_method,o.payment_status,o.total_amount,o.final_amount,
			order by o.order_id desc
		`
	err := config.DB.Raw(query, status).Scan(&order).Error
	return order, err
}

func (r *orderRepository) GetViewAllOrderUser(user_id int64) ([]OrderList, error) {
	var order []OrderList
	query := `
			select o.order_id,o.created_at,a.full_name,count(od.orderdetail_id) as order_item,
			o.order_status,o.payment_method,o.payment_status,o.final_amount
			from orders o
			join addresses a on a.address_id = o.address_id
			left join vouchers v on v.voucher_id = o.voucher_id
			join shippings s on s.shipping_id = o.shipping_id
			join orderdetails od on od.order_id = o.order_id
			where o.user_id = ?
			group by 
				o.order_id,o.created_at,a.full_name,o.order_status,
				o.payment_method,o.payment_status,o.total_amount,o.final_amount
			order by o.order_id desc
			
		`
	err := config.DB.Raw(query, user_id).Scan(&order).Error
	return order, err
}

func (r *orderRepository) GetViewOrderStatusByUser(user_id int64, status string) ([]OrderList, error) {
	var order []OrderList
	query := `
			select o.order_id,o.created_at,a.full_name,count(od.orderdetail_id) as order_item,
			o.order_status,o.payment_method,o.payment_status,o.final_amount
			from orders o
			join addresses a on a.address_id = o.address_id
			join shippings s on s.shipping_id = o.shipping_id
			join orderdetails od on od.order_id = o.order_id
			where o.order_status = ? and o.user_id = ?
			group by 
				o.order_id,o.created_at,a.full_name,o.order_status,
				o.payment_method,o.payment_status,o.total_amount,o.final_amount
			order by o.order_id desc
		`
	err := config.DB.Raw(query, status, user_id).Scan(&order).Error
	// log.Print("repo",order)
	return order, err
}

type OrderDetail struct {
	MedicineName string  `json:"medicine_name"`
	UnitPrice    float64 `json:"unit_price"`
	Type         string  `json:"type"`
	Quantity     int64   `json:"quantity"`
}

func (r *orderRepository) GetViewOrderDetail(order_id int64) ([]OrderDetail, error) {
	var orderdetail []OrderDetail
	query := `
			select m.name as medicine_name, od.unit_price, od.select_type as type, od.quantity 
			from orderdetails od
			join orders o on o.order_id = od.order_id
			join medicines m on m.medicine_id = od.medicine_id
			where od.order_id = ?
		`
	err := config.DB.Raw(query, order_id).Scan(&orderdetail).Error
	return orderdetail, err
}

type Order struct {
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

func (r *orderRepository) GetViewOrder(order_id int64) (Order, error) {
	var order Order
	query := `
			select o.order_id, a.full_name, a.phone, a.address, v.code as voucher_code,
			s.name as shipping_name, o.order_status, o.payment_method, o.payment_status,
			o.total_amount, o.final_amount
			from orders o 
			join addresses a on a.address_id = o.address_id
			join vouchers v on v.voucher_id = o.voucher_id
			join shippings s on s.shipping_id = o.shipping_id
			where o.order_id = ?
		`
	err := config.DB.Raw(query, order_id).Scan(&order).Error
	return order, err
}

func (r *orderRepository) GetOrderID(order_id int64) (model.Order, error) {
	var order model.Order
	err := config.DB.Where("order_id = ?", order_id).First(&order).Error
	if err != nil {
		return order, err
	}
	log.Print("order", order)
	return order, nil
}

// CRUD
func (r *orderRepository) CreateOrder(tx *gorm.DB, order model.Order) (model.Order, error) {
	// useSelf := false
	// if tx == nil {
	// 	tx = config.DB.Begin()
	// 	useSelf = true
	// }

	// err := tx.Create(&order).Error
	// if err != nil {
	// 	if useSelf {
	// 		tx.Rollback()

	// 	}
	// 	return order, err
	// }
	if tx == nil {
		tx = config.DB
	}
	err := tx.Create(&order).Error
	return order, err
}

func (r *orderRepository) CreateOrderDetail(tx *gorm.DB, orderdetail model.OrderDetail) (model.OrderDetail, error) {
	if tx == nil {
		tx = config.DB
	}
	err := tx.Create(&orderdetail).Error
	return orderdetail, err
}

// func (r *orderRepository)

func (r *orderRepository) UpdateStatus(order_id int64, status string) (model.Order, error) {
	var order model.Order
	err := config.DB.Model(&order).
		Where("order_id = ?", order_id).
		Update("order_status", status).
		Error

	if err != nil {
		return order, err
	}

	err = config.DB.Where("order_id = ?", order_id).First(&order).Error

	return order, err
}
