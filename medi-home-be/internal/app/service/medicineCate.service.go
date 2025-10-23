package service

import (
	// "fmt"
	"fmt"
	"medi-home-be/internal/app/model"
	"medi-home-be/internal/app/repository"
)

type MedicineCateService interface {
	GetAll() ([]model.MedicineCate, error)
	ListChildren(id int64) (model.MedicineCate, error)
	Create(medicineCate model.MedicineCate) (model.MedicineCate, error)
	CreateParentCate(medicineCate model.MedicineCate) (model.MedicineCate, error)
	Patch(id uint, updates map[string]interface{}) (model.MedicineCate, error)
	Delete(id int64) error
}

type medicineCateService struct {
	repo repository.MedicineCateRepository
}

func NewMedicineCateService(repo repository.MedicineCateRepository) MedicineCateService {
	return &medicineCateService{repo}
}

func (s *medicineCateService) GetAll() ([]model.MedicineCate, error) {
	return s.repo.FindAll()
}

func (s *medicineCateService) ListChildren(id int64) (model.MedicineCate, error) {
	return s.repo.ListChildren(id)
}

func (s *medicineCateService) CreateParentCate(medicineCate model.MedicineCate) (model.MedicineCate, error) {
	return s.repo.CreateParentCate(medicineCate)
}

func (s *medicineCateService) Create(medicineCate model.MedicineCate) (model.MedicineCate, error) {
	return s.repo.Create(medicineCate)
}

func (s *medicineCateService) Patch(id uint, data map[string]interface{}) (model.MedicineCate, error) {
	medicineCate, err := s.repo.FindByID(int64(id))
	if err != nil {
		return model.MedicineCate{}, err
	}

	allowedFields := map[string]bool{
		"name":      true,
		"icon":      true,
		"parent_id": true,
	}

	// Lọc ra những field hợp lệ từ client
	updates := make(map[string]interface{})
	for k, v := range data {
		if allowedFields[k] && v != nil {
			updates[k] = v
		}
	}

	if len(updates) == 0 {
		return model.MedicineCate{}, fmt.Errorf("no valid fields to update")
	}

	return s.repo.Patch(uint(medicineCate.MedicineCateID), updates)
}

func (s *medicineCateService) Delete(id int64) error {
	return s.repo.Delete(id)
}
