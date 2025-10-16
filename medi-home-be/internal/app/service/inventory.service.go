package service

import (
	"medi-home-be/internal/app/model"
	"medi-home-be/internal/app/repository"
)

type InventoryService interface {
	GetAll() ([]model.Inventory, error)
	Create(inventory model.Inventory) (model.Inventory, error)
	Delete(id int64) error
}

type inventoryService struct {
	repo repository.InventoryRepository
}

func NewInventoryService(repo repository.InventoryRepository) InventoryService {
	return &inventoryService{repo}
}

func (s *inventoryService) GetAll() ([]model.Inventory, error) {
	return s.repo.FindAll()
}

func (s *inventoryService) Create(inventory model.Inventory) (model.Inventory, error) {
	return s.repo.Create(inventory)
}

// func (s *inventoryService) Patch(id int64, data map[string]interface{}) (model.Inventory, error){
// 	inventory, err := s.repo.FindByID(id)
// 	return s.repo.Patch(inventory)
// }

func (s *inventoryService) Delete(id int64) error{
	return s.repo.Delete(id)
}