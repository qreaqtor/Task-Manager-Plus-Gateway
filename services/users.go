package services

import (
	"io"
	"net/http"
	"task-manager-plus-gateway/utils"
	"time"

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

	outputDataChan := make(chan map[string]string)
	statusCodeChan := make(chan int)

	go func() {
		outputData, statusCode := utils.SendRequest(utils.Get, utils.USERS_BACKEND+"/users/get/"+userId, jsonData)
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

func updateUser(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(string)

	jsonData, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	outputDataChan := make(chan map[string]string)
	statusCodeChan := make(chan int)

	go func() {
		outputData, statusCode := utils.SendRequest(utils.Patch, utils.USERS_BACKEND+"/users/update/"+userId, jsonData)
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

func deleteUser(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(string)

	jsonData, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	outputDataChan := make(chan map[string]string)
	statusCodeChan := make(chan int)

	go func() {
		outputData, statusCode := utils.SendRequest(utils.Delete, utils.USERS_BACKEND+"/users/delete/"+userId, jsonData)
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

func RegisterUsersRoutes(rg *gin.RouterGroup) {
	rg.GET("/get/:id", getUser)
	rg.GET("/get/me", getUser)
	rg.PATCH("/update", updateUser)
	rg.DELETE("/delete", deleteUser)
}
