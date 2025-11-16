package filestorage

import (
	"encoding/json"
	"goNiki/CheckLink/internal/storage/models"
	"os"
)

func (s *storage) LoadLinks() (models.StorageLinks, error) {
	jsondate, err := os.ReadFile(s.filepathLinks)
	if err != nil {
		return models.StorageLinks{}, err
	}
	var result models.StorageLinks
	if err := json.Unmarshal(jsondate, &result); err != nil {
		return models.StorageLinks{}, err
	}

	return result, nil
}
