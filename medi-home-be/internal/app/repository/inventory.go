package repository

import (
	"fmt"
	"medi-home-be/config"
	"medi-home-be/internal/app/model"
	"time"

	"gorm.io/gorm"
)

type InventoryRepository interface {
	GetInventory() ([]ListInventory, error)
	// ADMIN
	FindAll() ([]model.Inventory, error)
	FindMedicine(medicine_id int64) ([]model.Inventory, error)
	Create(inventory model.Inventory) (model.Inventory, error)
	DecreaseQuantity(tx *gorm.DB, inventory_id, quantity int64) error
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

func (r *inventoryRepository) FindByID(id int64) (model.Inventory, error) {
	var inventory model.Inventory
	err := config.DB.First(&inventory, id).Error
	return inventory, err
}

func (r *inventoryRepository) FindMedicine(medicine_id int64) ([]model.Inventory, error) {
	var inventory []model.Inventory
	err := config.DB.Where("medicine_id = ?", medicine_id).Find(&inventory).Error
	return inventory, err
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

type ListInventory struct {
	Code           string    `json:"code"`
	Name           string    `json:"name"`
	Thumbnail      string    `json:"thumbnail"`
	BatchNumber    string    `json:"batch_number"`
	ImportPrice    float64   `json:"import_price"`
	MarkUp         float64   `json:"mark_up"`
	SellingPrice   float64   `json:"selling_price"`
	ExpirationDate time.Time `json:"expiration_date"`
	Quantity       int64     `json:"quantity"`
	Status         string    `json:"status"`
}

func (r *inventoryRepository) GetInventory() ([]ListInventory, error) {
	var inventory []ListInventory
	query := `select m.code, m.name, m.thumbnail, i.batch_number, i.expiration_date , i.import_price, i.mark_up,
		round(i.import_price * (1 + i.mark_up/100)) as selling_price,i.quantity, i.status
		from batchsellings bs
		join inventories i on i.inventory_id = bs.inventory_id
		join medicines m on m.medicine_id = i.medicine_id
		order by i.quantity asc`
	err := config.DB.Raw(query).Scan(&inventory).Error
	return inventory, err
}

func (r *inventoryRepository) SelectInventory(id int64) (model.Inventory, error) {
	var inventory model.Inventory
	err := config.DB.Where("inventory_id = ?", id).Find(&inventory).Error
	return inventory, err
}

func (r *inventoryRepository) DecreaseQuantity(tx *gorm.DB, inventory_id, quantity int64) error {
	if tx == nil {
		tx = config.DB
	}
	var inventory model.Inventory
	if err := config.DB.Where("inventory_id = ?", inventory_id).First(&inventory).Error; err != nil {
		return err
	}

	if inventory.Quantity < quantity {
		if err := r.UpdateStatus(inventory_id, "Hết hàng"); err != nil {
			return err
		}
		return fmt.Errorf("not enough stock: have %d, need %d", inventory.Quantity, quantity)
	}

	inventory.Quantity -= quantity

	if inventory.Quantity == 0 {
		inventory.Status = "Hết hàng"
	}

	if err := tx.Save(&inventory).Error; err != nil {
		return err
	}

	return nil
}

func (r *inventoryRepository) UpdateStatus(inventoryID int64, status string) error {
	return config.DB.Model(&model.Inventory{}).
		Where("inventory_id = ?", inventoryID).
		Update("status", status).Error
}
