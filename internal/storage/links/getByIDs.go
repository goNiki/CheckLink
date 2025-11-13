package linksstorage

import (
	"context"
	"goNiki/CheckLink/internal/domain"
)

func (s *storage) GetByIDs(_ context.Context, ids []int64) ([]domain.LinkBatch, error) {
	s.mU.RLock()
	defer s.mU.RUnlock()

	var result []domain.LinkBatch

	for _, id := range ids {
		if batch, exist := s.LinkBatch[id]; exist {
			result = append(result, *batch)
		}
	}
	return result, nil
}
