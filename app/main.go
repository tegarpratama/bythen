package main

import (
	"app/config"
	"app/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	config.ConnectDB()

	r := gin.Default()
	r.GET("/healthchecker", func(c *gin.Context) {
		message := "Its Works"
		c.JSON(http.StatusOK, gin.H{"status": "ok", "message": message})
	})

	routes.AuthRoute(r)
	routes.BlogRoute(r)

	log.Fatal(r.Run(fmt.Sprintf(":%v", config.ENV.PORT)))
}
