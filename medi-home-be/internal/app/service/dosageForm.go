package service

import (
	"medi-home-be/internal/app/dto"
	"medi-home-be/internal/app/model"
	"medi-home-be/internal/app/repository"
)

type DosageFormService interface {
	GetAllShort() ([]dto.DosageFormShortDTO, error)
	GetAll() ([]model.DosageForm, error)
	GetByID(id int64) (model.DosageForm, error)
	Create(dosage model.DosageForm) (model.DosageForm, error)
	Patch(id int64, data map[string]interface{}) (model.DosageForm, error)
	Delete(id int64) error
}

type dosageFormService struct {
	repo repository.DosageForm
}

func NewDosageFormService(repo repository.DosageForm) DosageFormService {
	return &dosageFormService{repo}
}

func (s *dosageFormService) GetAllShort() ([]dto.DosageFormShortDTO, error) {
	models, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	// mapping model -> dto
	var result []dto.DosageFormShortDTO
	for _, m := range models {
		result = append(result, dto.DosageFormShortDTO{
			DosageFormID: m.DosageFormID,
			Name:         m.Name,
		})
	}
	return result, nil
}

func (s *dosageFormService) GetAll() ([]model.DosageForm, error) {
	return s.repo.GetAll()
}

func (s *dosageFormService) GetByID(id int64) (model.DosageForm, error) {
	return s.repo.FindByID(id)
}

func (s *dosageFormService) Create(dosage model.DosageForm) (model.DosageForm, error) {
	return s.repo.Create(dosage)
}

func (s *dosageFormService) Patch(dosageform_id int64, data map[string]interface{}) (model.DosageForm, error) {
	// dosage, err := s.repo.FindByID(id)
	// if err != nil {
	// 	return model.DosageForm{}, err
	// }
	// allowedFields := map[string]bool{
	// 	"name":        true,
	// 	"icon":        true,
	// 	"description": true,
	// }
	// updates := make(map[string]interface{})

	// for k, v := range data {
	// 	if allowedFields[k] {
	// 		updates[k] = v
	// 	}
	// }

	return s.repo.Patch(dosageform_id, data)
}

func (s *dosageFormService) Delete(id int64) error {
	return s.repo.Delete(id)
}
