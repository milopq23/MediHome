package repository

import (
	"medi-home-be/config"
	"medi-home-be/internal/app/model"
)

type StockTransactionRepository interface {
	CreateLogTransaction(log model.StockTransaction) (model.StockTransaction, error)
}

type stockTransactionRepository struct{}

func NewStockTransactionRepository() StockTransactionRepository {
	return &stockTransactionRepository{}
}


func (r *stockTransactionRepository) CreateLogTransaction(log model.StockTransaction) (model.StockTransaction, error) {
	if err := config.DB.Create(&log).Error; err != nil {
		return model.StockTransaction{}, err
	}
	return log, nil
}
