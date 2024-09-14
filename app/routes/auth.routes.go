package routes

import (
	"app/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoute(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
}
