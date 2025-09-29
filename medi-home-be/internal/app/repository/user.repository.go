package repository

import (
	// "fmt"
	// "math/rand"
	"fmt"
	"medi-home-be/config"
	"medi-home-be/internal/app/model"
	// "time"
)

type UserRepository interface {
	FindAll() ([]model.User, error)
	FindByID(id uint) (model.User, error)
	Create(user model.User) (model.User, error)
	Update(user model.User) (model.User, error)
	Patch(id uint, updates map[string]interface{}) (model.User, error)
	Delete(id uint) error


	FindByEmail(email string) (model.User, error)
	Register(name string, email string, phone string, password string,gender string) (model.User, error)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) FindAll() ([]model.User, error) {
	var users []model.User
	err := config.DB.Find(&users).Error
	return users, err
}

func (r *userRepository) FindByID(id uint) (model.User, error) {
	var user model.User
	err := config.DB.Find(&user, id).Error
	return user, err
}

func (r *userRepository) Create(user model.User) (model.User, error) {
	err := config.DB.Debug().Create(&user).Error
	fmt.Println("Created user ID:", *user.UserID)
	return user, err
}

func (r *userRepository) Update(user model.User) (model.User, error) {
	err := config.DB.Save(&user).Error
	return user, err
}

func (r *userRepository) Patch(id uint, updates map[string]interface{}) (model.User, error) {
	var user model.User

	if err := config.DB.First(&user, id).Error; err != nil {
		return user, err
	}

	// if err := config.DB.Model(&user).Updates(updates).Error; err != nil {
	// 	return user, err
	// }
	if err := config.DB.Model(&model.User{}).Where("user_id=?", id).Updates(updates).Error; err != nil {
		return user, err
	}

	if err := config.DB.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) Delete(id uint) error {
	return config.DB.Delete(&model.User{}, id).Error
}

func (r *userRepository) LoginUser(email string, password string) (model.User, error) {
	var user model.User
	err := config.DB.Where("email = ? AND password = ?", email, password).First(&user).Error
	return user, err
}

func (r *userRepository) FindByEmail(email string) (model.User, error) {
	var user model.User
	err := config.DB.Where("email = ?", email).First(&user).Error
	return user, err
}

func (r *userRepository) Register(name string, email string, phone string, password string,gender string) (model.User, error) {
	user := model.User{
		Name : name,
		Email : email,
		Phone : phone,
		Password : password,
		Gender: gender,
		Point : 0,
	}
	err := config.DB.Create(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, err
}

