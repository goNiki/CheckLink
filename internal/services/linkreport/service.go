package linkreport

import (
	"goNiki/CheckLink/internal/storage"
)

type service struct {
	linksstorage storage.LinksStorage
}

func NewReportService(linksstorage storage.LinksStorage) *service {
	return &service{
		linksstorage: linksstorage,
	}
}
