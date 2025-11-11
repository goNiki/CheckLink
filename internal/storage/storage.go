package storage

import (
	"goNiki/CheckLink/internal/domain"
)

type LinksStorage interface {
	SaveDate(linkBatch *domain.LinkBatch) error
	NextID() int64
}
