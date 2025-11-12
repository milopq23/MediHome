package model

import "time"

type Inventory struct {
	InventoryID    int64     `gorm:"primaryKey; column:inventory_id; autoIncrement" json:"inventory_id"`
	MedicineID     int64     `gorm:"column:medicine_id; not null" json:"medicine_id"`
	BatchNumber    string    `gorm:"column:batch_number; not null" json:"batchnumber"`
	ImportPrice    float64   `gorm:"column:import_price; not null" json:"importprice"`
	MarkUp         float64   `gorm:"column:mark_up; not null" json:"markup"`
	ExpirationDate time.Time `gorm:"column:expiration_date; type:date" json:"expiration_date"`
	Quantity       int64     `gorm:"column:quantity" json:"quantity"`
	Status         string    `gorm:"column:status" json:"status"`
	CreatedAt      time.Time `gorm:"column:created_at;  autoCreateTime"`
	UpdatedAt      time.Time `gorm:"column:updated_at;autoUpdateTime"`

	Medicine *Medicine `gorm:"foreignKey:MedicineID; references:MedicineID" json:"medicine,omitempty"`
}

func (Inventory) TableName() string {
	return "inventories"
}
