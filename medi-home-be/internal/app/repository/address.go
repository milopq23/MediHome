package repository

import (
	"log"
	"medi-home-be/config"
	"medi-home-be/internal/app/model"
)

type AddressRepository interface {
	GetAddressUser(user_id int64) (model.User, error)
	GetAddress(user_id int64) ([]model.Address, error)
	CreateAddressForUser(address model.Address) (model.Address, error)
	GetByID(id int64) (model.Address, error)
}

type addressRepository struct{}

func NewAddressRepository() AddressRepository {
	return &addressRepository{}
}

func (r *addressRepository) GetAddressUser(user_id int64) (model.User, error) {
	var user model.User
	err := config.DB.Where("user_id = ? ", user_id).Find(&user).Error
	log.Print(user)
	return user, err
}

func (r *addressRepository) CreateAddressForUser(address model.Address) (model.Address, error) {
	err := config.DB.Create(&address).Error
	return address, err
}

func (r *addressRepository) GetByID(id int64) (model.Address, error) {
	var address model.Address
	err := config.DB.Find(&address).Error
	return address, err
}

func (r *addressRepository) GetAddress(user_id int64) ([]model.Address, error) {
	var address []model.Address
	err := config.DB.Where("user_id = ?", user_id).Find(&address).Error
	return address, err
}
