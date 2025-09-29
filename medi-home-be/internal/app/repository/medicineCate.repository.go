package repository

import (
	"medi-home-be/config"
	"medi-home-be/internal/app/model"
)

type MedicineCateRepository interface {
	FindAll() ([]model.MedicineCate, error)
	ListChildren(id int64) (model.MedicineCate, error)
	Create(medicineCate model.MedicineCate) (model.MedicineCate, error)
	Patch(medicineCate model.MedicineCate) (model.MedicineCate, error)
	Delete(id int64) error	
}

type medicineCateRepository struct{}

func NewMedicineCateRepository() MedicineCateRepository {
	return &medicineCateRepository{}
}

func (r *medicineCateRepository)  FindAll() ([]model.MedicineCate, error) {
	var medicineCates []model.MedicineCate
	err := config.DB.Preload("Children").Where("parent_id IS NULL").Find(&medicineCates).Error
	return medicineCates, err
}

func (r *medicineCateRepository) ListChildren(id int64) (model.MedicineCate, error) {
	var medicineCate model.MedicineCate
	err := config.DB.Preload("Children").First(&medicineCate, id).Error
	return medicineCate, err
}

func (r *medicineCateRepository) Create(medicineCate model.MedicineCate) (model.MedicineCate, error) {
	err := config.DB.Create(&medicineCate).Error
	return medicineCate, err
}

func (r *medicineCateRepository) Patch(medicineCate model.MedicineCate) (model.MedicineCate, error) {
	err := config.DB.Save(&medicineCate).Error
	return medicineCate, err
}

func (r *medicineCateRepository) Delete(id int64) error {
	err := config.DB.Delete(&model.MedicineCate{}, id).Error
	return err
}