package service

import (
	"medi-home-be/internal/app/model"
	"medi-home-be/internal/app/repository"
)

type MedicineCateService interface {
	GetAll() ([]model.MedicineCate, error)
	ListChildren(id int64) (model.MedicineCate, error)
	Create(medicineCate model.MedicineCate) (model.MedicineCate, error)
	Patch(medicineCate model.MedicineCate) (model.MedicineCate, error)
	Delete(id int64) error
}

type medicineCateService struct {
	repo repository.MedicineCateRepository
}

func NewMedicineCateService(repo repository.MedicineCateRepository) MedicineCateService {
	return &medicineCateService{}
}

func (s *medicineCateService) GetAll() ([]model.MedicineCate, error) {
	return s.repo.FindAll()
}

func (s *medicineCateService) ListChildren(id int64) (model.MedicineCate, error) {
	return s.repo.ListChildren(id)
}

func (s *medicineCateService) Create(medicineCate model.MedicineCate) (model.MedicineCate, error) {
	return s.repo.Create(medicineCate)
}

func (s *medicineCateService) Patch(medicineCate model.MedicineCate) (model.MedicineCate, error) {
	return s.repo.Patch(medicineCate)
}

func (s *medicineCateService) Delete(id int64) error {
	return s.repo.Delete(id)
}