package model

type MedicineCate struct {
	MedicineCateID int64  `gorm:"column:medicinecate_id;primarykey;autoIncrement" json:"medicinecate_id"`
	Name           string `gorm:"column:name;not null" json:"name"`
	Icon           string `gorm:"column:icon" json:"icon"`
	ParentID       *int64 `gorm:"column:parent_id" json:"parent_id"`

	Parent   *MedicineCate  `gorm:"foreignKey:ParentID;references:MedicineCateID" json:"parent,omitempty"`
	Children []MedicineCate `gorm:"foreignKey:ParentID;references:MedicineCateID" json:"children,omitempty"`
}

func (MedicineCate) TableName() string {
	return "medicinecates"
}
