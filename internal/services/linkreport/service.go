package linkreport

import (
	"goNiki/CheckLink/internal/storage"
)

type service struct {
	linksstorage storage.LinksStorage
}

func NewLinkReportService(linksstorage storage.LinksStorage) *service {
	return &service{
		linksstorage: linksstorage,
	}
}
