package repository

import (
	"fmt"
	"medi-home-be/config"
	"medi-home-be/internal/app/model"
)

type InventoryRepository interface {
	// ADMIN
	FindAll() ([]model.Inventory, error)
	Create(inventory model.Inventory) (model.Inventory, error)
	DecreaseQuantity(inventory_id, quantity int64) error
	// Patch(id int64, updates map[string]interface{}) (model.Inventory, error)
	Delete(id int64) error
}

type inventoryRepository struct{}

func NewInventoryRepository() InventoryRepository {
	return &inventoryRepository{}
}

func (r *inventoryRepository) FindAll() ([]model.Inventory, error) {
	var inventories []model.Inventory
	err := config.DB.Find(&inventories).Error
	return inventories, err
}

func (r *inventoryRepository) Create(inventory model.Inventory) (model.Inventory, error) {
	err := config.DB.Create(&inventory).Error
	return inventory, err
}

// func (r *inventoryRepository) Patch(id int64, updates map[string]interface{}) (model.Inventory, error) {
// 	err := config.DB.Model(&model.Inventory{}).Where("inventory_id", id).Updates(updates).Error
// 	if err != nil {
// 		return model.Inventory{}, err
// 	}
// 	var updated model.Inventory
// 	err = config.DB.First(&updated, id).Error
// 	return updated, err
// }

func (r *inventoryRepository) Delete(id int64) error {
	err := config.DB.Delete(&model.Inventory{}, id).Error
	return err
}

func (r *inventoryRepository) FindByID(id int64) (model.Inventory, error) {
	var inventory model.Inventory
	err := config.DB.First(&inventory, id).Error
	return inventory, err
}

func (r *inventoryRepository) SelectInventory(id int64) (model.Inventory, error) {
	var inventory model.Inventory
	err := config.DB.Where("inventory_id = ?", id).Find(&inventory).Error
	return inventory, err
}

func (r *inventoryRepository) DecreaseQuantity(inventory_id, quantity int64) error {
	var inventory model.Inventory
	if err := config.DB.Where("inventory_id = ?", inventory_id).First(&inventory).Error; err != nil {
		return err
	}

	if inventory.Quantity < quantity {
		return fmt.Errorf("not enough stock: have %d, need %d", inventory.Quantity, quantity)
	}

	inventory.Quantity -= quantity

	if err := config.DB.Save(&inventory).Error; err != nil {
		return err
	}

	return nil
}
