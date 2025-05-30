package handler

import (
	"company-profile/internal/dto/request"
	"company-profile/internal/dto/response"
	"company-profile/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandler struct {
	service usecase.AuthUsecase
}

func NewAuthHandler(s usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{s}
}

// @Summary Login user
// @Tags Auth
// @Accept json
// @Produce json
// @Param data body request.LoginRequest true "Login data"
// @Success 200 {object} response.MessageResponse
// @Failure 400 {object} response.MessageResponse
// @Router /api/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.MessageResponse{
			Message: "Data tidak valid",
		})
		return
	}
	token, err := h.service.Login(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.MessageResponse{
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// @Summary Logout user
// @Tags Auth
// @Security BearerAuth
// @Produce json
// @Success 200 {object} response.MessageResponse
// @Router /api/logout [post]
func (h *AuthHandler) Logout(c *gin.Context) {
	token := c.GetHeader("Authorization")

	h.service.Logout(token)
	c.JSON(http.StatusOK, response.MessageResponse{
		Message: "Berhasil logout",
	})
}
