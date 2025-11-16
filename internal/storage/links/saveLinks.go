package linksstorage

import (
	"context"
	"goNiki/CheckLink/internal/domain"
)

func (s *storage) SaveLinks(_ context.Context, linkBatch *domain.LinkBatch) error {
	s.mU.Lock()
	defer s.mU.Unlock()
	s.LinkBatch[linkBatch.Number] = linkBatch
	return nil
}

func (s *storage) NextID() int64 {
	return s.counter.Add(1)
}
