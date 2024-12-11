package api

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/news", CreateNewsHandler)
		api.GET("/news", FetchNewsHandler)
		api.GET("/news/:id", FetchNewsByIDHandler)
		api.GET("/news/category", FetchNewsByCategory)
		api.GET("/category", FetchCategories)
		api.PUT("/news/:id", UpdateNewsHandler)
		api.DELETE("/news/:id", DeleteNewsHandler)
	}
}
