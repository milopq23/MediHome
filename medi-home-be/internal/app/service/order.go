package service

import (
	"math"
	"medi-home-be/internal/app/dto"
	"medi-home-be/internal/app/model"
	"medi-home-be/internal/app/repository"
)

type OrderService interface {
	CheckOut(req dto.OrderRequestDTO) (model.Order, error)
	GetAllOrder() ([]dto.AllOrderResponse, error)
	GetStatusOrder(status string) ([]dto.AllOrderResponse, error)
	GetDetailOrder(order_id int64) (dto.OrderResponse, error)
	GetAllUserOrder(user_id int64) ([]dto.AllOrderResponse, error)
	GetStatusUserOrder(user_id int64, status string) ([]dto.AllOrderResponse, error)
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

		err = s.repoInventory.DecreaseQuantity(selectType.InventoryID, item.Quantity)
		if err != nil {
			return model.Order{}, err
		}

		_, err = s.LogTransactionOut(selectType.InventoryID, item.Quantity)
		if err != nil {
			return model.Order{}, err
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
func (s *orderService) GetAllOrder() ([]dto.AllOrderResponse, error) {
	order, err := s.repoOrder.GetViewAllOrder()
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

// lọc đơn hàng theo loại
func (s *orderService) GetStatusOrder(status string) ([]dto.AllOrderResponse, error) {
	order, err := s.repoOrder.GetStatusTypeOrder(status)
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

func (s *orderService) GetStatusUserOrder(user_id int64, status string) ([]dto.AllOrderResponse, error) {
	order, err := s.repoOrder.GetViewOrderStatusByUser(user_id, status)
	if err != nil {
		return []dto.AllOrderResponse{}, err
	}
	// log.Print(order)
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
	// log.Print("res",response)
	return response, err
}

// duyệt đơn hàng
// func (s *orderService) ApproveStatus(order_id int64) (model.Order,error){
// 	order,err := s.repoOrder.UpdateStatus
// }

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
