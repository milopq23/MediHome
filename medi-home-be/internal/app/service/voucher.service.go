package service

import (
	"medi-home-be/internal/app/model"
	"medi-home-be/internal/app/repository"
)

type VoucherService interface {
	GetAllVoucher() ([]model.Voucher, error)
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

func (s *voucherService) GetAllVoucher() ([]model.Voucher, error) {
	return s.repo.GetAllVoucher()
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

// func (s *voucherService) ValidateVoucher() (model.Voucher,error){
// 	vouchers, err := s.repo.GetAllVoucher()
// 	if err != nil {
// 		return model.Voucher{},err
// 	}
// 	now := time.Now()
// 	for _, v := range vouchers {
//         active := now.After(v.StartDate) && now.Before(v.EndDate)
//         if v.IsActive != active {
//             v.IsActive = active
//             if err := .Model(&v).Update("is_active", active).Error; err != nil {
//                 return err
//             }
//         }
//     }
// 	return vouchers,nill
// }
