package service

import (
	"errors"
	"fmt"
	"medi-home-be/internal/app/dto"
	"medi-home-be/internal/app/model"
	"medi-home-be/internal/app/repository"
	"medi-home-be/pkg/util"

	"golang.org/x/crypto/bcrypt"
	// "gorm.io/gorm"
	// "golang.org/x/crypto/bcrypt"
)

type UserService interface {
	TotalActive() (int64, error)
	GetAll(page, pageSize int) (dto.UserPaginationResponseDTO, error)
	GetByID(id uint) (model.User, error)
	Create(user model.User) (model.User, error)
	Patch(id uint, updates map[string]interface{}) (model.User, error)
	Delete(id uint) error
	LoginUser(email string, password string) (string, model.User, error)
	RegisterUser(email string, password string, checkpassword string, name string, phone string, gender string) (*model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) TotalActive() (int64, error) {
	return s.repo.TotalActive()
}

func (s *userService) GetAll(page, pageSize int) (dto.UserPaginationResponseDTO, error) {

	pagination, err := s.repo.GetAll(page, pageSize)
	if err != nil {
		return dto.UserPaginationResponseDTO{}, err
	}

	users, ok := pagination.Data.([]model.User)
	if !ok {
		return dto.UserPaginationResponseDTO{}, fmt.Errorf("failed load pagination data")
	}

	total_active, err := s.repo.TotalActive()
	if err != nil {
		return dto.UserPaginationResponseDTO{}, err
	}

	// mapping model -> dto
	var result []dto.UserListDTO
	for _, u := range users {
		result = append(result, dto.UserListDTO{
			UserID:     uint(u.UserID),
			Name:       u.Name,
			Email:      u.Email,
			Phone:      u.Phone,
			Gender:     u.Gender,
			IsVerified: u.IsVerified,
		})
	}

	user := dto.UserPaginationResponseDTO{
		Page:            pagination.Page,
		PageSize:        pagination.PageSize,
		Total:           pagination.Total,
		Data:            result,
		TotalUserActive: total_active,
	}

	return user, nil
}

func (s *userService) GetByID(id uint) (model.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) Create(user model.User) (model.User, error) {
	return s.repo.Create(user)
}

func (s *userService) Patch(id uint, data map[string]interface{}) (model.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return model.User{}, err
	}
	allowedFields := map[string]bool{
		"name":        true,
		"email":       true,
		"phone":       true,
		"avatar":      true,
		"pass	word":   true,
		"gender":      true,
		"is_verified": true,
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
		// Không tiết lộ lý do cụ thể để tránh lộ thông tin
		return "", model.User{}, fmt.Errorf("invalid email or password")
	}

	// So sánh password (hash)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", model.User{}, fmt.Errorf("invalid email or password")
	}

	// Tạo JWT đầy đủ
	token, err := util.GenerateJWT(uint(user.UserID), user.Name, "user")
	if err != nil {
		return "", model.User{}, fmt.Errorf("could not generate token")
	}

	return token, user, nil
}

func (s *userService) RegisterUser(email string, password string, checkpassword string, name string, phone string, gender string) (*model.User, error) {
	if password != checkpassword {
		return nil, errors.New("mật khẩu không khớp")
	}
	_, err := s.repo.FindByEmail(email)
	if err == nil {
		return nil, errors.New("email đã tồn tại")
	}
	hashPassword, err := util.HashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}
	user, err := s.repo.Register(name, email, phone, hashPassword, gender)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
