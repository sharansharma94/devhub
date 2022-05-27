package main

import (
	"devhub/src/controllers"
	"devhub/src/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/", controllers.FindAllCode)
	r.GET("/:id", controllers.FindCode)
	r.POST("/", controllers.CreateCode)
	r.DELETE("/:id", controllers.DeleteCode)

	r.Run()
}
