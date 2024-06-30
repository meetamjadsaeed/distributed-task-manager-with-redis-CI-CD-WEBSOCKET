package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/meetamjadsaeed/task-manager/internal/models"
	"github.com/meetamjadsaeed/task-manager/internal/services"
	"github.com/meetamjadsaeed/task-manager/pkg/config"
	"github.com/stretchr/testify/assert"
)

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
