package service

import (
	"medi-home-be/internal/app/dto"
	"medi-home-be/internal/app/model"
	"medi-home-be/internal/app/repository"
)

type MedicineService interface {
	//Admin
	GetAll(page, pageSize int) (model.Pagination, error)
	GetByID(id int64) (model.Medicine, error)
	Create(medicine model.Medicine) (model.Medicine, error)
	Patch(medicine model.Medicine) (model.Medicine, error)
	Delete(id int64) error

	SaveImages(medicineID int64, urls []string) error

	//User
	ListMedicine(page, pageSize int) (model.Pagination, error)
	// DetailMedicine(id int64) (dto.UserDetailMedicineDTO, error)
	DetailMedicineUser(id int64) (dto.DetailMedicine, error)
}

type medicineService struct {
	repo     repository.MedicineRepository
	cartRepo repository.CartRepository
}

func NewMedicineService(repo repository.MedicineRepository, cartRepo repository.CartRepository) MedicineService {
	return &medicineService{repo: repo, cartRepo: cartRepo}
}

func (s *medicineService) GetAll(page, pageSize int) (model.Pagination, error) {
	return s.repo.FindAll(page, pageSize)
}

func (s *medicineService) GetByID(id int64) (model.Medicine, error) {
	return s.repo.FindByID(id)
}

func (s *medicineService) Create(medicine model.Medicine) (model.Medicine, error) {
	return s.repo.Create(medicine)
}

func (s *medicineService) SaveImages(medicineID int64, urls []string) error {

	if len(urls) == 0 {
		return nil // không save ảnh khi không có gì
	}

	return s.repo.AddImages(medicineID, urls)
}

func (s *medicineService) Patch(medicine model.Medicine) (model.Medicine, error) {
	return s.repo.Patch(medicine)
}

func (s *medicineService) Delete(id int64) error {
	return s.repo.Delete(id)
}

func (s *medicineService) ListMedicine(page, pageSize int) (model.Pagination, error) {
	return s.repo.ListMedicineUser(page, pageSize)
}

// func (s *medicineService) DetailMedicine(id int64) (dto.UserDetailMedicineDTO, error) {
// 	return s.repo.DetailMedicine(id)
// }

func (s *medicineService) DetailMedicineUser(id int64) (dto.DetailMedicine, error) {
	medicine, err := s.repo.DetailMedicineUser(id)
	if err != nil {
		return dto.DetailMedicine{}, err
	}
	return dto.DetailMedicine{
		MedicineID:       medicine.MedicineID,
		Code:             medicine.Code,
		Name:             medicine.Name,
		Thumbnail:        medicine.Thumbnail,
		Image:            medicine.Image,
		UnitPerStrip:     medicine.UnitPerStrip,
		UnitPerBox:       medicine.UnitPerBox,
		MedCategoryName:  medicine.MedCategoryName,
		DosageFormName:   medicine.DosageFormName,
		PriceForStrip:    medicine.PriceForStrip,
		PriceForBox:      medicine.PriceForBox,
		Prescription:     medicine.Prescription,
		Usage:            medicine.Usage,
		Package:          medicine.Package,
		Indication:       medicine.Indication,
		Adverse:          medicine.Adverse,
		Contraindication: medicine.Contraindication,
		Precaution:       medicine.Precaution,
		Ability:          medicine.Ability,
		Pregnancy:        medicine.Pregnancy,
		DrugInteraction:  medicine.DrugInteraction,
		Storage:          medicine.Storage,
		Manufacturer:     medicine.Manufacturer,
		Note:             medicine.Note,
	}, nil
}

// func (s *medicineService) AddProductToCart(medicine_id int64) (model.CartItemDetail, error) {
// 	medicine, err := s.repo.FindByID(medicine_id)
// 	if err != nil {
// 		return model.CartItemDetail{}, fmt.Errorf("cannot find medicine with ID %d: %w", medicine_id, err)
// 	}

// 	// Kiểm tra số lượng tồn kho
// 	// if medicine.Stock <= 0 {
// 	// 	return model.CartItemDetail{}, fmt.Errorf("medicine '%s' is out of stock", medicine.Name)
// 	// }

// 	// Tạo thông tin sản phẩm trong giỏ
// 	item := model.CartItemDetail{
// 		MedicineID: medicine_id,
// 		Name:       medicine.Name,
// 		Quantity:   1,
// 		// PriceStrip: medicine.Price,
// 		// PriceBox:   medicine.Price * 1,
// 	}

// 	// // Lưu vào giỏ hàng (nếu có repository riêng cho cart)
// 	// if err := s.cartRepo.AddItem(item); err != nil {
// 	// 	return model.CartItemDetail{}, fmt.Errorf("failed to add item to cart: %w", err)
// 	// }

// 	return item, nil
// }
