package service

import (
	"medi-home-be/internal/app/dto"
	"medi-home-be/internal/app/model"
	"medi-home-be/internal/app/repository"
)

type InventoryService interface {
	GetAll() ([]model.Inventory, error)
	Create(inventory model.Inventory) (model.Inventory, error)
	Delete(id int64) error
	GetInventory() ([]dto.ListInventoryDTO, error)
	FindMedicine(medicine_id int64) ([]model.Inventory, error)
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

func (s *inventoryService) FindMedicine(medicine_id int64) ([]model.Inventory, error) {
	return s.repo.FindMedicine(medicine_id)
}

func (s *inventoryService) GetInventory() ([]dto.ListInventoryDTO, error) {
	inventories, err := s.repo.GetInventory()
	if err != nil {
		return []dto.ListInventoryDTO{}, err
	}

	var result []dto.ListInventoryDTO
	for _, inv := range inventories {
		result = append(result, dto.ListInventoryDTO{
			Code:           inv.Code,
			Name:           inv.Name,
			Thumbnail:      inv.Thumbnail,
			BatchNumber:    inv.BatchNumber,
			ImportPrice:    inv.ImportPrice,
			MarkUp:         inv.MarkUp,
			SellingPrice:   inv.SellingPrice,
			ExpirationDate: inv.ExpirationDate,
			Quantity:       inv.Quantity,
			Status:         inv.Status,
		})
	}

	return result, err
}

// func (s *inventoryService) Patch(id int64, data map[string]interface{}) (model.Inventory, error){
// 	inventory, err := s.repo.FindByID(id)
// 	return s.repo.Patch(inventory)
// }

func (s *inventoryService) Delete(id int64) error {
	return s.repo.Delete(id)
}
