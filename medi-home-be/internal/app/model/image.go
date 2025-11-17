package model

import ()

type Image struct {
	Image_ID   int64  `gorm:"primaryKey; column:image_id; autoIncrement" json:"image_id"`
	MedicineID int64  `gorm:"column:medicine_id; not null" json:"medicine_id"`
	ImageURL   string `gorm:"column:image_url; not null" json:"image_url"`
}

func (Image) TableName() string {
	return "medicineimage"
}
