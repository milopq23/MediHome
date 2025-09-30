package repository

import (
	"medi-home-be/config"
	"medi-home-be/internal/app/model"
)

type MedicineRepository interface {
	GetAll() ([]model.Medicine,error)
	FindByID(id int64) (model.Medicine, error)
	Create(medicine model.Medicine) (model.Medicine, error)
	Patch(medicine model.Medicine) (model.Medicine, error)
	Delete(id int64) error
}

type medicineRepository struct{}

func NewMedicineRepository() MedicineRepository {
	return &medicineRepository{}
}

func (r *medicineRepository)  GetAll() ([]model.Medicine,error) {
	var medicines []model.Medicine
	err := config.DB.Find(&medicines).Error
	return medicines, err
}

func (r *medicineRepository) FindByID(id int64) (model.Medicine, error) {
	var medicine model.Medicine
	err := config.DB.Find(&medicine, id).Error
	return medicine, err
}

func (r *medicineRepository) Create(medicine model.Medicine) (model.Medicine, error) {
	err := config.DB.Create(&medicine).Error
	return medicine, err
}


func (r *medicineRepository) Patch(medicine model.Medicine) (model.Medicine, error) {
	err := config.DB.Save(&medicine).Error
	return medicine, err
}

func (r *medicineRepository) Delete(id int64) error {
	err := config.DB.Delete(&model.Medicine{}, id).Error
	return err
}