package model

import ()

type DosageForm struct {
	DosageFormID int64  `gorm:"column:dosageform_id;primarykey;autoIncrement" json:"dosageform_id"`
	Name         string `gorm:"column:name;not null" json:"name"`
	Description  string `gorm:"column:description" json:"description"`
}

func (DosageForm) TableName() string{
	return "dosageforms"
}
