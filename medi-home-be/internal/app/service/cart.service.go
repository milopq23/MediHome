package service

import (
	"log"
	"math"
	"medi-home-be/internal/app/dto"
	"medi-home-be/internal/app/model"
	"medi-home-be/internal/app/repository"
)

type CartService interface {
	GetCartByUser(user_id int64) ([]dto.CartItemDetailDTO, error)
	GetCart(user_id int64) (model.Cart, error)
	// AddMedicineToCart(cartID, medicineID int64) (model.CartItem, error)
	AddMedicineToCart(cartItem dto.AddCartRequestDTO) (model.CartItem, error)
}

type cartService struct {
	repo repository.CartRepository
}

func NewCartService(repo repository.CartRepository) CartService {
	return &cartService{repo}
}

func (s *cartService) GetCartByUser(user_id int64) ([]dto.CartItemDetailDTO, error) {
	// Lấy cart của user
	cart, err := s.repo.GetCartUser(user_id)
	if err != nil {
		return nil, err
	}

	// Lấy chi tiết các item trong giỏ
	items, err := s.repo.GetCartItems(cart.CartID)
	if err != nil {
		return nil, err
	}
	log.Print(items)

	var itemsDTO []dto.CartItemDetailDTO
	for _, item := range items {
		itemsDTO = append(itemsDTO, dto.CartItemDetailDTO{
			Name:       item.Name,
			Thumbnail:  item.Thumbnail,
			Quantity:   item.Quantity,
			PriceStrip: item.PriceStrip,
			PriceBox:   item.PriceBox,
		})
	}

	return itemsDTO, nil
}

func (s *cartService) GetCart(user_id int64) (model.Cart, error) {
	return s.repo.GetCartUser(user_id)
}

// func (s *cartService) AddMedicineToCart(cartID, medicineID int64) (model.CartItem, error) {
// 	type := s.SelectType()

//		item := model.CartItem{
//			CartID:     cartID,
//			MedicineID: medicineID,
//			Quantity:   1,
//		}
//		return s.repo.AddCartItem(item)
//	}
func (s *cartService) AddMedicineToCart(cartItem dto.AddCartRequestDTO) (model.CartItem, error) {
	select_type, err := s.SelectType(cartItem.SelectType, cartItem.MedicineID)
	if err != nil {
		return model.CartItem{}, err
	}

	item := model.CartItem{
		CartID:     cartItem.CartID,
		MedicineID: cartItem.MedicineID,
		Quantity:   1,
		SelectType: cartItem.SelectType,
		Price:      select_type.Price,
	}
	return s.repo.AddCartItem(item)
}

func (s *cartService) SelectType(select_type string, medicine_id int64) (*dto.SelectTypeMedicineDTO, error) {
	medicine, err := s.repo.FindMedicine(medicine_id)
	if err != nil {
		return nil, err
	}
	batch, err := s.repo.FindBatchSelling(medicine_id)
	if err != nil {
		return nil, err
	}
	// =>inventory và medicine
	inventory, err := s.repo.FindInventory(batch.InventoryID)
	if err != nil {
		return nil, err
	}
	// price

	var price float64
	if select_type == "Box" {
		price = inventory.ImportPrice * (1 + inventory.MarkUp/100)
	} else {
		price = (inventory.ImportPrice * (1 + inventory.MarkUp/100)) / float64(medicine.UnitPerBox)
	}

	// 5. Trả về DTO
	selected := &dto.SelectTypeMedicineDTO{
		MedicineID: medicine_id,
		SelectType: select_type,
		Price:      math.Round(price*100) / 100, // làm tròn 2 chữ số
	}

	return selected, nil
}
