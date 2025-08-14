package router

import (
	"open_gamba/internal/auth"

	"github.com/gin-gonic/gin"
)

func authRoutes(superRoute *gin.RouterGroup) {
	authRouter := superRoute.Group("/api/v1/auth")
	{
		authRouter.POST("/login", auth.Login)
		authRouter.POST("/signup", auth.SignUp)
	}
}
