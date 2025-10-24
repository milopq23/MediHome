package handler

import (
	"log"
	"medi-home-be/internal/app/model"
	"medi-home-be/internal/app/service"
	"medi-home-be/pkg/util"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.UserService
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
type UserTokenDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

type LoginResponse struct {
	Token     string       `json:"token"`
	UserToken UserTokenDTO `json:"user"`
}

type RegisterRequest struct {
	Email         string `json:"email" binding:"required,email"`
	Password      string `json:"password" binding:"required,min=6"`
	CheckPassword string `json:"checkpassword" binding:"required,eqfield=Password"`
	Name          string `json:"name" binding:"required"`
	Phone         string `json:"phone" binding:"required"`
	Gender        string `json:"gender" binding:"required"`
}

// type LoginResponse struct {
// 	Token string        `json:"token"`
// 	User  UserDTO `json:"user"`
// }

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service}
}

func (h *UserHandler) TotalActive(c *gin.Context) {
	total, err := h.service.TotalActive()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, total)
}

func (h *UserHandler) GetAll(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	pagination, err := h.service.GetAll(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, pagination)
}

func (h *UserHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	log.Print(id)
	user, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) Create(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newUser, err := h.service.Create(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newUser)
}

func (h *UserHandler) Patch(c *gin.Context) {
	claimsRaw, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No claims found"})
		return
	}
	log.Print("claimsRaw", claimsRaw)
	claims := claimsRaw.(*util.Claims)
	log.Print(claims)

	var input map[string]interface{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu lỗi"})
		return
	}
	updatedUser, err := h.service.Patch(uint(claims.UserID), input)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patch lỗi"})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

func (h *UserHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

//User

func (h *UserHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	token, user, err := h.service.LoginUser(strings.ToLower(req.Email), req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("token", token, 3600, "/", "localhost", false, true)
	log.Printf(token)
	resp := LoginResponse{
		Token: token,
		UserToken: UserTokenDTO{
			ID:   uint(user.UserID),
			Name: user.Name,
			Role: "user",
		},
	}
	c.JSON(http.StatusOK, resp)
}

func (h *UserHandler) Register(c *gin.Context) {
	var user RegisterRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newUser, err := h.service.RegisterUser(user.Email, user.Password, user.CheckPassword, user.Name, user.Phone, user.Gender)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newUser)
}

func (h *UserHandler) Profile(c *gin.Context) {
	claimsRaw, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No claims found"})
		return
	}
	log.Println("Lỗi nhận claim",claimsRaw)
	claims := claimsRaw.(*util.Claims)
	log.Println(claims)
	user, err := h.service.GetByID(claims.UserID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) LogOut(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
