package filestorage

import (
	"encoding/json"
	"goNiki/CheckLink/internal/storage/models"
	"os"
)

const (
	DateDir     = "date"
	StorageFile = "links_storage.json"
)

type storage struct {
	filepath string
}

func NewFileStorage() *storage {
	return &storage{
		filepath: DateDir + "/" + StorageFile,
	}
}

func (s *storage) Save(date models.StorageDate) error {
	jsondate, err := json.MarshalIndent(date, "", "")
	if err != nil {
		return err
	}

	os.WriteFile(s.filepath, jsondate, 0644)
	return nil

}

func (s *storage) LoadDate() (models.StorageDate, error) {
	jsondate, err := os.ReadFile(s.filepath)
	if err != nil {
		return models.StorageDate{}, err
	}
	var result models.StorageDate
	if err := json.Unmarshal(jsondate, &result); err != nil {
		return models.StorageDate{}, err
	}

	return result, nil
}
