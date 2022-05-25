package controllers

import (
	"devhub/src/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindCode(ctx *gin.Context) {
	// code := models.Code{ID: 1, Code: "console.log('hello world')"}
	// models.DB.Find(&code)
	var code []models.Code
	DB := models.GetDB()

	// DB.Find(&code)

	fmt.Println(DB)
	// DB.Create(&models.Code{ID: 3, Code: "alert('test')"})
	ctx.JSON(http.StatusOK, gin.H{"data": code})
}
