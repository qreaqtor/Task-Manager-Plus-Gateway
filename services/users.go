package services

import (
	"io"
	"net/http"
	"task-manager-plus-gateway/utils"

	"github.com/gin-gonic/gin"
)

func getUser(ctx *gin.Context) {
	userId, ok := ctx.Params.Get("id")
	if !ok {
		userId = ctx.MustGet("userId").(string)
	}

	jsonData, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	outputData, statusCode := utils.SendRequest(utils.Get, utils.USERS_BACKEND, "/users/get/", userId, jsonData)
	ctx.JSON(statusCode, outputData)
}

func updateUser(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(string)

	jsonData, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	outputData, statusCode := utils.SendRequest(utils.Patch, utils.USERS_BACKEND, "/users/update/", userId, jsonData)
	ctx.JSON(statusCode, outputData)
}

func deleteUser(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(string)

	jsonData, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	outputData, statusCode := utils.SendRequest(utils.Delete, utils.USERS_BACKEND, "/users/delete/", userId, jsonData)
	ctx.JSON(statusCode, outputData)
}

func RegisterUsersRoutes(rg *gin.RouterGroup) {
	rg.GET("/get/:id", getUser)
	rg.GET("/get/me", getUser)
	rg.PATCH("/update", updateUser)
	rg.DELETE("/delete", deleteUser)
}
