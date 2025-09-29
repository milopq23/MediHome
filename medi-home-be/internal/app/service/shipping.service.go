package service

import (
	"medi-home-be/internal/app/model"
	"medi-home-be/internal/app/repository"
)

type ShippingService interface {
	GetAll() ([]model.Shipping, error)
	GetByID(id int64) (model.Shipping, error)
	Create(shipping model.Shipping) (model.Shipping, error)
	Update(shipping model.Shipping) (model.Shipping, error)
	Patch(shipping model.Shipping) (model.Shipping, error)
	Delete(id int64) error
}

type shippingService struct {
	repo repository.ShippingRepository
}

func NewShippingService(repo repository.ShippingRepository) ShippingService {
	return &shippingService{}
}

func (s *shippingService) GetAll() ([]model.Shipping, error) {
	return s.repo.FindAll()
}

func (s *shippingService) GetByID(id int64) (model.Shipping, error) {
	return s.repo.FindByID(id)
}

func (s *shippingService) Create(shipping model.Shipping) (model.Shipping, error) {
	return s.repo.Create(shipping)
}

func (s *shippingService) Update(shipping model.Shipping) (model.Shipping, error) {
	return s.repo.Update(shipping)
}

func (s *shippingService) Patch(shipping model.Shipping) (model.Shipping, error) {
	return s.repo.Patch(shipping)
}

func (s *shippingService) Delete(id int64) error {
	return s.repo.Delete(id)
}