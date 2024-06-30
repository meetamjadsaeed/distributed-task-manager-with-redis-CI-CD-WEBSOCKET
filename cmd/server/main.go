package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/meetamjadsaeed/task-manager/docs"
	"github.com/meetamjadsaeed/task-manager/internal/config"
	"github.com/meetamjadsaeed/task-manager/internal/handlers"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func main() {
	config.LoadEnv()
	config.InitRedis()
	config.InitDB()

	r := gin.Default()

	// r.GET("/ws", handlers.HandleConnections)

	// User Authentication
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	// Task Management
	r.POST("/tasks", handlers.CreateTask)
	r.GET("/tasks/:id", handlers.GetTask)
	r.PUT("/tasks/:id", handlers.UpdateTask)
	r.DELETE("/tasks/:id", handlers.DeleteTask)

	// Real-time Notifications
	r.GET("/ws", handlers.HandleWebSocket)

	// Report Generation
	r.GET("/report", handlers.GenerateReport)

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	go handlers.BroadcastMessages()

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
