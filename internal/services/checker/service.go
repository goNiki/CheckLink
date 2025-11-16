package checker

import (
	"goNiki/CheckLink/internal/client"
	"goNiki/CheckLink/internal/storage"
	"log/slog"
)

type service struct {
	client      client.HTTPClient
	linkstorage storage.LinksStorage
	log         *slog.Logger
}

func NewLinksChecker(client client.HTTPClient, linkstorage storage.LinksStorage, log *slog.Logger) *service {
	return &service{
		client:      client,
		linkstorage: linkstorage,
		log:         log,
	}
}
