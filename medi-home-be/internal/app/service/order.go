package service

import (
	"fmt"
	"log"
	"math"
	"medi-home-be/config"
	"medi-home-be/internal/app/dto"
	"medi-home-be/internal/app/model"
	"medi-home-be/internal/app/repository"
)

type OrderService interface {
	//VIEW
	GetAllOrder() ([]dto.AllOrderResponse, error)
	GetDetailOrder(order_id int64) (dto.OrderResponse, error)
	GetStatusOrder(status string) ([]dto.AllOrderResponse, error)
	GetAllUserOrder(user_id int64) ([]dto.AllOrderResponse, error)
	GetStatusUserOrder(user_id int64, status string) ([]dto.AllOrderResponse, error)

	//ACTION
	ApproveStatus(order_id int64, newStatus string) (model.Order, error)
	CheckOut(req dto.OrderRequestDTO) (model.Order, error)
}

type orderService struct {
	repoOrder     repository.OrderRepository
	repoCart      repository.CartRepository
	repoShipping  repository.ShippingRepository
	repoVoucher   repository.VoucherRepository
	repoInventory repository.InventoryRepository
	repoLog       repository.StockTransactionRepository
}

func NewOrderService(
	repoOrder repository.OrderRepository,
	repoCart repository.CartRepository,
	repoShipping repository.ShippingRepository,
	repoVoucher repository.VoucherRepository,
	repoInventory repository.InventoryRepository,
	repoLog repository.StockTransactionRepository,

) OrderService {
	return &orderService{
		repoOrder:     repoOrder,
		repoCart:      repoCart,
		repoShipping:  repoShipping,
		repoVoucher:   repoVoucher,
		repoInventory: repoInventory,
		repoLog:       repoLog,
	}
}

// lấy tất cả đơn hàng
func (s *orderService) GetAllOrder() ([]dto.AllOrderResponse, error) {
	order, err := s.repoOrder.GetViewAllOrder()
	if err != nil {
		log.Printf("Lỗi không view được  đơn hàng: %v", err)
		return []dto.AllOrderResponse{}, err
	}
	var response []dto.AllOrderResponse
	for _, o := range order {
		response = append(response, dto.AllOrderResponse{
			OrderID:       o.OrderID,
			Date:          o.Date,
			FullName:      o.FullName,
			OrderItem:     o.OrderItem,
			OrderStatus:   o.OrderStatus,
			PaymentMethod: o.PaymentMethod,
			PaymentStatus: o.PaymentStatus,
			FinalAmount:   o.FinalAmount,
		})
	}
	return response, err
}

// chi tiết đơn hàng
func (s *orderService) GetDetailOrder(order_id int64) (dto.OrderResponse, error) {
	detail_items, err := s.repoOrder.GetViewOrderDetail(order_id)
	if err != nil {
		return dto.OrderResponse{}, err
	}
	order_info, err := s.repoOrder.GetViewOrder(order_id)
	if err != nil {
		return dto.OrderResponse{}, err
	}
	var item_response []dto.OrderDetailResponse
	for _, res := range detail_items {
		item_response = append(item_response, dto.OrderDetailResponse{
			MedicineName: res.MedicineName,
			UnitPrice:    res.UnitPrice,
			Type:         res.Type,
			Quantity:     res.Quantity,
		})
	}

	order_response := dto.OrderResponse{
		OrderID:       order_info.OrderID,
		FullName:      order_info.FullName,
		Phone:         order_info.Phone,
		Address:       order_info.Address,
		VoucherCode:   order_info.VoucherCode,
		ShippingName:  order_info.ShippingName,
		OrderStatus:   order_info.OrderStatus,
		PaymentMethod: order_info.PaymentMethod,
		PaymentStatus: order_info.PaymentStatus,
		TotalAmount:   order_info.TotalAmount,
		FinalAmount:   order_info.FinalAmount,
		OrderDetail:   item_response,
	}
	return order_response, err
}

// lọc đơn hàng theo loại
func (s *orderService) GetStatusOrder(status string) ([]dto.AllOrderResponse, error) {
	order, err := s.repoOrder.GetViewOrderStatus(status)
	if err != nil {
		log.Printf("Lỗi không view được đơn hàng theo loại: %v", err)
		return []dto.AllOrderResponse{}, err
	}
	var response []dto.AllOrderResponse
	for _, o := range order {
		response = append(response, dto.AllOrderResponse{
			OrderID:       o.OrderID,
			Date:          o.Date,
			FullName:      o.FullName,
			OrderItem:     o.OrderItem,
			OrderStatus:   o.OrderStatus,
			PaymentMethod: o.PaymentMethod,
			PaymentStatus: o.PaymentStatus,
			FinalAmount:   o.FinalAmount,
		})
	}
	return response, err
}

// lấy đơn hàng của user
func (s *orderService) GetAllUserOrder(user_id int64) ([]dto.AllOrderResponse, error) {
	order, err := s.repoOrder.GetViewAllOrderUser(user_id)
	if err != nil {
		return []dto.AllOrderResponse{}, err
	}
	var response []dto.AllOrderResponse
	for _, o := range order {
		response = append(response, dto.AllOrderResponse{
			OrderID:       o.OrderID,
			Date:          o.Date,
			FullName:      o.FullName,
			OrderItem:     o.OrderItem,
			OrderStatus:   o.OrderStatus,
			PaymentMethod: o.PaymentMethod,
			PaymentStatus: o.PaymentStatus,
			FinalAmount:   o.FinalAmount,
		})
	}
	return response, err
}

// lấy đơn hàng của user theo loại
func (s *orderService) GetStatusUserOrder(user_id int64, status string) ([]dto.AllOrderResponse, error) {
	order, err := s.repoOrder.GetViewOrderStatusByUser(user_id, status)
	if err != nil {
		return []dto.AllOrderResponse{}, err
	}
	var response []dto.AllOrderResponse
	for _, o := range order {
		response = append(response, dto.AllOrderResponse{
			OrderID:       o.OrderID,
			Date:          o.Date,
			FullName:      o.FullName,
			OrderItem:     o.OrderItem,
			OrderStatus:   o.OrderStatus,
			PaymentMethod: o.PaymentMethod,
			PaymentStatus: o.PaymentStatus,
			FinalAmount:   o.FinalAmount,
		})
	}
	return response, err
}

func (s *orderService) ShippingPrice(shipping_id int64) float64 {
	shipping, err := s.repoShipping.FindByID(shipping_id)
	if err != nil {
		log.Printf("Lỗi không lấy giá shipping: %v", err)
		return 0
	}
	return shipping.Price
}

var flow_order = map[string][]string{
	"Chờ xác nhận": {"Đã xác nhận", "Đã hủy"},
	"Đã xác nhận":  {"Đang giao", "Đã hủy"},
	"Đang giao":    {"Hoàn thành"},
}

func FlowOrder(current, next string) bool {
	for _, s := range flow_order[current] {
		if s == next {
			return true
		}
	}
	return false
}

// duyệt đơn hàng
func (s *orderService) ApproveStatus(order_id int64, newStatus string) (model.Order, error) {
	var order model.Order
	order, err := s.repoOrder.GetOrderID(order_id)
	if err != nil {
		return model.Order{}, err
	}

	if !FlowOrder(order.OrderStatus, newStatus) {
		return model.Order{}, fmt.Errorf("không thể chuyển từ '%s' sang '%s'", order.OrderStatus, newStatus)
	}

	updatedOrder, err := s.repoOrder.UpdateStatus(order_id, newStatus)
	log.Print("order", order)
	if err != nil {
		return model.Order{}, err
	}
	return updatedOrder, err
}

// thanh toán
func (s *orderService) CheckOut(req dto.OrderRequestDTO) (model.Order, error) {
	tx := config.DB.Begin()

	if tx.Error != nil {
		return model.Order{}, tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	user, err := s.repoCart.GetCartUser(req.UserID)
	if err != nil {
		tx.Rollback()
		return model.Order{}, err
	}

	cartItems, err := s.repoCart.GetCartItem(user.CartID)
	if err != nil {
		tx.Rollback()
		return model.Order{}, err
	}

	var totalPrice float64
	for _, item := range cartItems {
		selectType, err := s.SelectType(item.SelectType, item.MedicineID)
		if err != nil {
			tx.Rollback()
			return model.Order{}, err
		}
		totalPrice += selectType.Price * float64(item.Quantity)
	}

	// Xử lý voucher
	var discount float64
	var voucher model.Voucher
	if req.VoucherCode != "" {
		discount, err = s.repoVoucher.ClassifyVoucher(req.VoucherCode, totalPrice)
		if err != nil {
			tx.Rollback()
			return model.Order{}, err
		}
		voucher, err = s.repoVoucher.GetVoucherByCode(req.VoucherCode)
		if err != nil {
			tx.Rollback()
			return model.Order{}, err
		}
	}

	finalAmount := totalPrice + s.ShippingPrice(req.ShippingID) - discount

	order := model.Order{
		UserID:        req.UserID,
		AddressID:     3,
		VoucherID:     voucher.VoucherID,
		ShippingID:    req.ShippingID,
		OrderStatus:   "Chờ xác nhận",
		PaymentMethod: req.PaymentMethod,
		PaymentStatus: "Chưa thanh toán",
		TotalAmount:   totalPrice,
		FinalAmount:   finalAmount,
	}

	// gọi repo với transaction
	ordered, err := s.repoOrder.CreateOrder(tx, order)
	if err != nil {
		tx.Rollback()
		return model.Order{}, err
	}

	for _, item := range cartItems {
		selectType, err := s.SelectType(item.SelectType, item.MedicineID)
		if err != nil {
			tx.Rollback()
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

		if err := s.repoInventory.DecreaseQuantity(tx, selectType.InventoryID, item.Quantity); err != nil {
			tx.Rollback()
			return model.Order{}, err
		}

		if _, err := s.repoOrder.CreateOrderDetail(tx, orderDetail); err != nil {
			tx.Rollback()
			return model.Order{}, err
		}
	}

	// commit toàn bộ
	if err := tx.Commit().Error; err != nil {
		return model.Order{}, err
	}

	// user, err := s.repoCart.GetCartUser(req.UserID)
	// if err != nil {
	// 	return model.Order{}, err
	// }

	// cartItems, err := s.repoCart.GetCartItem(user.CartID)
	// if err != nil {
	// 	return model.Order{}, err
	// }

	// var totalPrice float64
	// for _, item := range cartItems {
	// 	selectType, err := s.SelectType(item.SelectType, item.MedicineID)
	// 	if err != nil {
	// 		return model.Order{}, err
	// 	}
	// 	totalPrice += selectType.Price * float64(item.Quantity)
	// }

	// //giảm giá
	// var discount float64

	// var voucher model.Voucher
	// if req.VoucherCode != "" {
	// 	discount, err = s.repoVoucher.ClassifyVoucher(req.VoucherCode, totalPrice)
	// 	if err != nil {
	// 		return model.Order{}, err
	// 	}

	// 	voucher, err = s.repoVoucher.GetVoucherByCode(req.VoucherCode)
	// 	if err != nil {
	// 		return model.Order{}, err
	// 	}
	// }

	// //giá sau khi giảm
	// final_amount := totalPrice + s.ShippingPrice(req.ShippingID) - discount

	// order := model.Order{
	// 	UserID:        req.UserID,
	// 	AddressID:     3,
	// 	VoucherID:     voucher.VoucherID,
	// 	ShippingID:    req.ShippingID,
	// 	OrderStatus:   "Chờ xác nhận",
	// 	PaymentMethod: req.PaymentMethod,
	// 	PaymentStatus: "Chưa thanh toán",
	// 	TotalAmount:   totalPrice,
	// 	FinalAmount:   final_amount,
	// }
	// ordered, err := s.repoOrder.CreateOrder(order)
	// if err != nil {
	// 	log.Printf("Lỗi không tạo order: %v", err)
	// 	return model.Order{}, err
	// }

	// for _, item := range cartItems {
	// 	selectType, err := s.SelectType(item.SelectType, item.MedicineID)
	// 	if err != nil {
	// 		log.Printf("Lỗi không lấy được giá tiền: %v", err)
	// 		return model.Order{}, err
	// 	}

	// 	orderDetail := model.OrderDetail{
	// 		OrderID:     ordered.OrderID,
	// 		InventoryID: selectType.InventoryID,
	// 		MedicineID:  item.MedicineID,
	// 		Quantity:    item.Quantity,
	// 		UnitPrice:   selectType.Price,
	// 		SelectType:  item.SelectType,
	// 	}

	// 	err = s.repoInventory.DecreaseQuantity(selectType.InventoryID, item.Quantity)
	// 	if err != nil {
	// 		log.Printf("Lỗi giảm số lượng inventory: %v", err)
	// 		return model.Order{}, err
	// 	}

	// 	_, err = s.LogTransactionOut(selectType.InventoryID, item.Quantity)
	// 	if err != nil {
	// 		log.Printf("Lỗi ghi log transaction: %v", err)
	// 		return model.Order{}, err
	// 	}

	// 	_, err = s.repoOrder.CreateOrderDetail(orderDetail)
	// 	if err != nil {
	// 		log.Printf("Lỗi tạo order detail: %v", err)
	// 		return model.Order{}, err
	// 	}
	// }

	return ordered, nil
}

func (s *orderService) SelectType(select_type string, medicine_id int64) (dto.SelectTypeMedicineDTO, error) {
	//lấy unit_per_box
	medicine, err := s.repoCart.FindMedicine(medicine_id)
	if err != nil {
		log.Printf("Lỗi tìm thuốc: %v", err)
		return dto.SelectTypeMedicineDTO{}, err
	}
	//lấy lô của thuốc hiện bán
	batch, err := s.repoCart.FindBatchSelling(medicine_id)
	if err != nil {
		log.Printf("Lỗi tìm lô hiện bán: %v", err)
		return dto.SelectTypeMedicineDTO{}, err
	}
	//lấy giá đã tính markup
	inventory, err := s.repoCart.FindInventory(batch.InventoryID)
	if err != nil {
		log.Printf("Lỗi tìm hàng tồn: %v", err)
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

func (s *orderService) LogTransactionOut(inventory_id, quantity int64) (model.StockTransaction, error) {
	log := model.StockTransaction{
		InventoryID:     inventory_id,
		TransactionType: "OUT",
		Quantity:        quantity,
		Note:            "Bán hàng",
	}
	_, err := s.repoLog.CreateLogTransaction(log)
	if err != nil {
		return model.StockTransaction{}, err
	}
	return log, err
}
