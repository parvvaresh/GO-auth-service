package handler

import (
	"net/http"

	"auth-service/internal/domain"

	"github.com/gin-gonic/gin"
)

type ProfileUsecase interface {
	Get(userID int) (*domain.Profile, error)
	Update(profile *domain.Profile) error
}

type ProfileHandler struct {
	Service ProfileUsecase
}

func NewProfileHandler(s ProfileUsecase) *ProfileHandler {
	return &ProfileHandler{Service: s}
}

type UpdateProfileRequest struct {
	FirstName string `json:"first_name" binding:"required,min=2,max=100"`
	LastName  string `json:"last_name" binding:"required,min=2,max=100"`
	Bio       string `json:"bio" binding:"max=500"`
}

func (h *ProfileHandler) Me(c *gin.Context) {
	userID := c.GetInt("user_id")

	profile, err := h.Service.Get(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "profile not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"profile": profile,
	})
}

func (h *ProfileHandler) UpdateMe(c *gin.Context) {
	userID := c.GetInt("user_id")

	var req UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	profile := &domain.Profile{
		UserID:    userID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Bio:       req.Bio,
	}

	if err := h.Service.Update(profile); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not update profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "profile updated successfully",
	})
}
