package checker

import (
	"goNiki/CheckLink/internal/storage"
	"net/http"
)

type service struct {
	client      *http.Client
	linkstorage storage.LinksStorage
}

func NewChecker(client *http.Client, linkstorage storage.LinksStorage) *service {
	return &service{
		client:      client,
		linkstorage: linkstorage,
	}
}
