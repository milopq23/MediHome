package repository

import (
	"medi-home-be/config"
	"medi-home-be/internal/app/model"
)

type BatchSellingRepository interface {
}

type batchSellingRepository struct{}

func NewBatchSellingRepository() BatchSellingRepository {
	return &batchSellingRepository{}
}

type BatchSelling struct {
	MedicineID    int64   `json:"medicine_id"`
	NameMedicine  string  `json:"name"`
	BatchNumber   int64   `json:"batch_number"`
	StockQuantity int64   `json:"stock_quantity"`
	PriceStrip    float64 `json:"price_strip"`
	PriceBox      float64 `json:"price_box"`
}

func (r *batchSellingRepository) CurrentBatch() ([]BatchSelling, error) {
	var curentBatch []BatchSelling
	query := `
		select m.medcine_id ,m.name, i.batch_number, bs.stock_quantity,
		round((i.import_price * (1 + i.mark_up/100))/m.unit_per_strip) as price_for_strip,
		round((i.import_price * (1 + i.mark_up/100))) as price_for_box,
		from batchsellings bs
		join inventories i on i.inventory_id = bs.inventory_id
		join medicines m on m.medicine_id = bs.medicine_id
	`
	err := config.DB.Raw(query).Scan(&curentBatch).Error
	return curentBatch, err
}

func (r *batchSellingRepository) CreateBatch(batch model.BatchSelling) (model.BatchSelling, error) {
	err := config.DB.Create(&batch).Error
	return batch, err
}

func (r *batchSellingRepository) UpdateBatchSelling(medicine_id, inventory_id int64) (model.BatchSelling, error) {
	var batchselling model.BatchSelling
	err := config.DB.Model(&batchselling).Where("medicine_id  = ?", medicine_id).Update("inventory_id", inventory_id).Error
	if err != nil {
		return model.BatchSelling{}, err
	}
	err = config.DB.Where("medicine_id = ?", medicine_id).First(&batchselling).Error
	if err != nil {
		return model.BatchSelling{}, err
	}

	return batchselling, nil
}
