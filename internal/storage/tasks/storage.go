package tasks

import (
	"fmt"
	"goNiki/CheckLink/internal/domain"
	repo "goNiki/CheckLink/internal/storage"
	"goNiki/CheckLink/pkg/errorsAPP"
	"sync"
)

type storage struct {
	TaskBatch   map[string]*domain.Task
	mU          sync.RWMutex
	filestorage repo.FileStorage
}

func NewTaskStorage(filestorage repo.FileStorage) (*storage, error) {
	tasks, err := filestorage.LoadTask()
	if err != nil {
		return &storage{
			TaskBatch:   make(map[string]*domain.Task),
			filestorage: filestorage}, fmt.Errorf("%w: %w", errorsAPP.ErrLoadTasks, err)
	}

	return &storage{
		TaskBatch:   tasks.Batches,
		filestorage: filestorage,
	}, nil
}
