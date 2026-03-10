package router

import (
	"codeDrop/internal/handlers"
	"codeDrop/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	authHandler *handlers.AuthHandler,
	pasteHandler *handlers.PasteHandler,
) *gin.Engine {

	r := gin.Default()

	auth := r.Group("/auth")
	{
		auth.POST("/signup", authHandler.Signup)
		auth.POST("/login", authHandler.Login)
	}

	paste := r.Group("/paste")
	paste.Use(middleware.AuthMiddleware())
	{
		paste.POST("/", pasteHandler.CreatePaste)
		paste.DELETE("/:id", pasteHandler.DeletePaste)
		paste.GET("/me", pasteHandler.GetByUser)
	}

	r.GET("/paste/:id", pasteHandler.GetById)

	return r
}
