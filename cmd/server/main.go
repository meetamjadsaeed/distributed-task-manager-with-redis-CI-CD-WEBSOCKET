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

	r := gin.Default()

	// r.GET("/ws", handlers.HandleConnections)

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	go handlers.BroadcastMessages()

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
