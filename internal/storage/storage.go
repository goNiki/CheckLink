package storage

import (
	"context"
	"goNiki/CheckLink/internal/domain"
	"goNiki/CheckLink/internal/storage/models"
)

type LinksStorage interface {
	SaveLinks(ctx context.Context, linkBatch *domain.LinkBatch) error
	NextID() int64
	GetAllLinks(ctx context.Context) (map[int64]*domain.LinkBatch, int64, error)
	GetByIDs(ctx context.Context, ids []int64) ([]domain.LinkBatch, error)
}

type FileStorage interface {
	SaveTasksToFile(_ context.Context, tasks map[string]*domain.Task) error
	SaveLinksToFile(_ context.Context, links map[int64]*domain.LinkBatch, lastID int64) error
	LoadLinks() (models.StorageLinks, error)
	LoadTask() (models.StorageTasks, error)
}

type TasksStorage interface {
	SaveDate(ctx context.Context, task *domain.Task) error
	GetAndCleanTasks(ctx context.Context) (map[string]*domain.Task, error)
}
