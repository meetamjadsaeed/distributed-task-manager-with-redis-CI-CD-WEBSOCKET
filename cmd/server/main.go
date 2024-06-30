package main

import (
	"github.com/gin-gonic/gin"
	"github.com/meetamjadsaeed/task-manager/internal/config"
	"github.com/meetamjadsaeed/task-manager/internal/handlers"
	"github.com/meetamjadsaeed/task-manager/pkg/config"
)

func main() {
	config.LoadEnv()
	config.InitDB()
	config.InitRedis()

	r := gin.Default()

	// User Authentication
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	// Task Management
	r.POST("/tasks", handlers.CreateTask)
	r.GET("/tasks/:id", handlers.GetTask)
	r.PUT("/tasks/:id", handlers.UpdateTask)
	r.DELETE("/tasks/:id", handlers.DeleteTask)

	// Real-time Notifications
	r.GET("/ws", handlers.HandleConnections)

	// Report Generation
	r.GET("/report", handlers.GenerateReport)

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	go handlers.BroadcastMessages()

	r.Run(":8080")
}
