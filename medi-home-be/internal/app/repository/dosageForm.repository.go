package repository

import (
	"medi-home-be/config"
	"medi-home-be/internal/app/model"
)

type DosageForm interface {
	GetAll() ([] model.DosageForm, error)
	FindByID(id int64) (model.DosageForm,error)
	Create(dosageForm model.DosageForm) (model.DosageForm,error)
	Patch(id int64, updates map[string] interface{}) (model.DosageForm, error)
	Delete(id int64) error
}

type dosageFormRepository struct{}

func NewDosageFormRepository() DosageForm {
	return &dosageFormRepository{}
}

func (r *dosageFormRepository)  GetAll() ([]model.DosageForm,error) {
	var dosageForms []model.DosageForm
	err := config.DB.Find(&dosageForms).Error
	return dosageForms, err
}

func (r *dosageFormRepository) FindByID(id int64) (model.DosageForm, error) {
	var dosageForm model.DosageForm
	err := config.DB.Find(&dosageForm, id).Error
	return dosageForm, err
}

func (r *dosageFormRepository) Create(dosageForm model.DosageForm) (model.DosageForm, error) {
	err := config.DB.Create(&dosageForm).Error
	return dosageForm, err
}

func (r *dosageFormRepository) Patch(id int64, updates map[string] interface{}) (model.DosageForm, error) {
	err := config.DB.Model(&model.DosageForm{}).Where("dosageform_id=?").Updates(updates).Error
	if err != nil{
		return model.DosageForm{},err
	}

	var updated model.DosageForm

	err = config.DB.First(&updated,id).Error
	
	return updated, err
}
func (r *dosageFormRepository) Delete(id int64) error {
	err := config.DB.Delete(&model.DosageForm{}, id).Error
	return err
}