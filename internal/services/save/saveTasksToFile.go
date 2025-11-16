package save

import (
	"context"
	"fmt"
	"goNiki/CheckLink/pkg/errorsAPP"
)

func (s *service) SaveTasksToFile(ctx context.Context) error {
	tasks, err := s.taskstorage.GetAndCleanTasks(ctx)
	if err != nil {
		return fmt.Errorf("%w: %w", errorsAPP.ErrSaveTasksToFile, err)
	}

	err = s.filestorage.SaveTasksToFile(ctx, tasks)
	if err != nil {
		return fmt.Errorf("%w: %w", errorsAPP.ErrSaveTasksToFile, err)
	}
	return nil
}
