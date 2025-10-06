package service

import(
	"medi-home-be/internal/app/model"
	"medi-home-be/internal/app/repository"
)

type AdminService interface{
	GetAll(page,pageSize int) (model.Pagination,error)
	GetByID(id uint) (model.Admin,error)
	TotalAdmin(role string) (int64,error)
	Create(admin model.Admin) (model.Admin,error)
	Patch(id uint, updates map[string]interface{}) (model.Admin,error)
	Delete(id uint) error
}

type adminService struct{
	repo repository.AdminRepository
}

func NewAdminService(repo repository.AdminRepository) AdminService{
	return &adminService{repo}
}

func (s *adminService) GetAll(page,pageSize int) (model.Pagination,error){
	return s.repo.FindAll(page, pageSize)
}

func (s *adminService) TotalAdmin(role string) (int64,error){
	return s.repo.TotalAdmin(role)
}

func (s *adminService) GetByID(id uint) (model.Admin,error){
	return s.repo.FindByID(id)
}

func (s *adminService) Create(admin model.Admin) (model.Admin,error){
	return s.repo.Create((admin))
}

func (s *adminService) Patch(id uint, data map[string] interface{} ) (model.Admin,error){
	admin, err := s.repo.FindByID(id)
	if err != nil{
		return model.Admin{},err
	}
	allowedFields := map[string]bool{
		"phone" :true,
		"name":true,
		"password" : true,
		"role" :true,
	}
	updates := make(map[string] interface{})
	for k,v := range data{
		if allowedFields[k]{
			updates[k] = v
		}
	}
	return s.repo.Patch(uint(admin.AdminID),updates)
}

func (s *adminService) Delete(id uint ) error{
	return s.repo.Delete(id)
}