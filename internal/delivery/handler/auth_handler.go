package handler

import (
	"company-profile/internal/delivery/handler/dto"
	"company-profile/internal/domain"
	"company-profile/internal/usecase"
	"company-profile/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := &domain.User{
		Name: req.Name,
		Username: req.Username,
		Email: req.Email,
		Password: req.Password,
	}

	err := usecase.RegisterUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Registrasi berhasil!",
	})
}

func Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := usecase.LoginUser(req.Login, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
	}

	var roles []string
	for _, role := range user.Roles{
		roles = append(roles, role.Name)
	}

	token, err := utils.GenerateJWT(user.ID, roles)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}