package linksstorage

import (
	"context"
	"goNiki/CheckLink/internal/domain"
	"goNiki/CheckLink/pkg/errorsAPP"
)

func (s *storage) GetAllLinks(_ context.Context) (map[int64]*domain.LinkBatch, int64, error) {
	if s.LinkBatch == nil {
		return nil, 0, errorsAPP.ErrNoLinks
	}
	lastID := s.counter.Load()

	return s.LinkBatch, lastID, nil

}
