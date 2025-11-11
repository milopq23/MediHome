package service

import (
	"log"
	"math"
	"medi-home-be/internal/app/dto"
	"medi-home-be/internal/app/model"
	"medi-home-be/internal/app/repository"
)

type OrderService interface {
	CheckOut(req dto.OrderRequestDTO) (model.Order, error)
	GetAllOrder() ([]dto.OrderResponse, error)
	GetStatusOrder(status string) ([]dto.OrderResponse, error)
}

type orderService struct {
	repoOrder    repository.OrderRepository
	repoCart     repository.CartRepository
	repoShipping repository.ShippingRepository
	repoVoucher  repository.VoucherRepository
}

func NewOrderService(
	repoOrder repository.OrderRepository,
	repoCart repository.CartRepository,
	repoShipping repository.ShippingRepository,
	repoVoucher repository.VoucherRepository,

) OrderService {
	return &orderService{repoOrder: repoOrder, repoCart: repoCart, repoShipping: repoShipping, repoVoucher: repoVoucher}
}

// thanh toán
func (s *orderService) CheckOut(req dto.OrderRequestDTO) (model.Order, error) {
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

	//giá shipping
	shippingPrice := s.ShippingPrice(req.ShippingID)

	discount, err := s.repoVoucher.ClassifyVoucher(req.VoucherCode, totalPrice)
	if err != nil {
		return model.Order{}, err
	}

	log.Print("discount ", discount)

	//giá sau khi giảm
	final_amount := totalPrice - shippingPrice - discount

	order := model.Order{
		UserID: req.UserID,
		// AddressID: 3,
		VoucherID:     1,
		ShippingID:    req.ShippingID,
		OrderStatus:   "Chờ xác nhận",
		PaymentMethod: req.PaymentMethod,
		PaymentStatus: "Chưa thanh toán",
		TotalAmount:   totalPrice,
		FinalAmount:   final_amount,
	}
	ordered, err := s.repoOrder.CreateOrder(order)
	if err != nil {
		return model.Order{}, err
	}

	for _, item := range cartItems {
		selectType, err := s.SelectType(item.SelectType, item.MedicineID)
		if err != nil {
			return model.Order{}, err
		}

		orderDetail := model.OrderDetail{
			OrderID:     ordered.OrderID,
			InventoryID: selectType.InventoryID,
			MedicineID:  item.MedicineID,
			Quantity:    item.Quantity,
			UnitPrice:   selectType.Price,
			SelectType:  item.SelectType,
		}

		_, err = s.repoOrder.CreateOrderDetail(orderDetail)
		if err != nil {
			return model.Order{}, err
		}
	}
	return ordered, nil
}

func (s *orderService) ShippingPrice(shipping_id int64) float64 {
	shipping, err := s.repoShipping.FindByID(shipping_id)
	if err != nil {
		return 0
	}
	return shipping.Price
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

// lấy tất cả đơn hàng
func (s *orderService) GetAllOrder() ([]dto.OrderResponse, error) {
	order, err := s.repoOrder.GetAllOrder()
	if err != nil {
		return []dto.OrderResponse{}, err
	}
	var response []dto.OrderResponse
	for _, o := range order {
		response = append(response, dto.OrderResponse{
			OrderID:       o.OrderID,
			FullName:      o.FullName,
			Phone:         o.Phone,
			Address:       o.Address,
			VoucherCode:   o.VoucherCode,
			ShippingName:  o.ShippingName,
			OrderStatus:   o.OrderStatus,
			PaymentMethod: o.PaymentMethod,
			PaymentStatus: o.PaymentStatus,
			TotalAmount:   o.TotalAmount,
			FinalAmount:   o.FinalAmount,
		})
	}
	return response, err
}

// lọc đơn hàng theo loại
func (s *orderService) GetStatusOrder(status string) ([]dto.OrderResponse, error) {
	order, err := s.repoOrder.GetStatusTypeOrder(status)
	if err != nil {
		return []dto.OrderResponse{}, err
	}
	var response []dto.OrderResponse
	for _, o := range order {
		response = append(response, dto.OrderResponse{
			OrderID:       o.OrderID,
			FullName:      o.FullName,
			Phone:         o.Phone,
			Address:       o.Address,
			VoucherCode:   o.VoucherCode,
			ShippingName:  o.ShippingName,
			OrderStatus:   o.OrderStatus,
			PaymentMethod: o.PaymentMethod,
			PaymentStatus: o.PaymentStatus,
			TotalAmount:   o.TotalAmount,
			FinalAmount:   o.FinalAmount,
		})
	}
	return response, err
}

// duyệt đơn hàng

// chi tiết đơn hàng
