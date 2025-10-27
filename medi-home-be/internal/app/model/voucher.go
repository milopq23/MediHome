package model

import "time"

type Voucher struct {
	VoucherID        int64     `gorm:"column:voucher_id;primarykey;autoIncrement"`
	Code             string    `gorm:"column:code;unique;not null" json:"code"`
	Name             string    `gorm:"column:name;not null" json:"name"`
	Description      string    `gorm:"column:description" json:"description"`
	StartDate        time.Time `gorm:"column:start_date;not null" json:"start_date"`
	EndDate          time.Time `gorm:"column:end_date;not null" json:"end_date"`
	IsActive         bool      `gorm:"column:is_active;not null" json:"is_active"`
	DiscountType     string    `gorm:"column:discount_type;not null" json:"discount_type"` // percentage or fixed
	DiscountValue    int64     `gorm:"column:discount_value;not null" jsson:"discount_value"`
	MinOrderValue    int64     `gorm:"column:min_order_value;not null" json:"min_order_value"`
	MaxDiscountValue int64     `gorm:"column:max_discount_value;not null" json:"max_discount_value"`
	UsageLimit       int64     `gorm:"column:usage_limit;not null" json:"usage_limit"`
	UsedCount        int64     `gorm:"column:used_count;not null" json:"used_count"`
	CreatedAt        time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt        time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (Voucher) TableName() string {
	return "vouchers"
}
