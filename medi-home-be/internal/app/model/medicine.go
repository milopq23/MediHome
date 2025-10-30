package model

import "time"

type Medicine struct {
	MedicineID       int64     `gorm:"column:medicine_id;primarykey;autoIncrement" json:"medicine_id"`
	MedicineCateID   int64     `gorm:"column:medicinecate_id;not null" json:"medicinecate_id"`
	DosageFormID     int64     `gorm:"column:dosageform_id;not null" json:"dosageform_id"`
	Code             string    `gorm:"column:code;unique;not null" json:"code"`
	Name             string    `gorm:"column:name;not null" json:"name"`
	Thumbnail        string    `gorm:"column:thumbnail" json:"thumbnail"`
	Image            string    `gorm:"column:image" json:"image"`
	Prescription     bool      `gorm:"column:prescription" json:"prescription"`
	UnitPerStrip     int64     `gorm:"column:unit_per_strip" json:"unitstrip"`
	UnitPerBox       int64     `gorm:"column:unit_per_box" json:"unitbox"`
	Usage            string    `gorm:"column:usage" json:"usage"`
	Package          string    `gorm:"column:package" json:"package"`
	Indication       string    `gorm:"column:indication" json:"indication"`
	Adverse          string    `gorm:"column:adverse" json:"adverse"`
	Contraindication string    `gorm:"column:contraindication" json:"contraindication"`
	Precaution       string    `gorm:"column:precaution" json:"precaution"`
	Ability          string    `gorm:"column:ability" json:"ability"`
	Pregnancy        string    `gorm:"column:pregnancy" json:"pregnancy"`
	DrugInteraction  string    `gorm:"column:drug_interaction" json:"drug_interaction"`
	Storage          string    `gorm:"column:storage" json:"storage"`
	Manufacturer     string    `gorm:"column:manufacturer" json:"manufacturer"`
	Note             string    `gorm:"column:note" json:"note"`
	CreatedAt        time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt        time.Time `gorm:"column:updated_at;autoUpdateTime"`

	MedicineCate *MedicineCate `gorm:"foreignKey:MedicineCateID;references:MedicineCateID" json:"medicine_cate,omitempty"`
	DosageForm   *DosageForm   `gorm:"foreignKey:DosageFormID;references:DosageFormID" json:"dosage_form,omitempty"`
}

func (Medicine) TableName() string {
	return "medicines"
}

type DetailMedicineVM struct {
	MedicineID       uint    `json:"medicine_id" gorm:"column:medicine_id"`
	Code             string  `json:"code" gorm:"column:code"`
	Name             string  `json:"name" gorm:"column:name"`
	PriceForStrip    float64 `json:"price_for_strip" gorm:"column:price_for_strip"`
	PriceForBox      float64 `json:"price_for_box" gorm:"column:price_for_box"`
	CategoryName     string  `json:"cate_name" gorm:"column:cate_name"`
	DosageFormName   string  `json:"dosage_name" gorm:"column:dosage_name"`
	Thumbnail        string  `json:"thumbnail" gorm:"column:thumbnail"`
	Image            string  `json:"image" gorm:"column:image"`
	Prescription     string  `json:"prescription" gorm:"column:prescription"`
	Usage            string  `json:"usage" gorm:"column:usage"`
	Package          string  `json:"package" gorm:"column:package"`
	Indication       string  `json:"indication" gorm:"column:indication"`
	Adverse          string  `json:"adverse" gorm:"column:adverse"`
	Contraindication string  `json:"contraindication" gorm:"column:contraindication"`
	Precaution       string  `json:"precaution" gorm:"column:precaution"`
	Ability          string  `json:"ability" gorm:"column:ability"`
	Pregnancy        string  `json:"pregnancy" gorm:"column:pregnancy"`
	DrugInteraction  string  `json:"drug_interaction" gorm:"column:drug_interaction"`
	Storage          string  `json:"storage" gorm:"column:storage"`
	Manufacturer     string  `json:"manufacturer" gorm:"column:manufacturer"`
	Note             string  `json:"note" gorm:"column:note"`
}
