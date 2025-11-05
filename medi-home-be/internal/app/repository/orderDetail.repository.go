package repository

import (
	"medi-home-be/config"
	"medi-home-be/internal/app/model"
)

type OrderDetailRepository interface {
}

type orderDetailRepository struct {
}

func NewOrderDetailRepository() OrderDetailRepository {
	return &orderDetailRepository{}
}

func (r *orderDetailRepository) CreateOrderDetail(orderdetail model.OrderDetail) (model.OrderDetail, error) {
	err := config.DB.Create(&orderdetail).Error
	return orderdetail, err
}
