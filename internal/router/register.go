package router

import (
	"github.com/NoANameGroup/DAOld-Backend/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	// UserApi
	userGroup := router.Group("/api/users")
	{
		userGroup.POST("/", handler.CreateUser)
		userGroup.GET("/me", handler.GetMyProfile)
		userGroup.PATCH("/me", handler.UpdateMyProfile)
		userGroup.PATCH("/me/password", handler.UpdateMyPassword)
		userGroup.DELETE("/me", handler.DeleteMyAccount)
		userGroup.PATCH("/:userId/role", handler.UpdateUserRole)
		//userGroup.PATCH("/me/email", handler.UpdateMyEmail)
		//userGroup.PATCH("/me/phone", handler.UpdateMyPhone)
		userGroup.PATCH("/me/username", handler.UpdateMyUsername)
	}

	// SessionApi
	sessionGroup := router.Group("/api/sessions")
	{
		sessionGroup.POST("/", handler.CreateSession)
		sessionGroup.DELETE("/", handler.DeleteSession)
	}

	return router
}
