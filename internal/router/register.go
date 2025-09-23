package router

import (
	"github.com/NoANameGroup/DAOld-Backend/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	// UserApi
	userGroup := router.Group("/api/user")
	{
		userGroup.POST("/register", handler.Register)
		userGroup.POST("/login", handler.Login)
	}

	return router
}
