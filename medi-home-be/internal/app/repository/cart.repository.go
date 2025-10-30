package repository

import (
	"medi-home-be/config"
	"medi-home-be/internal/app/model"
)

type CartRepository interface {
	GetCartUser(user_id int64) (model.Cart, error)
	GetCartItems(cart_id int64) ([]model.CartItemDetail, error)
}

type cartRepository struct{}

func NewCartRepository() CartRepository {
	return &cartRepository{}
}

// func (r *cartRepository) GetAll() ([]model.Cart, error) {

// }

// đầu tiên lấy từ model cart truyền user vào ra được cart
// sau đó từ cart lấy đc cart item truyền cart id vào để lấy đc list cart item
func (r *cartRepository) GetCartUser(user_id int64) (model.Cart, error) {
	var cart model.Cart
	err := config.DB.Where("user_id = ?", user_id).Find(&cart).Error
	return cart, err
}

func (r *cartRepository) CreateCart(cart_id model.CartItems) (model.CartItems, error) {
	err := config.DB.Create(&cart_id).Error
	return cart_id, err
}

func (r *cartRepository) GetCartItems(cart_id int64) ([]model.CartItemDetail, error) {
	var items []model.CartItemDetail
	query := `
        SELECT m.name, ci.quantity,
        ROUND((i.import_price * (1 + i.mark_up/100)) * m.unit_per_box) AS price_strip,
        ROUND(i.import_price * (1 + i.mark_up/100)) AS price_box
        FROM cartitems ci
        JOIN medicines m ON ci.medicine_id = m.medicine_id
        JOIN batchsellings bs ON bs.medicine_id = ci.medicine_id
        JOIN inventories i ON i.inventory_id = bs.inventory_id
        WHERE cart_id = ?
    `
	err := config.DB.Raw(query, cart_id).Scan(&items).Error
	return items, err
}
