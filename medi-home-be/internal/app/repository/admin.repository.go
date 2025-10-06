package repository

import (
	"medi-home-be/config"
	"medi-home-be/internal/app/model"
	"medi-home-be/pkg/helper"
)

type AdminRepository interface{
	FindAll(page, pageSize int) (model.Pagination,error)
	FindByID(id uint) (model.Admin,error)
	TotalAdmin(role string) (int64,error)
	Create(admin model.Admin) (model.Admin,error)
	Patch(id uint, updates map[string]interface{}) (model.Admin, error)
	Delete(id uint) error
}

type adminRepository struct{}

func NewAdminRepository() AdminRepository{
	return &adminRepository{}
}

type AdminResponse struct {
	AdminID     uint   `json:"admin_id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Role	   string `json:"role"`
}

func (r *adminRepository) FindAll(page,pageSize int) (model.Pagination,error){
	var admin []model.Admin
	var total int64

	pagination := model.NewPagination(page, pageSize)

	err := config.DB.Model(&model.Admin{}).Count(&total).Error
	if err != nil {
		return model.Pagination{}, err
	}
	pagination.Total = total

	err = config.DB.Scopes(helper.Paginate(pagination.Page, pagination.PageSize)).
		Order("admin_id ASC").
		Find(&admin).Error

	if err != nil {
		return model.Pagination{}, err
	}

	var adminResponses []AdminResponse
	for _, a := range admin {
		adminResponses = append(adminResponses, AdminResponse{
			AdminID: uint(a.AdminID),
			Name : a.Name,
			Email:   a.Email,
			Phone:   a.Phone,
			Role:    a.Role,
		})
	}
	pagination.Data = adminResponses

	return *pagination, nil
}

func (r *adminRepository) TotalAdmin(role string) (int64, error) {
	var count int64
	err:=config.DB.Model(&model.Admin{}).Where("role = ?", role).Count(&count).Error
	return count,err
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