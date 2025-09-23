package model

import()

type OrderDetail struct{
	OrderDetailID int64
	InventoryID int64
	OrderID int64
	Quantity int64
	UnitPrice int64
}