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
	c.SetCookie(
		"refresh_token",
		refresh,
		7 * 24 * 3600,
		"/",
		"localhost",
		false,
		true,
	)

	c.JSON(200, gin.H{"access_token": access})
}

func(h *AuthHandler) RefreshToken(c *gin.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	
}