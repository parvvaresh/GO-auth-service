package handler

import (
	"net/http"

	"auth-service/internal/domain"

	"github.com/gin-gonic/gin"
)

type AuthUsecase interface {
	SendOTP(phone string) error
	Register(phone, code, username, password string) (string, *domain.User, error)
	Login(phone, password string) (string, *domain.User, error)
}

type AuthHandler struct {
	Service AuthUsecase
}

func NewAuthHandler(s AuthUsecase) *AuthHandler {
	return &AuthHandler{Service: s}
}

type SendOTPRequest struct {
	Phone string `json:"phone" binding:"required,min=10,max=20"`
}

type RegisterRequest struct {
	Phone    string `json:"phone" binding:"required,min=10,max=20"`
	Code     string `json:"code" binding:"required,len=6"`
	Username string `json:"username" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required,min=6,max=100"`
}

type LoginRequest struct {
	Phone    string `json:"phone" binding:"required,min=10,max=20"`
	Password string `json:"password" binding:"required,min=6,max=100"`
}

func (h *AuthHandler) SendOTP(c *gin.Context) {
	var req SendOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if err := h.Service.SendOTP(req.Phone); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send otp"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "otp sent successfully",
	})
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	jwtToken, user, err := h.Service.Register(req.Phone, req.Code, req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user registered successfully",
		"token":   jwtToken,
		"user": gin.H{
			"id":       user.ID,
			"phone":    user.Phone,
			"username": user.Username,
		},
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	jwtToken, user, err := h.Service.Login(req.Phone, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login successful",
		"token":   jwtToken,
		"user": gin.H{
			"id":       user.ID,
			"phone":    user.Phone,
			"username": user.Username,
		},
	})
}
