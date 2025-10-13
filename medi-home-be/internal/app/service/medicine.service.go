package service

import (
	"errors"
	"log"
	"medi-home-be/internal/app/model"
	"medi-home-be/internal/app/repository"
)

type MedicineService interface {
	GetAll(page, pageSize int) (model.Pagination, error)
	GetByID(id int64) (model.Medicine, error)
	Create(medicine model.Medicine) (model.Medicine, error)
	Patch(medicine model.Medicine) (model.Medicine, error)
	Delete(id int64) error
}

type medicineService struct {
	repo repository.MedicineRepository
}

func NewMedicineService(repo repository.MedicineRepository) MedicineService {
	return &medicineService{repo}
}

func (s *medicineService) GetAll(page, pageSize int) (model.Pagination, error) {
	if s.repo == nil {
		log.Println("s.repo is nil!") // Log lỗi rõ ràng
		return model.Pagination{}, errors.New("internal error: repository not initialized")
	}
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
