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
		userGroup.POST("/register", handler.Register)
		userGroup.POST("/login", handler.Login)
		userGroup.GET("/me", handler.GetMyProfile)
		userGroup.PATCH("/me", handler.UpdateMyProfile)
		userGroup.PATCH("/me/password", handler.ChangePassword)
		userGroup.DELETE("/me", handler.DeleteAccount)
		userGroup.POST("/logout", handler.Logout)
		userGroup.PATCH("/:userId/role", handler.UpdateUserRole)
	}

	return router
}
