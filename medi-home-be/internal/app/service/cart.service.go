package service

import (
	"medi-home-be/internal/app/model"
	"medi-home-be/internal/app/repository"
)

type CartService interface {
	 GetCartByUser(user_id int64) ([]model.CartItemDetail, error)
}

type cartService struct {
	repo repository.CartRepository
}

func NewCartService(repo repository.CartRepository) CartService {
	return &cartService{repo}
}

func (s *cartService) GetCartByUser(user_id int64) ([]model.CartItemDetail, error) {
    // Lấy cart của user
    cart, err := s.repo.GetCartUser(user_id)
    if err != nil {
        return []model.CartItemDetail{}, err // slice rỗng + error
    }

    // Lấy chi tiết các item trong giỏ
    items, err := s.repo.GetCartItems(cart.CartID)
    if err != nil {
        return []model.CartItemDetail{}, err
    }

    return items, nil
}
