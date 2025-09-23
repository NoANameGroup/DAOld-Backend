package router

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	// CommentApi
	//commentGroup := router.Group("/api/comment")
	//{
	//	commentGroup.POST("/add", handler.CreateComment)
	//}
	//
	//searchGroup := router.Group("/api/search")
	//{
	//	searchGroup.GET("/recent", handler.GetSearchHistory)
	//	searchGroup.POST("", handler.LogSearch)
	//}

	return router
}
