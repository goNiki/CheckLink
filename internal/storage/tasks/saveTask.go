package tasks

import (
	"context"
	"goNiki/CheckLink/internal/domain"
)

func (s *storage) SaveDate(_ context.Context, task *domain.Task) error {

	s.mU.Lock()
	defer s.mU.Unlock()

	s.TaskBatch[task.ID] = task

	return nil
}
