package tasks

import (
	"context"
	"goNiki/CheckLink/internal/domain"
	"goNiki/CheckLink/pkg/errorsAPP"
)

func (s *storage) GetAndCleanTasks(_ context.Context) (map[string]*domain.Task, error) {
	s.mU.Lock()
	defer s.mU.Unlock()

	if s.TaskBatch == nil {
		return nil, errorsAPP.ErrNoTasks
	}

	tasks := s.TaskBatch

	s.TaskBatch = make(map[string]*domain.Task)

	return tasks, nil
}
