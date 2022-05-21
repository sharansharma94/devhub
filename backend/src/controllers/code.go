package controllers

import (
	"devhub/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)




func FindCode(ctx *gin.Context){
	var code []models.Code
	models.DB.Find(&code)

	ctx.JSON(http.StatusOK,gin.H{"data":code})
}