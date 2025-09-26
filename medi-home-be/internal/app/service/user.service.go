package service

import (
	"fmt"
	"medi-home-be/internal/app/model"
	"medi-home-be/internal/app/repository"
	"medi-home-be/pkg/util"
	"strings"
	// "gorm.io/gorm"
	// "golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetAll() ([]model.User, error)
	GetByID(id uint) (model.User, error)
	Create(user model.User) (model.User, error)
	Update(id uint, data model.User) (model.User, error)
	Patch(id uint, updates map[string]interface{}) (model.User, error)
	Delete(id uint) error

	LoginUser(email string, password string) (string, model.User, error)
	RegisterUser(email string, password string, checkpassword string, name string, phone string) (*model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) GetAll() ([]model.User, error) {
	return s.repo.FindAll()
}

func (s *userService) GetByID(id uint) (model.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) Create(user model.User) (model.User, error) {
	return s.repo.Create(user)
}

func (s *userService) Update(id uint, data model.User) (model.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return model.User{}, err
	}
	user.Name = data.Name
	user.Phone = data.Phone
	user.Avatar = data.Avatar
	user.Password = data.Password
	user.Gender = data.Gender

	return s.repo.Update(user)
}

func (s *userService) Patch(id uint, data map[string]interface{}) (model.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return model.User{}, err
	}
	allowedFields := map[string]bool{
		"name":     true,
		"phone":    true,
		"avatar":   true,
		"password": true,
		"gender":   true,
	}
	updates := make(map[string]interface{})
	for k, v := range data {
		if allowedFields[k] {
			updates[k] = v
		}
	}
	return s.repo.Patch(uint(user.UserID), updates)
}

func (s *userService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *userService) LoginUser(email string, password string) (string, model.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", model.User{}, fmt.Errorf("invalid email or password")
	}
	// if err:= bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password)); err!=nil{
	// 	return "", model.User{}, fmt.Errorf("invalid email or password")
	// }
	if strings.Compare(user.Password, password) != 0 {
		return "", model.User{}, fmt.Errorf("invalid password")
	}

	token, err := util.GenerateJWT(user.Name, "user")
	if err != nil {
		return "", model.User{}, err
	}
	return token, user, nil
}

// func (s *userService) RegisterUser(email, password, checkpassword, name, phone string) (*model.User, error) {
// 	if password != checkpassword {
// 		return nil, fmt.Errorf("passwords do not match")
// 	}

// 	existingUser, err := s.repo.FindByEmail(email)
// 	if err == nil {
// 		return nil, fmt.Errorf("email already in use")
// 	}

// 	hashPassword, err := util.HashPassword(password)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to hash password: %w", err)
// 	}

// 	user, err := s.repo.Register(name, email, phone, hashPassword)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to create user: %w", err)
// 	}

// 	return user, nil
// }

func (s *userService) RegisterUser(email string, password string, checkpassword string, name string, phone string) (*model.User, error) {
	_, err := s.repo.FindByEmail(email)
	if err == nil {
		return nil, fmt.Errorf("email already in use")
	}
	// if err != gorm.ErrRecordNotFound {
	// 	// ❗ Đây là lỗi thật (không phải "not found")
	// 	return nil, fmt.Errorf("error checking existing user: %w", err)
	// }
	if password != checkpassword {
		return nil, fmt.Errorf("passwords do not match")
	}

	hashPassword, err := util.HashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// otp := util.GenderateOTPCode()

	// err = util.SendEmailOTP(email, otp)
	if err != nil {
		return nil, fmt.Errorf("failed to generate OTP: %w", err)
	}

	user, err := s.repo.Register(name, email, phone, hashPassword)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
