package model


type Address struct {
	AddressID int64  `json:"address_id" gorm:"column:address_id; primaryKey; autoIncrement"`
	UserID    int64  `json:"user_id" gorm:"column:user_id"`
	Phone     string `json:"phone" gorm:"column:phone"`
	FullName  string `json:"full_name" gorm:"column:full_name"`
	Address   string `json:"address" gorm:"column:address"`
	IsDefault bool   `json:"is_default" gorm:"column:is_default"`
}

func (Address) TableName() string{
	return "addresses"
}