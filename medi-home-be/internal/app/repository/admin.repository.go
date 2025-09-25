package repository

import (
	"medi-home-be/config"
	"medi-home-be/internal/app/model"
	
)

type AdminRepository interface{
	FindAll() ([]model.Admin,error)
	FindByID(id uint) (model.Admin,error)
	Create(admin model.Admin) (model.Admin,error)
	// Update(admin model.Admin) (model.Admin, error)
	Patch(id uint, updates map[string]interface{}) (model.Admin, error)
	Delete(id uint) error
}

type adminRepository struct{}

func NewAdminRepository() AdminRepository{
	return &adminRepository{}
}

func (r *adminRepository) FindAll() ([]model.Admin,error){
	var admin []model.Admin
	err:= config.DB.Find(&admin).Error
	return admin,err
}

func (r *adminRepository) FindByID(id uint) (model.Admin,error){
	var admin model.Admin
	err:= config.DB.Find(&admin,id).Error
	return admin,err
}

func (r *adminRepository) Create(admin model.Admin) (model.Admin,error){
	err:= config.DB.Create(&admin).Error
	return admin, err
}

func (r *adminRepository) Patch(id uint, updates map[string]interface{}) (model.Admin,error){
	var admin model.Admin
	if err := config.DB.First(&admin,id).Error; err!=nil{
		return admin, err
	}

	if err:= config.DB.Model(&model.Admin{}).Where("admin_id=?",id).Updates(updates).Error;err!=nil{
		return admin,err
	}
	if err := config.DB.First(&admin,id).Error; err!=nil{
		return admin, err
	}
	return admin, nil
}

func (r *adminRepository) Delete(id uint) error{
	return config.DB.Delete(&model.Admin{},id).Error
}