package dto

type DosageFormShortDTO struct {
	DosageFormID int64  `json:"dosageform_id"`
	Name         string `json:"name"`
}

type DosageFormList struct {
	DosageFormID int64  `json:"dosageform_id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
}
