package dto

type UserListDTO struct {
	UserID     uint   `json:"user_id"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Gender     string `json:"gender"`
	IsVerified bool   `json:"is_verified"`
}

type UserPaginationResponseDTO struct {
	Page            int           `json:"page"`
	PageSize        int           `json:"page_size"`
	Total           int64         `json:"total"`
	TotalUserActive int64         `json:"active_user"`
	Data            []UserListDTO `json:"data"`
}
