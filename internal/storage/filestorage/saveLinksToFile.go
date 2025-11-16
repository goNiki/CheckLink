package filestorage

import (
	"context"
	"encoding/json"
	"fmt"
	"goNiki/CheckLink/internal/domain"
	"goNiki/CheckLink/internal/storage/models"
	"goNiki/CheckLink/pkg/errorsAPP"
	"os"
)

func (s *storage) SaveLinksToFile(_ context.Context, links map[int64]*domain.LinkBatch, lastID int64) error {
	const op = "storage.filestorage.savelinkstofile"
	date := models.StorageLinks{
		Batches: links,
		LastID:  lastID,
	}

	jsondate, err := json.MarshalIndent(date, "", "")
	if err != nil {
		return fmt.Errorf("%s: %w: %v", op, errorsAPP.ErrMarshalIndent, err)
	}

	err = os.WriteFile(s.filepathLinks, jsondate, 0644)
	if err != nil {
		return fmt.Errorf("%s: %w: %v", op, errorsAPP.ErrWriteFile, err)

	}

	return nil

}
