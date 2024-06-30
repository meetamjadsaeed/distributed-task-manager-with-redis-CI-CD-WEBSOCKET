package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"time"

	"testing"

	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
	"github.com/meetamjadsaeed/task-manager/internal/models"
	"github.com/meetamjadsaeed/task-manager/internal/services"
	"github.com/meetamjadsaeed/task-manager/pkg/config"
	"github.com/stretchr/testify/assert"
)

// func CreateTask(c *gin.Context) {
// 	var task models.Task
// 	if err := c.ShouldBindJSON(&task); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	if err := services.CreateTask(&task); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, task)
// }

// @Summary Create a task
// @Description Create a new task
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body models.Task true "Task"
// @Success 200 {object} models.Task
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /tasks [post]
func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.CreateTask(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	message := []byte("New task created: " + task.Title)
	broadcast <- message

	c.JSON(http.StatusOK, task)
}

// func GetTasks(c *gin.Context) {
// 	tasks, err := services.GetTasks()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, tasks)
// }

func GetTask(c *gin.Context) {
	id := c.Param("id")
	ctx := context.Background()

	// Try to get the task from Redis
	val, err := config.RDB.Get(ctx, id).Result()
	if err == redis.Nil {
		// If not found, get it from the database
		task, err := services.GetTask(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Cache the task in Redis
		jsonTask, _ := json.Marshal(task)
		config.RDB.Set(ctx, id, jsonTask, 10*time.Minute)

		c.JSON(http.StatusOK, task)
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		var task models.Task
		json.Unmarshal([]byte(val), &task)
		c.JSON(http.StatusOK, task)
	}
}

func UpdateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.UpdateTask(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	if err := services.DeleteTask(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}

func TestCreateTask(t *testing.T) {
	router := gin.Default()
	router.POST("/tasks", CreateTask)

	task := models.Task{
		Title:       "Test Task",
		Description: "This is a test task",
		Status:      "To Do",
	}

	jsonTask, _ := json.Marshal(task)

	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonTask))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var createdTask models.Task
	json.Unmarshal(w.Body.Bytes(), &createdTask)
	assert.Equal(t, task.Title, createdTask.Title)
}

func GenerateReport(c *gin.Context) {
	report, err := services.GenerateReport()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"report": report})
}

// Upgrade connection to a WebSocket
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	Conn *websocket.Conn
	Send chan []byte
}

var clients = make(map[*Client]bool)
var broadcast = make(chan []byte)

// Handle WebSocket connections
func HandleConnections(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Error upgrading connection: %v", err)
		return
	}
	defer conn.Close()

	client := &Client{Conn: conn, Send: make(chan []byte)}
	clients[client] = true

	go handleMessages(client)

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			delete(clients, client)
			break
		}
		broadcast <- message
	}
}

// Handle incoming messages
func handleMessages(client *Client) {
	for {
		message := <-client.Send
		err := client.Conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Printf("Error writing message: %v", err)
			client.Conn.Close()
			delete(clients, client)
			break
		}
	}
}

// Broadcast messages to all clients
func BroadcastMessages() {
	for {
		message := <-broadcast
		for client := range clients {
			select {
			case client.Send <- message:
			default:
				close(client.Send)
				delete(clients, client)
			}
		}
	}
}
