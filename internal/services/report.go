package services

import (
	"github.com/meetamjadsaeed/task-manager/internal/models"
)

func GenerateReport() (interface{}, error) {
	// Implement your report generation logic here
	// For example, count the number of tasks by status
	var countByStatus []struct {
		Status string
		Count  int
	}
	if result := models.DB.Table("tasks").Select("status, count(*) as count").Group("status").Scan(&countByStatus); result.Error != nil {
		return nil, result.Error
	}

	return countByStatus, nil
}
