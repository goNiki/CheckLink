package models

import "goNiki/CheckLink/internal/domain"

const (
	DateDir     = "date"
	StorageFile = "links_storage.json"
)

type StorageDate struct {
	Batches map[int64]*domain.LinkBatch `json:"batches"`
	LastID  int64                       `json:"last_id"`
}
