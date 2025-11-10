package dto

type UserListDTO struct {
	UserID     uint   `json:"user_id"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Gender	 string `json:"gender"`
	IsVerified bool   `json:"is_verified"`
}