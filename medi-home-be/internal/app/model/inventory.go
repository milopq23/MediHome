package model

import "time"

type Inventory struct{
	InventoryID int64
	MedicineID int64
	BatchNumber string
	ImportPrice int64
	MarkUp int64
	ExpirationDate time.Time
	Quantity int64
	CreatedAt time.Time
	UpdatedAt time.Time
}