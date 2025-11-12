package repository

import (
	"medi-home-be/config"
	"medi-home-be/internal/app/model"
)

type StockTransactionRepository interface {
}

type stockTransactionRepository struct{}

func NewStockTransactionRepository() StockTransactionRepository {
	return &stockTransactionRepository{}
}

// func (r *stockTransactionRepository) GetViewLogTransaction() ([]model.StockTransaction,error){
// 	var log model.StockTransaction
// 	query := `
	
// 	`
// }

func (r *stockTransactionRepository) CreateLogTransaction(log model.StockTransaction) (model.StockTransaction, error) {
	err := config.DB.Create(&log).Error
	return log, err
}

