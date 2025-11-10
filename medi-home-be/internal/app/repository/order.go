package repository

import (
	"medi-home-be/config"
	"medi-home-be/internal/app/model"
)

type OrderRepository interface {
	CreateOrder(order model.Order) (model.Order, error)
	CreateOrderDetail(orderdetail model.OrderDetail) (model.OrderDetail, error)
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

// func(r *orderRepository)

func (r *orderRepository) CreateOrderDetail(orderdetail model.OrderDetail) (model.OrderDetail, error) {
	err := config.DB.Create(&orderdetail).Error
	return orderdetail, err
}
