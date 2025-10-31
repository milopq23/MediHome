package model

type BatchSelling struct {
	MedicineID    int64 `gorm:"medicine_id" json:"medicine_id"`
	InventoryID   int64 `gorm:"inventory_id" json:"inventory_id"`
	StockQuantity int64 `gorm:"stock_quantity" json:"stock_quantity"`
	UpdatedBy     int64 `gorm:"updated_by" json:"updated_by"`
}

func (BatchSelling) TableName() string {
	return "batchsellings"
}
