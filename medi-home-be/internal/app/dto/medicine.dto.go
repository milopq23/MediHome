package dto

type ListMedicineDTO struct {
	MedicineID     int64  `json:"medicine_id"`
	Code           string `json:"code"`
	Name           string `json:"name"`
	Thumbnail      string `json:"thumbnail"`
	Prescription   bool   `json:"prescription"`
	CategoryName   string `json:"category_name"`
	DosageFormName string `json:"dosageform_name"`
}
