package services

import (
	"io"
	"net/http"
	"task-manager-plus-gateway/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func createTask(ctx *gin.Context) {
	jsonData, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	outputDataChan := make(chan map[string]interface{})
	statusCodeChan := make(chan int)

	go func() {
		outputData, statusCode := utils.SendRequest(utils.Post, utils.TASKS_BACKEND+"/users/tasks/create", jsonData)
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

func getWorkspaceTasks(ctx *gin.Context) {
	taskId := ctx.Param("taskId")

	jsonData, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	outputDataChan := make(chan map[string]interface{})
	statusCodeChan := make(chan int)

	go func() {
		outputData, statusCode := utils.SendRequest(utils.Get, utils.TASKS_BACKEND+"/users/tasks/get/"+taskId, jsonData)
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

func deleteTask(ctx *gin.Context) {
	taskId := ctx.Param("taskId")

	jsonData, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	outputDataChan := make(chan map[string]interface{})
	statusCodeChan := make(chan int)

	go func() {
		outputData, statusCode := utils.SendRequest(utils.Delete, utils.TASKS_BACKEND+"/users/tasks/delete/"+taskId, jsonData)
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

func RegisterTasksRoutes(rg *gin.RouterGroup) {
	rg.POST("/create", createTask)
	rg.GET("/get/:taskId", getWorkspaceTasks)
	rg.DELETE("/delete/:taskId", deleteTask)
}
