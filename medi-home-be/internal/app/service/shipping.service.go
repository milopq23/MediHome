package service

import (
	"medi-home-be/internal/app/model"
	"medi-home-be/internal/app/repository"
)

type ShippingService interface {
	GetAll() ([]model.Shipping, error)
	GetByID(id int64) (model.Shipping, error)
	Create(shipping model.Shipping) (model.Shipping, error)
	Patch(id int64, data map[string]interface{}) (model.Shipping, error)
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

func (s *shippingService) Patch(id int64, data map[string]interface{}) (model.Shipping, error) {
	ship, err := s.repo.FindByID(id)
	if err != nil {
		return model.Shipping{}, err
	}

	allowedFields := map[string]bool{
		"name":  true,
		"price": true,
	}

	updates := make(map[string]interface{})

	for k, v := range data {
		if allowedFields[k] {
			updates[k] = v
		}
	}

	return s.repo.Patch(int64(ship.ShippingID), updates)
}

func (s *shippingService) Delete(id int64) error {
	return s.repo.Delete(id)
}
