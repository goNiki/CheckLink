package save

import (
	"context"
	"fmt"
	"goNiki/CheckLink/pkg/errorsAPP"
)

func (s *service) SaveLinksToFile(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	links, lastID, err := s.linkstorage.GetAllLinks(ctx)
	if err != nil {
		return fmt.Errorf("%w: %w", errorsAPP.ErrSaveLinksToFile, err)
	}

	err = s.filestorage.SaveLinksToFile(ctx, links, lastID)
	if err != nil {
		return fmt.Errorf("%w, %w", errorsAPP.ErrSaveLinksToFile, err)
	}

	return nil
}
