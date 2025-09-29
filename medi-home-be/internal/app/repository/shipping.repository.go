package repository

import (
	"medi-home-be/config"
	"medi-home-be/internal/app/model"
)

type ShippingRepository interface {
	FindAll() ([]model.Shipping,error)
	FindByID(id int64) (model.Shipping, error)
	Create(shipping model.Shipping) (model.Shipping, error)
	Update(shipping model.Shipping) (model.Shipping, error)
	Patch(shipping model.Shipping) (model.Shipping, error)
	Delete(id int64) error
}

type shippingRepository struct{}

func NewShippingRepository() ShippingRepository {
	return &shippingRepository{}
}

func (r *shippingRepository)  FindAll() ([]model.Shipping,error) {
	var shippings []model.Shipping
	err := config.DB.Find(&shippings).Error
	return shippings, err
}

func (r *shippingRepository) FindByID(id int64) (model.Shipping, error) {
	var shipping model.Shipping
	err := config.DB.Find(&shipping, id).Error
	return shipping, err
}

func (r *shippingRepository) Create(shipping model.Shipping) (model.Shipping, error) {
	err := config.DB.Create(&shipping).Error
	return shipping, err
}

func (r *shippingRepository) Update(shipping model.Shipping) (model.Shipping, error) {
	err := config.DB.Save(&shipping).Error
	return shipping, err
}

func (r *shippingRepository) Patch(shipping model.Shipping) (model.Shipping, error) {
	err := config.DB.Save(&shipping).Error
	return shipping, err
}

func (r *shippingRepository) Delete(id int64) error {
	err := config.DB.Delete(&model.Shipping{}, id).Error
	return err
}