package handlers

import (
	"codeDrop/internal/dto"
	"codeDrop/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Service *service.AuthService
}

func (a *AuthHandler) Signup (c *gin.Context) {
	var req dto.SignupRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid input"})
		return
	}

	err := a.Service.Signup(req.UserName, req.Email, req.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "user created successfully"})
}

func (a *AuthHandler) Login (c *gin.Context) {
	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid credentials"})
		return
	}

	access, refresh, err := a.Service.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	res := dto.AuthResponse{
		AccessToken: access,
		RefreshToken: refresh,
	}

	c.JSON(200, res)
}