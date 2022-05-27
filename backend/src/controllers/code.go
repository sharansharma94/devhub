package controllers

import (
	"devhub/src/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindAllCode(ctx *gin.Context) {
	// code := models.Code{ID: 1, Code: "console.log('hello world')"}
	// models.DB.Find(&code)
	var code []models.Code
	DB := models.GetDB()

	DB.Find(&code)

	fmt.Println(code)

	ctx.JSON(http.StatusOK, code)
}

func FindCode(ctx *gin.Context) {

	DB := models.GetDB()

	id := ctx.Params.ByName("id")
	var code models.Code

	DB.First(&code, id)

	if code.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "No record found with the given id"})
		return
	}

	ctx.JSON(http.StatusOK, code)

}

func CreateCode(ctx *gin.Context) {
	var code models.Code
	err := ctx.BindJSON(&code)

	if err != nil {

		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Request data is not valid"})
		return
	}

	DB := models.GetDB()

	DB.Create(&code)
	ctx.JSON(http.StatusOK, code)
}

func DeleteCode(ctx *gin.Context) {

	DB := models.GetDB()

	id := ctx.Params.ByName("id")
	var code models.Code

	DB.First(&code, id)

	if code.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	} else {
		_ = ctx.BindJSON(&code)

		DB.Delete(&code)
		ctx.JSON(http.StatusOK, gin.H{"message": "delete sucess"})
	}

	ctx.JSON(http.StatusOK, code)

}
