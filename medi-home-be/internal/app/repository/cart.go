package repository

import (
	"log"
	"medi-home-be/config"
	"medi-home-be/internal/app/model"
)

type CartRepository interface {
	//VIEW
	GetCartItemDetail(cart_id int64) ([]CartItemDetail, error)
	GetCartUser(user_id int64) (model.Cart, error)
	GetCartItem(cart_id int64) ([]model.CartItem, error)
	GetItem(cartitem_id int64) (model.CartItem, error)

	//CRUD
	AddCartItem(item model.CartItem) (model.CartItem, error)
	UpdateCartItem(item model.CartItem) error
	DeleteItem(cartitem_id int64) error

	FindInventory(inventory_id int64) (model.Inventory, error)
	FindBatchSelling(medicine_id int64) (model.BatchSelling, error)
	FindMedicine(medicine_id int64) (model.Medicine, error)
}

type cartRepository struct{}

func NewCartRepository() CartRepository {
	return &cartRepository{}
}

//VIEW

// trả về cart_id từ user_id
func (r *cartRepository) GetCartUser(user_id int64) (model.Cart, error) {
	var cart model.Cart
	err := config.DB.Where("user_id = ?", user_id).Find(&cart).Error
	return cart, err
}

// trả về cart_item theo cart_id
func (r *cartRepository) GetCartItem(cart_id int64) ([]model.CartItem, error) {
	var items []model.CartItem
	err := config.DB.Where("cart_id = ?", cart_id).Find(&items).Error
	return items, err
}

func (r *cartRepository) GetItem(cartitem_id int64) (model.CartItem, error) {
	var item model.CartItem
	err := config.DB.Where("cartitem_id = ?", cartitem_id).Find(&item).Error
	return item, err
}

type CartItemDetail struct {
	CartItemID int64   `json:"cartitem_id" gorm:"column:cartitem_id"`
	Name       string  `json:"name"`
	Thumbnail  string  `json:"thumbnail"`
	Quantity   int     `json:"quantity"`
	SelectType string  `json:"select_type"`
	Price      float64 `json:"price"`
	Total      float64 `json:"total"`
}

func (r *cartRepository) GetCartItemDetail(cart_id int64) ([]CartItemDetail, error) {
	var items []CartItemDetail
	query := `
        select ci.cartitem_id, m.name, m.thumbnail, ci.quantity, ci.select_type, ci.price,
		round(ci.quantity*ci.price) as total
		from cartitems ci
		join medicines m on m.medicine_id = ci.medicine_id
		join inventories i on i.medicine_id = ci.medicine_id
		join batchsellings bs on bs.inventory_id = i.inventory_id
		where cart_id = ?
	`
	err := config.DB.Raw(query, cart_id).Scan(&items).Error
	log.Print(items)
	return items, err
}

//CRUD

func (r *cartRepository) CreateCart(cart model.CartItem) (model.CartItem, error) {
	err := config.DB.Create(&cart).Error
	return cart, err
}

func (r *cartRepository) AddCartItem(item model.CartItem) (model.CartItem, error) {
	if err := config.DB.Create(&item).Error; err != nil {
		return model.CartItem{}, err
	}
	return item, nil
}

func (r *cartRepository) UpdateCartItem(item model.CartItem) error {
	// Cập nhật quantity và price (hoặc các field khác nếu cần)
	return config.DB.Model(&model.CartItem{}).
		Where("cartitem_id = ?", item.CartItemID).
		Updates(map[string]interface{}{
			"quantity":    item.Quantity,
			"price":       item.Price,
			"select_type": item.SelectType,
		}).Error
}

func (r *cartRepository) DeleteItem(cartitem_id int64) error {
	if err := config.DB.Delete(&model.CartItem{}, cartitem_id).Error; err != nil {
		return err
	}
	return nil
}

func (r *cartRepository) FindInventory(inventory_id int64) (model.Inventory, error) {
	var inventory model.Inventory
	err := config.DB.Where("inventory_id = ?", inventory_id).Find(&inventory).Error
	return inventory, err
}

func (r *cartRepository) FindBatchSelling(medicine_id int64) (model.BatchSelling, error) {
	var batch model.BatchSelling
	err := config.DB.Where("medicine_id = ?", medicine_id).Find(&batch).Error
	return batch, err
}

func (r *cartRepository) FindMedicine(medicine_id int64) (model.Medicine, error) {
	var medicine model.Medicine
	err := config.DB.Where("medicine_id = ?", medicine_id).Find(&medicine).Error
	return medicine, err
}
