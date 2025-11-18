package service

import (
	"medi-home-be/internal/app/model"
	"medi-home-be/internal/app/repository"
)

type VoucherService interface {
	FindVoucherActive(total float64) ([]model.Voucher, error)

	GetVoucherByCode(code string) (model.Voucher, error)

	GetAllVoucher() ([]model.Voucher, error)
	GetDetailVoucher(id int64) (model.Voucher, error)
	CreateVoucher(voucher model.Voucher) (model.Voucher, error)
	UpdateVoucher(voucher_id int64, data map[string]interface{}) (model.Voucher, error)
	DeleteVoucher(id int64) error
}

type voucherService struct {
	repo repository.VoucherRepository
}

func NewVoucherService(repo repository.VoucherRepository) VoucherService {
	return &voucherService{repo}
}

func (s *voucherService) FindVoucherActive(total float64) ([]model.Voucher, error) {
	return s.repo.FindActiveVoucher(total)
}

func (s *voucherService) GetVoucherByCode(code string) (model.Voucher, error) {
	return s.repo.GetVoucherByCode(code)
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

func (s *voucherService) GetAllVoucher() ([]model.Voucher, error) {
	return s.repo.GetAllVoucher()
}

func (s *voucherService) GetDetailVoucher(id int64) (model.Voucher, error) {
	return s.repo.GetDetailVoucher(id)
}

func (s *voucherService) CreateVoucher(voucher model.Voucher) (model.Voucher, error) {
	return s.repo.CreateVoucher(voucher)
}

func (s *voucherService) UpdateVoucher(voucher_id int64, data map[string]interface{}) (model.Voucher, error) {
	return s.repo.UpdateVoucher(voucher_id, data)
}

func (s *voucherService) DeleteVoucher(id int64) error {
	return s.repo.DeleteVoucher(id)
}
