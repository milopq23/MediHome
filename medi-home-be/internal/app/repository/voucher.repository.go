package repository

import (
	"medi-home-be/config"
	"medi-home-be/internal/app/model"
)

type VoucherRepository interface {
	// Define voucher-related methods here
	CreateVoucher(voucher model.Voucher) (model.Voucher, error)
	GetVoucherByCode(code string) (model.Voucher, error)
	PatchVoucher(voucher model.Voucher) (model.Voucher, error)
	DeleteVoucher(id int64) error
}

type voucherRepository struct{}

func NewVoucherRepository() VoucherRepository {
	return &voucherRepository{}
}

func (r *voucherRepository) CreateVoucher(voucher model.Voucher) (model.Voucher, error) {
	err := config.DB.Create(&voucher).Error
	if err != nil {
		return model.Voucher{}, err
	}
	return voucher, nil
}

func (r *voucherRepository) GetVoucherByCode(code string) (model.Voucher, error) {
	var voucher model.Voucher
	err := config.DB.Where("code = ?", code).First(&voucher).Error	
	if err != nil {
		return model.Voucher{}, err
	}
	return voucher, nil
}

func (r *voucherRepository) PatchVoucher(voucher model.Voucher) (model.Voucher, error) {
	err := config.DB.Save(&voucher).Error
	if err != nil {
		return model.Voucher{}, err
	}	
	return voucher, nil
}

func (r *voucherRepository) DeleteVoucher(id int64) error {
	err := config.DB.Delete(&model.Voucher{}, id).Error
	if err != nil {
		return err
	}
	return nil
}	