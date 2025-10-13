package model

import "time"

type Medicine struct {
	MedicineID       int64     `gorm:"column:medicine_id;primarykey;autoIncrement"`
	MedicineCateID   int64     `gorm:"column:medicinecate_id;not null" json:"medicinecate_id"`
	DosageFormID     int64     `gorm:"column:dosageform_id;not null" json:"dosageform_id"`
	Code             string    `gorm:"column:code;unique;not null" json:"code"`
	Name             string    `gorm:"column:name;not null" json:"name"`
	Thumbnail        string    `gorm:"column:thumbnail" json:"thumbnail"`
	Image            string    `gorm:"column:image" json:"image"`
	Prescription     bool      `gorm:"column:prescription" json:"prescription"`
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

