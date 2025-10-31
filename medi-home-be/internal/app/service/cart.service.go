package service

import (
	"log"
	"medi-home-be/internal/app/dto"
	"medi-home-be/internal/app/model"
	"medi-home-be/internal/app/repository"
)

type CartService interface {
	GetCartByUser(user_id int64) ([]dto.CartItemDetailDTO, error)
	GetCart(user_id int64) (model.Cart, error)
	AddMedicineToCart(cartID, medicineID int64) (model.CartItem, error)
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

// func (s *cartService) GetCartByUser(user_id int64) (model.CartItem, error) {
// 	// Lấy cart của user
// 	cart, err := s.repo.GetCartUser(user_id)
// 	if err != nil {
// 		return nil, err
// 	}

//     item := model.CartItem{

//         CartID : cart.CartID,
//         MedicineID: dto.MedicneID,
//         Quantity: 1,
//     }
//     return item,nil
// }

func (s *cartService) AddMedicineToCart(cartID, medicineID int64) (model.CartItem, error) {
	item := model.CartItem{
		CartID:     cartID,
		MedicineID: medicineID,
		Quantity:   1,
	}

	return s.repo.AddCartItem(item)
}
