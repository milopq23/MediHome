package service

import (
	"math"
	"medi-home-be/internal/app/dto"
	"medi-home-be/internal/app/model"
	"medi-home-be/internal/app/repository"
)

type CartService interface {
	GetCartByUser(user_id int64) ([]dto.CartItemResponseDTO, error)
	GetCart(user_id int64) (model.Cart, error)
	// AddMedicineToCart(cartID, medicineID int64) (model.CartItem, error)
	AddMedicineToCart(cartItem dto.CartRequestDTO) (model.CartItem, error)
	UpdateCart(cartItem dto.CartRequestDTO) (model.CartItem, error)
	DeleteCart(cartitem_id int64) error
}

type cartService struct {
	repo repository.CartRepository
}

func NewCartService(repo repository.CartRepository) CartService {
	return &cartService{repo}
}

type UpdateCart struct {
	CartItemID int64 `json:"cartitem_id"`
	// CartID     int64  `json:"cart_id"`
	MedicineID int64  `json:"medicine_id"`
	Quantity   int64  `json:"quantity"`
	SelectType string `json:"select_type"`
}

func (s *cartService) GetCart(user_id int64) (model.Cart, error) {
	return s.repo.GetCartUser(user_id)
}

func (s *cartService) GetCartByUser(user_id int64) ([]dto.CartItemResponseDTO, error) {
	// Lấy cart của user
	cart, err := s.repo.GetCartUser(user_id)
	if err != nil {
		return nil, err
	}

	// Lấy chi tiết các item trong giỏ
	items, err := s.repo.GetCartItemDetail(cart.CartID)
	if err != nil {
		return nil, err
	}

	var itemsDTO []dto.CartItemResponseDTO
	for _, item := range items {
		itemsDTO = append(itemsDTO, dto.CartItemResponseDTO{
			CartItemID: item.CartItemID,
			Name:       item.Name,
			Thumbnail:  item.Thumbnail,
			Quantity:   item.Quantity,
			SelectType: item.SelectType,
			PriceStrip: item.PriceStrip,
			PriceBox:   item.PriceBox,
			Price:      item.Price,
			Total:      item.Total,
		})
	}

	return itemsDTO, nil
}

func (s *cartService) UpdateCart(cartItem dto.CartRequestDTO) (model.CartItem, error) {
	item, err := s.repo.GetItem(cartItem.CartItemID)
	if err != nil {
		return model.CartItem{}, err
	}
	select_type, err := s.SelectType(cartItem.SelectType, item.MedicineID)
	if err != nil {
		return model.CartItem{}, err
	}
	item.Quantity = cartItem.Quantity
	item.SelectType = cartItem.SelectType
	item.Price = select_type.Price
	if err := s.repo.UpdateCartItem(item); err != nil {
		return model.CartItem{}, err
	}
	return item, err
}

func (s *cartService) AddMedicineToCart(cartItem dto.CartRequestDTO) (model.CartItem, error) {
	select_type, err := s.SelectType(cartItem.SelectType, cartItem.MedicineID)
	if err != nil {
		return model.CartItem{}, err
	}

	cartItems, err := s.repo.GetCartItem(cartItem.CartID)
	if err != nil {
		return model.CartItem{}, err
	}

	// 3. Kiểm tra medicine_id có trong cart chưa
	for i := range cartItems {
		if cartItems[i].MedicineID == cartItem.MedicineID {

			// nếu có rồi thì tăng quantity
			cartItems[i].Quantity += 1
			// cập nhật lại cart item trong DB
			err := s.repo.UpdateCartItem(cartItems[i])
			if err != nil {
				return model.CartItem{}, err
			}
			return cartItems[i], nil
		}
	}

	// 4. Nếu chưa có thì tạo mới với quantity = 1
	newItem := model.CartItem{
		CartID:     cartItem.CartID,
		MedicineID: cartItem.MedicineID,
		Quantity:   1,
		SelectType: cartItem.SelectType,
		Price:      select_type.Price,
	}

	return s.repo.AddCartItem(newItem)
}

func (s *cartService) SelectType(select_type string, medicine_id int64) (dto.SelectTypeMedicineDTO, error) {
	//lấy unit_per_box
	medicine, err := s.repo.FindMedicine(medicine_id)
	if err != nil {
		return dto.SelectTypeMedicineDTO{}, err
	}
	//lấy lô của thuốc hiện bán
	batch, err := s.repo.FindBatchSelling(medicine_id)
	if err != nil {
		return dto.SelectTypeMedicineDTO{}, err
	}
	//lấy giá đã tính markup
	inventory, err := s.repo.FindInventory(batch.InventoryID)
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
		MedicineID: medicine_id,
		SelectType: select_type,	
		Price:      math.Round(price*100) / 100,
	}

	return selected, nil
}

func (s *cartService) DeleteCart(cartitem_id int64) error {
	return s.repo.DeleteItem(cartitem_id)
}
