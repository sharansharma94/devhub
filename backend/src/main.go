package main

import (
	"devhub/src/controllers"
	"devhub/src/models"

	"github.com/gin-gonic/gin"
)


func main(){
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/",controllers.FindCode)

	r.Run()
}