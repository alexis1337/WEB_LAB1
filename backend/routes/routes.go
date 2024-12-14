package routes

import (
	"news_app/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, newsController *controllers.NewsController) {
	router.Static("/static", "../frontend/public")

	router.NoRoute(func(c *gin.Context) {
		c.File("./frontend/public/index.html")
	})

	api := router.Group("/api/news")
	{
		api.GET("", newsController.GetNews)
		api.POST("/create", newsController.CreateNews)
		api.PUT("/update/:id", newsController.UpdateNews)
		api.DELETE("/delete/:id", newsController.DeleteNews)
	}
}
