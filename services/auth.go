package services

import (
	"io"
	"net/http"
	"task-manager-plus-gateway/utils"

	"github.com/gin-gonic/gin"
)

func registerUser(ctx *gin.Context) {
	jsonData, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	outputData, statusCode := utils.SendRequest(utils.Post, utils.USERS_BACKEND, "/auth/register", "", jsonData)
	ctx.JSON(statusCode, outputData)
}

func login(ctx *gin.Context) {
	jsonData, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	outputData, statusCode := utils.SendRequest(utils.Post, utils.USERS_BACKEND, "/auth/login", "", jsonData)
	ctx.JSON(statusCode, outputData)
}

func RegisterAuthRoutes(rg *gin.RouterGroup) {
	rg.POST("/register", registerUser)
	rg.POST("/login", login)
}
