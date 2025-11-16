package filestorage

import (
	"context"
	"encoding/json"
	"fmt"
	"goNiki/CheckLink/internal/domain"
	"goNiki/CheckLink/internal/storage/models"
	"goNiki/CheckLink/pkg/errorsAPP"
	"os"
)

func (s *storage) SaveTasksToFile(_ context.Context, tasks map[string]*domain.Task) error {
	const op = "storage.filestorage.savetaskstofile"

	date := models.StorageTasks{
		Batches: tasks,
	}

	jsondate, err := json.MarshalIndent(date, "", "")
	if err != nil {
		return fmt.Errorf("%s: %w: %v", op, errorsAPP.ErrMarshalIndent, err)
	}

	if err := os.WriteFile(s.filepathTasks, jsondate, 0644); err != nil {
		return fmt.Errorf("%s: %w: %v", op, errorsAPP.ErrWriteFile, err)
	}

	return nil
}
