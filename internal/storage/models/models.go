package models

import "goNiki/CheckLink/internal/domain"

const (
	DateDir          = "date"
	StorageLinksFile = "links_storage.json"
	StorageTasksFile = "pending_tasks.json"
)

type StorageLinks struct {
	Batches map[int64]*domain.LinkBatch `json:"batches"`
	LastID  int64                       `json:"last_id"`
}

type StorageTasks struct {
	Batches map[string]*domain.Task `json:"batches"`
}
