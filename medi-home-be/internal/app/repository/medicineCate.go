package repository

import (
	"medi-home-be/config"
	"medi-home-be/internal/app/model"
)

type MedicineCateRepository interface {
	FindAll() ([]model.MedicineCate, error)
	FindByID(id int64) (model.MedicineCate, error)
	Children() ([]model.MedicineCate, error)
	GetChildren(id int64) ([]model.MedicineCate, error)
	Parent() ([]model.MedicineCate, error)
	ListChildren(id int64) (model.MedicineCate, error)
	Create(medicineCate model.MedicineCate) (model.MedicineCate, error)
	Patch(id uint, updates map[string]interface{}) (model.MedicineCate, error)
	Delete(id int64) error

	CreateParentCate(medicineCate model.MedicineCate) (model.MedicineCate, error)
}

type medicineCateRepository struct{}

func NewMedicineCateRepository() MedicineCateRepository {
	return &medicineCateRepository{}
}

func (r *medicineCateRepository) FindAll() ([]model.MedicineCate, error) {
	var medicineCates []model.MedicineCate
	err := config.DB.Preload("Children").Where("parent_id IS NULL").Find(&medicineCates).Error
	return medicineCates, err
}

func (r *medicineCateRepository) FindByID(id int64) (model.MedicineCate, error) {
	var medicineCate model.MedicineCate
	err := config.DB.Find(&medicineCate, id).Error
	return medicineCate, err
}

func (r *medicineCateRepository) GetChildren(id int64) ([]model.MedicineCate, error) {
	var category []model.MedicineCate
	err := config.DB.Where("parent_id = ?", id).Find(&category).Error
	return category, err
}

func (r *medicineCateRepository) ListChildren(id int64) (model.MedicineCate, error) {
	var medicineCate model.MedicineCate
	err := config.DB.Preload("Children").First(&medicineCate, id).Error
	return medicineCate, err
}

func (r *medicineCateRepository) Children() ([]model.MedicineCate, error) {
	var medicineCate []model.MedicineCate
	err := config.DB.Preload("Children").Where("parent_id IS NOT NULL").Find(&medicineCate).Error
	return medicineCate, err
}

func (r *medicineCateRepository) Parent() ([]model.MedicineCate, error) {
	var medicineCate []model.MedicineCate
	err := config.DB.Where("parent_id IS NULL").Find(&medicineCate).Error
	return medicineCate, err
}

func (r *medicineCateRepository) CreateParentCate(medicineCate model.MedicineCate) (model.MedicineCate, error) {
	var parent_id *int64 = nil

	medicineCate = model.MedicineCate{
		Name:     medicineCate.Name,
		Icon:     medicineCate.Icon,
		ParentID: parent_id,
	}

	err := config.DB.Create(&medicineCate).Error
	return medicineCate, err
}

func (r *medicineCateRepository) Create(medicineCate model.MedicineCate) (model.MedicineCate, error) {
	err := config.DB.Create(&medicineCate).Error
	return medicineCate, err
}

func (r *medicineCateRepository) Patch(id uint, updates map[string]interface{}) (model.MedicineCate, error) {
	err := config.DB.Model(&model.MedicineCate{}).Where("medicinecate_id=?", id).Updates(updates).Error
	if err != nil {
		return model.MedicineCate{}, err
	}

	var updated model.MedicineCate
	err = config.DB.First(&updated, id).Error
	return updated, err
}

func (r *medicineCateRepository) Delete(id int64) error {
	err := config.DB.Delete(&model.MedicineCate{}, id).Error
	return err
}
