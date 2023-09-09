package main

import (
	"log"
	"task-manager-plus-gateway/services"
	"task-manager-plus-gateway/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	authPath := server.Group("/auth")
	services.RegisterAuthRoutes(authPath)

	usersPath := server.Group("/users")
	usersPath.Use(utils.JwtAuthMiddleware())
	services.RegisterUsersRoutes(usersPath)

	tasksPath := server.Group("/tasks")
	tasksPath.Use(utils.JwtAuthMiddleware())
	services.RegisterTasksRoutes(tasksPath)

	workspacePath := server.Group("/workspaces")
	workspacePath.Use(utils.JwtAuthMiddleware())
	services.RegisterWorkspaceRoutes(workspacePath)

	log.Fatal(server.Run(":8080"))
}
