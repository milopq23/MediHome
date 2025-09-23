package model

import "time"

type Voucher struct{
	VoucherID int64
	Code string
	Name string
	Description string
	StartDate time.Time
	EndDate time.Time
	IsActive bool
	DiscountType string
	DiscountValue int64
	MinOrderValue int64
	MaxDiscountValue int64
	UsageLimit int64
	UsedCount int64
	CreatedAt time.Time
	UpdatedAt time.Time
}