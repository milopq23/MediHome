package service

import (
	"medi-home-be/internal/app/model"
	"medi-home-be/internal/app/repository"
)

type AddressService interface {
	AddAddress(user_id int64, address model.Address) (model.Address, error)
	GetAddressUser(user_id int64) (model.User, error)
}

type addressService struct {
	repo repository.AddressRepository
}

func NewAddressService(repo repository.AddressRepository) AddressService {
	return &addressService{repo}
}

func (s *addressService) AddAddress(user_id int64, address model.Address) (model.Address, error) {
	userAddress := model.Address{
		UserID:    user_id,
		Phone:     address.Phone,
		FullName:  address.FullName,
		Address:   address.Address,
		IsDefault: false,
	}
	return s.repo.CreateAddressForUser(userAddress)
}

func (s *addressService) GetAddressUser(user_id int64) (model.User, error) {
	return s.repo.GetAddressUser(user_id)
}
