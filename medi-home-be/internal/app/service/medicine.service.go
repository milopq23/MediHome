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

	//User
	ListMedicine(page, pageSize int) (model.Pagination, error)
	DetailMedicine(id int64) (dto.UserDetailMedicineDTO, error)
	DetailMedicinePrice(id int64) (model.DetailMedicineVM, error)
}

type medicineService struct {
	repo repository.MedicineRepository
}

func NewMedicineService(repo repository.MedicineRepository) MedicineService {
	return &medicineService{repo}
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

func (s *medicineService) Patch(medicine model.Medicine) (model.Medicine, error) {
	return s.repo.Patch(medicine)
}

func (s *medicineService) Delete(id int64) error {
	return s.repo.Delete(id)
}

func (s *medicineService) ListMedicine(page, pageSize int) (model.Pagination, error) {
	return s.repo.ListMedicineUser(page, pageSize)
}

func (s *medicineService) DetailMedicine(id int64) (dto.UserDetailMedicineDTO, error) {
	return s.repo.DetailMedicine(id)
}


func (s *medicineService) DetailMedicinePrice(id int64) (model.DetailMedicineVM, error) {
	return s.repo.DetailMedicineWithPrice(id)
}