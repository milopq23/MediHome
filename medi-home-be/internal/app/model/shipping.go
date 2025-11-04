package model

import ()

type Shipping struct {
	ShippingID int64   `json:"shipping_id" gorm:"column:shipping_id;primaryKey;autoIncrement"`
	Name       string  `json:"name" column:"name"`
	Price      float64 `json:"price" column:"price"`
}

func (Shipping) TableName() string{
	return "shipping"
}