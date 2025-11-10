package links

import (
	"goNiki/CheckLink/internal/services"
	"log/slog"
)

type handler struct {
	log         *slog.Logger
	linkchecker services.Linkchecker
}

func NewLinksHandler(log *slog.Logger, linkchecker services.Linkchecker) *handler {
	return &handler{
		log:         log,
		linkchecker: linkchecker,
	}
}
