package router

import "github.com/gin-gonic/gin"

func AddRoutes(superRoute *gin.RouterGroup) {
	authRoutes(superRoute)
}
