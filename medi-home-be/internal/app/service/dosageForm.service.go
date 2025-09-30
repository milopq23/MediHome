package service

import (
	"medi-home-be/internal/app/model"
	"medi-home-be/internal/app/repository"
)

type DosageFormService interface{
	GetByID(id int64) (model.DosageForm, error)
	Create(dosage model.DosageForm) (model.DosageForm, error)
	Patch(id int64, data map[string] interface{}) (model.DosageForm,error)
	Delete(id int64) error

}

type dosageFormService struct{
	repo repository.DosageForm
}

func NewDosageFormService(repo repository.DosageForm) DosageFormService{
	return &dosageFormService{repo}
}

// func (s *dosageFormService) GetAll

func (s *dosageFormService) GetByID(id int64) (model.DosageForm, error){
	return s.repo.FindByID(id)
}

func (s *dosageFormService) Create(dosage model.DosageForm) (model.DosageForm, error){
	return s.repo.Create(dosage)
}

func (s *dosageFormService) Patch(id int64, data map[string] interface{}) (model.DosageForm,error){
	dosage, err := s.repo.FindByID(id)
	if err != nil {
		return model.DosageForm{},err
	}
	allowedFields := map[string] bool{
		"name": true,
		"icon":true,
		"description": true,
	}
	updates := make(map[string]interface{})

	for k, v := range data{
		if allowedFields[k]{
			updates[k] = v
		}
	}

	return s.repo.Patch(dosage.DosageFormID,updates)
}

func (s *dosageFormService) Delete(id int64) error{
	return s.repo.Delete(id)
}