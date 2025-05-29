package handler

import (
	"company-profile/internal/domain"
	"company-profile/internal/dto/request"
	"company-profile/internal/dto/response"
	"company-profile/internal/usecase"
	"company-profile/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Register godoc
// @Summary Register new user
// @Description Register with email, username, password
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param   request body  request.RegisterRequest true "Register data"
// @Success 201 {object} response.UserResponse
// @Failure 400 {object} map[string]string
// @Router /api/register [post]
func Register(c *gin.Context) {
	var req request.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := &domain.User{
		Name:     req.Name,
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	err := usecase.RegisterUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusCreated, response.MessageResponse{
		Message: "Registrasi berhasil",
	})
}

// Login godoc
// @Summary Login user
// @Description Login menggunakan email/username dan password
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body request.LoginRequest true "Login data"
// @Success 200 {object} response.LoginResponse
// @Failure 401 {object} map[string]string
// @Router /api/login [post]
func Login(c *gin.Context) {
	var req request.LoginRequest
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
	for _, role := range user.Roles {
		roles = append(roles, role.Name)
	}

	token, err := utils.GenerateJWT(user.ID, roles)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, response.MessageResponse{
		Message: "Login success",
		Data:    token,
	})
}
