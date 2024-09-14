package routes

import (
	"app/controllers"
	"app/middleware"

	"github.com/gin-gonic/gin"
)

func BlogRoute(r *gin.Engine) {
	router := r.Group("posts")

	router.Use(middleware.Auth())

	router.GET("/:id/comments", controllers.GetComments)
	router.POST("/:id/comments", controllers.CreateComment)

	router.GET("/", controllers.GetBlogs)
	router.GET("/:id", controllers.DetailBlog)
	router.POST("/", controllers.CreateBlog)
	router.PUT("/:id", controllers.UpdateBlog)
	router.DELETE("/:id", controllers.DeleteBlog)

}
