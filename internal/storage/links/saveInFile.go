package linksstorage

import "goNiki/CheckLink/internal/storage/models"

func (s *storage) SaveInFile() error {
	date := models.StorageDate{
		Batches: s.LinkBatch,
		LastID:  s.counter.Load(),
	}

	err := s.filestorage.Save(date)
	if err != nil {
		return err
	}
	return nil
}
