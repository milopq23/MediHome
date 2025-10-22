package dto

type AdminListMedicineDTO struct {
	MedicineID     int64  `json:"medicine_id"`
	Code           string `json:"code"`
	Name           string `json:"name"`
	Thumbnail      string `json:"thumbnail"`
	Prescription   bool   `json:"prescription"`
	CategoryName   string `json:"category_name"`
	DosageFormName string `json:"dosageform_name"`
}

type UserListMedicineDTO struct {
	MedicineID int64  `json:"medicine_id"`
	Name       string `json:"name"`
	Thumbnail  string `json:"thumbnail"`
	// Price     int64  `json:"price"`
	Package string `json:"package"`
}

type UserDetailMedicineDTO struct {
	MedicineID       int64  `json:"medicine_id"`
	MedicineCate     string `gorm:"column:category" json:"medicinecate"`
	DosageForm       string `gorm:"column:dosageform" json:"dosageform"`
	Code             string `json:"code"`
	Name             string `json:"name"`
	Thumbnail        string `json:"thumbnail"`
	Image            string `json:"image"`
	Prescription     bool   `json:"prescription"`
	Usage            string `json:"usage"`
	Package          string `json:"package"`
	Indication       string `json:"indication"`
	Adverse          string `json:"adverse"`
	Contraindication string `json:"contraindication"`
	Precaution       string `json:"precaution"`
	Ability          string `json:"ability"`
	Pregnancy        string `json:"pregnancy"`
	DrugInteraction  string `json:"drug_interaction"`
	Storage          string `json:"storage"`
	Manufacturer     string `json:"manufacturer"`
	Note             string `json:"note"`
}
