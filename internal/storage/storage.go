package storage

import (
	"context"
	"goNiki/CheckLink/internal/domain"
	"goNiki/CheckLink/internal/storage/models"
)

type LinksStorage interface {
	SaveDate(ctx context.Context, linkBatch *domain.LinkBatch) error
	NextID() int64
	SaveInFile() error
	GetByIDs(ctx context.Context, ids []int64) ([]domain.LinkBatch, error)
}

type FileStorage interface {
	Save(date models.StorageDate) error
	LoadDate() (models.StorageDate, error)
}
