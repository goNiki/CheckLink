package links

import (
	"log/slog"
)

type handler struct {
	log *slog.Logger
}

func NewLinksHandler(log *slog.Logger) *handler {
	return &handler{
		log: log,
	}
}
