package service

import (
	"medi-home-be/internal/app/model"
	"medi-home-be/internal/app/repository"
)

type VoucherService interface {
	CreateVoucher(voucher model.Voucher) (model.Voucher, error)
	GetVoucherByCode(code string) (model.Voucher, error)
	PatchVoucher(voucher model.Voucher) (model.Voucher, error)
	DeleteVoucher(id int64) error
}

type voucherService struct {
	repo repository.VoucherRepository
}

func NewVoucherService(repo repository.VoucherRepository) VoucherService {
	return &voucherService{repo}
}

func (s *voucherService) CreateVoucher(voucher model.Voucher) (model.Voucher, error) {
	return s.repo.CreateVoucher(voucher)
}

func (s *voucherService) GetVoucherByCode(code string) (model.Voucher, error) {
	return s.repo.GetVoucherByCode(code)
}

func (s *voucherService) PatchVoucher(voucher model.Voucher) (model.Voucher, error) {
	return s.repo.PatchVoucher(voucher)
}

func (s *voucherService) DeleteVoucher(id int64) error {
	return s.repo.DeleteVoucher(id)
}
