package model

import "time"

type StockTransaction struct{
	StockTransactionID int64
	InventoryID int64
	TransactionType string
	Quantity int64
	TransactionDate time.Time
	Note string
	CreatedAt time.Time
	UpdatedAt time.Time
}