package service

import (
	"medi-home-be/internal/app/model"
	"medi-home-be/internal/app/repository"
)

type UserService interface{
	GetAll() ([]model.User,error)
	GetByID(id uint) (model.User,error)
	Create(user model.User) (model.User,error)
	Update(id uint, data model.User) (model.User,error)
	Patch(id uint, updates map[string] interface{}) (model.User, error)
	Delete(id uint) error
}

type userService struct{
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService{
	return &userService{repo}
}

func (s *userService) GetAll() ([]model.User,error){
	return s.repo.FindAll()
}

func (s *userService) GetByID(id uint) (model.User,error){
	return s.repo.FindByID(id)
}

func (s *userService) Create(user model.User) (model.User,error){
	return s.repo.Create(user)
}

func (s *userService) Update(id uint, data model.User) (model.User,error){
	user,err := s.repo.FindByID(id)
	if err != nil {
		return model.User{},err
	}
	user.Name = data.Name
	user.Phone = data.Phone
	user.Avatar = data.Avatar
	user.Password = data.Password
	user.Gender = data.Gender

	return  s.repo.Update(user)
}

func (s *userService) Patch(id uint, data map[string] interface{} ) (model.User,error){
	user, err := s.repo.FindByID(id)
	if err != nil {
		return model.User{},err
	}
	allowedFields := map[string]bool{
		"name" : true,
		"phone" :true,
		"avatar" :true,
		"password" : true,
		"gender" :true,
	}
	updates := make(map[string] interface{})
	for k,v := range data {
		if allowedFields[k]{
			updates[k] = v
		}
	}
	return s.repo.Patch(uint(user.UserID),updates)
}

func (s *userService) Delete(id uint) error{
	return s.repo.Delete(id)
}