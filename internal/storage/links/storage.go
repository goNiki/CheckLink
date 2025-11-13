package linksstorage

import (
	"goNiki/CheckLink/internal/domain"
	repo "goNiki/CheckLink/internal/storage"
	"sync"
	"sync/atomic"
)

type storage struct {
	LinkBatch   map[int64]*domain.LinkBatch
	counter     atomic.Int64
	mU          sync.RWMutex
	filestorage repo.FileStorage
}

func NewStorage(repo repo.FileStorage) *storage {
	date, err := repo.LoadDate()
	if err != nil {
		return &storage{
			LinkBatch:   make(map[int64]*domain.LinkBatch),
			filestorage: repo,
		}
	}
	var counter atomic.Int64
	counter.Store(date.LastID)

	return &storage{
		LinkBatch:   date.Batches,
		counter:     counter,
		filestorage: repo,
	}
}
