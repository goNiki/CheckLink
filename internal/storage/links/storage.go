package linksstorage

import (
	"goNiki/CheckLink/internal/domain"
	"sync"
	"sync/atomic"
)

type storage struct {
	LinkBatch map[int64]*domain.LinkBatch
	counter   atomic.Int64
	mU        sync.RWMutex
}

func NewStorage() *storage {
	return &storage{
		LinkBatch: make(map[int64]*domain.LinkBatch),
	}
}

func (s *storage) SaveDate(linkBatch *domain.LinkBatch) error {
	s.mU.Lock()
	defer s.mU.Unlock()
	s.LinkBatch[linkBatch.Number] = linkBatch
	return nil
}

func (s *storage) NextID() int64 {
	return s.counter.Add(1)
}
