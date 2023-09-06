package services

import (
	"io"
	"net/http"
	"task-manager-plus-gateway/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func registerUser(ctx *gin.Context) {
	jsonData, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	outputDataChan := make(chan map[string]string)
	statusCodeChan := make(chan int)

	go func() {
		outputData, statusCode := utils.SendRequest(utils.Post, utils.USERS_BACKEND+"/auth/register", jsonData)
		outputDataChan <- outputData
		statusCodeChan <- statusCode
	}()

	select {
	case outputData := <-outputDataChan:
		statusCode := <-statusCodeChan
		ctx.JSON(statusCode, outputData)
	case <-time.After(5 * time.Second):
		ctx.JSON(http.StatusGatewayTimeout, gin.H{"message": "request timeout"})
	}
}

func login(ctx *gin.Context) {
	jsonData, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	outputDataChan := make(chan map[string]string)
	statusCodeChan := make(chan int)

	go func() {
		outputData, statusCode := utils.SendRequest(utils.Post, utils.USERS_BACKEND+"/auth/login", jsonData)
		outputDataChan <- outputData
		statusCodeChan <- statusCode
	}()

	select {
	case outputData := <-outputDataChan:
		statusCode := <-statusCodeChan
		ctx.JSON(statusCode, outputData)
	case <-time.After(5 * time.Second):
		ctx.JSON(http.StatusGatewayTimeout, gin.H{"message": "request timeout"})
	}
}

func RegisterAuthRoutes(rg *gin.RouterGroup) {
	rg.POST("/register", registerUser)
	rg.POST("/login", login)
}
