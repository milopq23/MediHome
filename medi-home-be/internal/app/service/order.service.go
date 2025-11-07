package service

import (
	"math"
	"medi-home-be/internal/app/dto"
	"medi-home-be/internal/app/model"
	"medi-home-be/internal/app/repository"
)

type OrderService interface {
	CheckOut(req dto.OrderRequestDTO) (model.Order, error)
}

type orderService struct {
	repoOrder repository.OrderRepository
	repoCart  repository.CartRepository
}

func NewOrderService(
	repoOrder repository.OrderRepository,
	repoCart repository.CartRepository,

) OrderService {
	return &orderService{repoOrder: repoOrder, repoCart: repoCart}
}

func (s *orderService) CheckOut(req dto.OrderRequestDTO) (model.Order, error) {
	// totalAmount, err := s.TotalPriceInCart(req.UserID)
	// if err != nil {
	// 	return model.Order{}, err
	// }

	// cartItems, err := s.repoCart.GetCartItem(user.CartID)
	// if err != nil {
	// 	return 0, err
	// }
	// if err != nil {
	// 	return 0, err
	// }
	user, err := s.repoCart.GetCartUser(req.UserID)
	if err != nil {
		return model.Order{}, err
	}

	cartItems, err := s.repoCart.GetCartItem(user.CartID)
	if err != nil {
		return model.Order{}, err
	}

	var totalPrice float64
	for _, item := range cartItems {
		selectType, err := s.SelectType(item.SelectType, item.MedicineID)
		if err != nil {
			return model.Order{}, err
		}

		totalPrice += selectType.Price * float64(item.Quantity)
	}

	order := model.Order{
		UserID: req.UserID,
		// AddressID: 3,
		VoucherID:     1,
		ShippingID:    req.ShippingID,
		OrderStatus:   "Chờ xác nhận",
		PaymentMethod: req.PaymentMethod,
		PaymentStatus: "Chưa thanh toán",
		TotalAmount:   int64(totalPrice),
		FinalAmount:   int64(totalPrice),
	}
	ordered, err := s.repoOrder.CreateOrder(order)
	if err != nil {
		return model.Order{}, err
	}

	// for _, item := range cartItems {
	// 	orderDetail := model.OrderDetail{
	// 		OrderID:     ordered.OrderID,
	// 		InventoryID: selectType.InventoryID,
	// 		MedicineID:  item.MedicineID,
	// 		Quantity:    item.Quantity,
	// 		UnitPrice:   item.Price,
	// 		SelectType:  item.SelectType,
	// 	}
	// 	orderdetail, err := s.repoOrder.CreateOrderDetail(orderDetail)
	// 	if err != nil {
	// 		return model.Order{}, err
	// 	}
	// }

	for _, item := range cartItems {
		selectType, err := s.SelectType(item.SelectType, item.MedicineID)
		if err != nil {
			return model.Order{}, err
		}

		orderDetail := model.OrderDetail{
			OrderID:     ordered.OrderID,
			InventoryID: selectType.InventoryID, // không dùng biến ngoài vòng lặp
			MedicineID:  item.MedicineID,
			Quantity:    item.Quantity,
			UnitPrice:   selectType.Price, // thống nhất giá
			SelectType:  item.SelectType,
		}

		_, err = s.repoOrder.CreateOrderDetail(orderDetail)
		if err != nil {
			return model.Order{}, err
		}
	}
	return ordered, nil
}

func (s *orderService) TotalPriceInCart(user_id int64) (float64, error) {
	user, err := s.repoCart.GetCartUser(user_id)
	if err != nil {
		return 0, err
	}

	cartItems, err := s.repoCart.GetCartItem(user.CartID)
	if err != nil {
		return 0, err
	}

	var totalPrice float64
	for _, item := range cartItems {
		selectType, err := s.SelectType(item.SelectType, item.MedicineID)
		if err != nil {
			return 0, err
		}
		totalPrice += selectType.Price * float64(item.Quantity)
	}

	return totalPrice, err
}

func (s *orderService) SelectType(select_type string, medicine_id int64) (dto.SelectTypeMedicineDTO, error) {
	//lấy unit_per_box
	medicine, err := s.repoCart.FindMedicine(medicine_id)
	if err != nil {
		return dto.SelectTypeMedicineDTO{}, err
	}
	//lấy lô của thuốc hiện bán
	batch, err := s.repoCart.FindBatchSelling(medicine_id)
	if err != nil {
		return dto.SelectTypeMedicineDTO{}, err
	}
	//lấy giá đã tính markup
	inventory, err := s.repoCart.FindInventory(batch.InventoryID)
	if err != nil {
		return dto.SelectTypeMedicineDTO{}, err
	}
	//phân loại giá theo box hoặc strip
	var price float64
	if select_type == "Box" {
		price = inventory.ImportPrice * (1 + inventory.MarkUp/100)
	} else {
		price = (inventory.ImportPrice * (1 + inventory.MarkUp/100)) / float64(medicine.UnitPerBox)
	}

	// 5. Trả về DTO
	selected := dto.SelectTypeMedicineDTO{
		MedicineID:  medicine_id,
		SelectType:  select_type,
		InventoryID: batch.InventoryID,
		Price:       math.Round(price*100) / 100,
	}

	return selected, nil
}
