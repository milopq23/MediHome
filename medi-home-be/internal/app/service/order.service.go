package service

import (
	"medi-home-be/internal/app/dto"
	"medi-home-be/internal/app/repository"
)

type OrderService interface {
}

type orderService struct {
	orderRepo    repository.OrderRepository
	userRepo     repository.UserRepository
	cartRepo     repository.CartRepository
	voucherRepo  repository.VoucherRepository
	shippingRepo repository.ShippingRepository
}

func NewOrderService(
	orderRepo repository.OrderRepository,
	userRepo repository.UserRepository,
	cartRepo repository.CartRepository,
	voucherRepo repository.VoucherRepository,
	shippingRepo repository.ShippingRepository,
) OrderService {
	return &orderService{
		orderRepo:    orderRepo,
		userRepo:     userRepo,
		cartRepo:     cartRepo,
		voucherRepo:  voucherRepo,
		shippingRepo: shippingRepo,
	}
}

//get user
//get voucher_id
//get shipping_id
//get cart push cart vào order

// func (s *orderService) AddOrder(order dto.OrderDTO) (dto.OrderDTO, error) {
// 	cart, err := s.cartRepo.GetCartUser(order.OrderID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	voucher, err := s.voucherRepo.GetVoucherByCode(order.VoucherCode)
// 	if err != nil {
// 		return nil, err
// 	}
// 	shipping, err := s.shippingRepo.FindByID(order.ShippingID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	orderItem := dto.OrderDTO{
// 		UserID: cart.UserID,
// 		VoucherCode: voucher.Code,
// 	}

// }

// nhận iduser từ cookie lấy được cart
// func (s *orderService) CheckOut(order dto.OrderDTO) (dto.OrderDTO, error) {

// }
func (s *orderService) GetCart(user_id int64) ([]dto.OrderDetailItemDTO, error) {
	cart, err := s.cartRepo.GetCartUser(user_id)
	if err != nil {
		return nil, err
	}
	items, err := s.cartRepo.GetCartItems(cart.CartID)
	if err != nil {
		return nil, err
	}
	var orderDetailDTO []dto.OrderDetailItemDTO
	for _, item := range items {
		orderDetailDTO = append(orderDetailDTO, dto.OrderDetailItemDTO{
			Name:     item.Name,
			Quantity: int64(item.Quantity),
			Price:    item.PriceBox,
		})
	}
	return orderDetailDTO, nil

}
