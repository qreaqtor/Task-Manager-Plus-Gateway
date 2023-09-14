package services

import (
	"fmt"
	"io"
	"net/http"
	"task-manager-plus-gateway/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func createWorkspace(ctx *gin.Context) {
	jsonData, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	outputDataChan := make(chan map[string]interface{})
	statusCodeChan := make(chan int)

	go func() {
		outputData, statusCode := utils.SendRequest(utils.Post, utils.TASKS_BACKEND+"/users/workspaces/create", jsonData)
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

func getUserWorkspaces(ctx *gin.Context) {
	username := ctx.MustGet("username").(string)

	jsonData, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	outputDataChan := make(chan map[string]interface{})
	statusCodeChan := make(chan int)

	go func() {
		outputData, statusCode := utils.SendRequest(utils.Get, utils.TASKS_BACKEND+"/users/workspaces/get/all/"+username, jsonData)
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

func deleteWorkspace(ctx *gin.Context) {
	workspaceId := ctx.Param("workspaceId")

	jsonData, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	outputDataChan := make(chan map[string]interface{})
	statusCodeChan := make(chan int)

	go func() {
		outputData, statusCode := utils.SendRequest(utils.Delete, utils.TASKS_BACKEND+"/users/workspaces/delete/"+workspaceId, jsonData)
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

func addUserToWorkspace(ctx *gin.Context) {
	workspaceId := ctx.Param("workspaceId")

	jsonData, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	outputDataChan := make(chan map[string]interface{})
	statusCodeChan := make(chan int)

	go func() {
		outputData, statusCode := utils.SendRequest(utils.Patch, utils.TASKS_BACKEND+"/users/workspaces/add/"+workspaceId, jsonData)
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

func deleteUserFromWorkspace(ctx *gin.Context) {
	workspaceId := ctx.Param("workspaceId")
	username := ctx.MustGet("username").(string)

	jsonData, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	outputDataChan := make(chan map[string]interface{})
	statusCodeChan := make(chan int)

	go func() {
		outputData, statusCode := utils.SendRequest(utils.Patch, utils.TASKS_BACKEND+fmt.Sprintf("/users/workspaces/remove/%s/%s", username, workspaceId), jsonData)
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

func RegisterWorkspaceRoutes(rg *gin.RouterGroup) {
	rg.POST("/create", createWorkspace)
	rg.GET("/get/all", getUserWorkspaces)
	rg.DELETE("/remove/:workspaceId", deleteUserFromWorkspace)
	rg.PATCH("/delete/:workspaceId", deleteWorkspace)
	rg.PATCH("/add/:workspaceId", addUserToWorkspace)
}
