package model

import ()

type DosageForm struct {
	DosageFormID int64  `gorm:"column:dosage_form_id;primarykey;autoIncrement"`
	Name         string `gorm:"column:name;not null" json:"name"`
	Description  string `gorm:"column:description" json:"description"`
}
