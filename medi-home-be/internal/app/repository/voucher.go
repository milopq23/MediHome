package repository

import (
	"fmt"
	"log"
	"medi-home-be/config"
	"medi-home-be/internal/app/model"
	"time"
)

type VoucherRepository interface {
	GetAllVoucher() ([]model.Voucher, error)
	CreateVoucher(voucher model.Voucher) (model.Voucher, error)
	GetVoucherByCode(code string) (model.Voucher, error)
	PatchVoucher(voucher model.Voucher) (model.Voucher, error)
	DeleteVoucher(id int64) error

	UpdateAllStatus() error
	ClassifyVoucher(code string, total float64) (float64, error)
}

type voucherRepository struct{}

func NewVoucherRepository() VoucherRepository {
	return &voucherRepository{}
}

type VoucherValidate struct {
	VoucherID   int64  `json:"voucher_id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	TypeVoucher string `json:"type"`
}

func (r *voucherRepository) GetAllVoucher() ([]model.Voucher, error) {
	var voucher []model.Voucher
	err := config.DB.Find(&voucher).Error
	if err != nil {
		return []model.Voucher{}, err
	}
	return voucher, err
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

func (r *voucherRepository) ClassifyVoucher(code string, total float64) (float64, error) {
	var voucher model.Voucher
	err := config.DB.Where("code = ?", code).First(&voucher).Error
	if err != nil {
		return 0, err
	}
	var discount float64

	if total < voucher.MinOrderValue {
		return 0, fmt.Errorf("đơn hàng chưa đạt giá trị tối thiểu để áp dụng voucher")
	}

	switch voucher.DiscountType {
	case "Phần trăm":
		// ví dụ tính phần trăm giảm
		log.Println("%", total)

		discount = total * (voucher.DiscountValue / 100)
		if discount > voucher.MaxDiscountValue {
			discount = voucher.MaxDiscountValue
		}
		log.Println("%", discount)
	case "Cố định":
		discount = voucher.MaxDiscountValue
	default:
		// log.Println("sai")
		return 0, err
	}

	return discount, err
}

func (r *voucherRepository) FindActiveVoucher() ([]model.Voucher, error) {
	var voucher []model.Voucher
	err := config.DB.Where("is_active = true").Find(&voucher).Error
	if err != nil {
		return []model.Voucher{}, err
	}
	return voucher, err
}

func (r *voucherRepository) UpdateAllStatus() error {
	now := time.Now()

	// active voucher đang có hiệu lực
	if err := config.DB.Model(&model.Voucher{}).
		Where("start_date <= ? AND end_date >= ?", now, now).
		Update("is_active", true).Error; err != nil {
		return err
	}

	// deactivate voucher hết hạn hoặc chưa tới ngày
	if err := config.DB.Model(&model.Voucher{}).
		Where("end_date < ? OR start_date > ?", now, now).
		Update("is_active", false).Error; err != nil {
		return err
	}

	var count int64
	if err := config.DB.Model(&model.Voucher{}).
		Where("is_active = ?", true).
		Count(&count).Error; err != nil {
		return err
	}
	log.Printf("Hiện có %d voucher đang còn hiệu lực\n", count)

	return nil
}
