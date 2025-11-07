package model

import ()

type OrderDetail struct {
	OrderDetailID int64   `json:"orderdetail_id" gorm:"primaryKey; autoIncrement; column:orderdetail_id"`
	InventoryID   int64   `json:"inventory_id" gorm:"column:inventory_id"`
	MedicineID    int64   `json:"medicine_id" gorm:"medicine_id"`
	OrderID       int64   `json:"order_id" gorm:"column:order_id"`
	Quantity      int64   `json:"quantity" gorm:"quantity"`
	SelectType    string  `json:"select_type" gorm:"select_type"`
	UnitPrice     float64 `json:"unit_price" gorm:"unit_price"`
}

func (OrderDetail) TableName() string {
	return "orderdetails"
}
