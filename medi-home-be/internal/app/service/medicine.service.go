package service

import (
	"medi-home-be/internal/app/model"
	"medi-home-be/internal/app/repository"
)

type MedicineService interface {
	GetAll() ([]model.Medicine, error)
	GetByID(id int64) (model.Medicine, error)
	Create(medicine model.Medicine) (model.Medicine, error)
	Patch(medicine model.Medicine) (model.Medicine, error)
	Delete(id int64) error
}

type medicineService struct {
	repo repository.MedicineRepository
}

func NewMedicineService(repo repository.MedicineRepository) MedicineService {
	return &medicineService{}
}

func (s *medicineService) GetAll() ([]model.Medicine, error) {
	return s.repo.GetAll()
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