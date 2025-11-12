package model

import "time"

type StockTransaction struct {
	StockTransactionID int64     `gorm:"column:transaction_id;primarykey;autoIncrement"`
	InventoryID        int64     `gorm:"column:inventory_id; not null" json:"inventory_id"`
	TransactionType    string    `gorm:"column:transaction_type" json:"transaction_type"`
	Quantity           int64     `gorm:"column:quantity" json:"quantity"`
	TransactionDate    time.Time `gorm:"column:transaction_date; type:date; autoCreateTime"`
	Note               string    `gorm:"column:note" json:"note"`
	// CreatedAt          time.Time `gorm:"column:created_at;autoCreateTime"`
	// UpdatedAt          time.Time `gorm:"column:updated_at;autoUpdateTime"`

	Inventory *Inventory `gorm:"foreignKey:InventoryID;references:InventoryID" json:"inventory,omitempty"`
}

func (StockTransaction) TableName() string{
	return "stocktransaction"
} 