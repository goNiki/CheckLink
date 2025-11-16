package filestorage

import (
	"encoding/json"
	"goNiki/CheckLink/internal/storage/models"
	"os"
)

func (s *storage) LoadTask() (models.StorageTasks, error) {
	tasksByte, err := os.ReadFile(s.filepathTasks)
	if err != nil {
		return models.StorageTasks{}, err
	}

	var tasks models.StorageTasks

	if err := json.Unmarshal(tasksByte, &tasks); err != nil {
		return models.StorageTasks{}, err
	}

	return tasks, nil

}
