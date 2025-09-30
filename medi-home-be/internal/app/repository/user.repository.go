package repository

import (
	"fmt"
	"medi-home-be/config"
	"medi-home-be/internal/app/model"
	"medi-home-be/pkg/helper"
	// "github.com/gohugoio/hugo/resources/page"
)

type UserRepository interface {
	FindAll(page, pageSize int) (model.Pagination, error)
	FindByID(id uint) (model.User, error)
	Create(user model.User) (model.User, error)
	Patch(id uint, updates map[string]interface{}) (model.User, error)
	Delete(id uint) error

	FindByEmail(email string) (model.User, error)
	Register(name string, email string, phone string, password string, gender string) (model.User, error)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

type UserResponse struct {
	UserID     uint   `json:"user_id"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Gender     string `json:"gender"`
	IsVerified bool   `json:"is_verified"`
}

// #region List All Pagination

func (r *userRepository) FindAll(page, pageSize int) (model.Pagination, error) {
	var users []model.User
	var total int64

	pagination := model.NewPagination(page, pageSize)

	err := config.DB.Model(&model.User{}).Count(&total).Error
	if err != nil {
		return model.Pagination{}, err
	}

	pagination.Total = total

	err = config.DB.Scopes(helper.Paginate(pagination.Page, pagination.PageSize)).
		Order("user_id ASC").
		Find(&users).Error
	if err != nil {
		return model.Pagination{}, err
	}
	var userResponses []UserResponse
	for _, u := range users {
		userResponses = append(userResponses, UserResponse{
			UserID:     uint(*u.UserID),
			Name:       u.Name,
			Email:      u.Email,
			Phone:      u.Phone,
			IsVerified: u.IsVerified,
		})
	}

	pagination.Data = userResponses

	return *pagination, nil
}

// #endregion

// #region Get By ID

func (r *userRepository) FindByID(id uint) (model.User, error) {
	var user model.User
	err := config.DB.Find(&user, id).Error
	return user, err
}

// #endregion

// #region Create User
func (r *userRepository) Create(user model.User) (model.User, error) {
	err := config.DB.Debug().Create(&user).Error
	fmt.Println("Created user ID:", *user.UserID)
	return user, err
}

// #endregion

// #region Patch User
func (r *userRepository) Patch(id uint, updates map[string]interface{}) (model.User, error) {
	err := config.DB.Model(&model.User{}).Where("user_id=?", id).Updates(updates).Error
	if err != nil {
		return model.User{}, err
	}

	var updated model.User
	err = config.DB.First(&updated, id).Error
	return updated, err
}

// #endregion

// #region Delete User
func (r *userRepository) Delete(id uint) error {
	return config.DB.Delete(&model.User{}, id).Error
}

// #endregion

// #region Find By Email
func (r *userRepository) FindByEmail(email string) (model.User, error) {
	var user model.User
	err := config.DB.Where("email = ?", email).First(&user).Error
	return user, err
}

// #endregion

// #region Login
func (r *userRepository) LoginUser(email string, password string) (model.User, error) {
	var user model.User
	err := config.DB.Where("email = ? AND password = ?", email, password).First(&user).Error
	return user, err
}

// #endregion

// #region Register
func (r *userRepository) Register(name string, email string, phone string, password string, gender string) (model.User, error) {
	user := model.User{
		Name:     name,
		Email:    email,
		Phone:    phone,
		Password: password,
		Gender:   gender,
		Point:    0,
	}
	err := config.DB.Create(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, err
}

// #endregion
