package model

import "time"

type Medicine struct{
	MedicineID int64
	MedicineCateID int64
	DosageFormID int64
	Code string
	Name string
	Thumbnail string
	Image string
	Prescription bool
	Usage string
	Package string
	Indication string
	Adverse string
	Contraindication string
	Precaution string
	Ability string
	Pregnancy string
	DrugInteraction string
	Storage string
	Manufacturer string
	Note string
	CreatedAt time.Time
	UpdatedAt time.Time
}