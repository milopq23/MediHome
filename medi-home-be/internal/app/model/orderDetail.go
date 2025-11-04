package model

import ()

type OrderDetail struct {
	OrderDetailID int64 `json:"orderdetail_id" gorm:"primaryKey, autoIncrement; column:orderdetail_id"`
	InventoryID   int64 `json:"inventory_id" gorm:"primaryKey, autoIncrement; column:orderdetail_id"`
	OrderID       int64 `json:"order_id" gorm:"column:order_id"`
	Quantity      int64 `json:"quantity" gorm:"quantity"`
	UnitPrice     int64 `json:"unit_price" gorm:"unit_price"`
}
func (OrderDetail) TableName() string {
	return "orderdetails"
}