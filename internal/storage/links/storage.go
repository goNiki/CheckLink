package linksstorage

import (
	"fmt"
	"goNiki/CheckLink/internal/domain"
	repo "goNiki/CheckLink/internal/storage"
	"goNiki/CheckLink/pkg/errorsAPP"
	"sync"
	"sync/atomic"
)

type storage struct {
	LinkBatch   map[int64]*domain.LinkBatch
	counter     atomic.Int64
	mU          sync.RWMutex
	filestorage repo.FileStorage
}

func NewLinksStorage(repo repo.FileStorage) (*storage, error) {
	date, err := repo.LoadLinks()
	if err != nil {
		return &storage{
			LinkBatch:   make(map[int64]*domain.LinkBatch),
			filestorage: repo,
		}, fmt.Errorf("%w: %w", errorsAPP.ErrLoadLinks, err)
	}

	s := &storage{
		LinkBatch:   date.Batches,
		filestorage: repo,
	}
	s.counter.Store(date.LastID)

	return s, nil
}
