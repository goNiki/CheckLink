package links

import (
	"goNiki/CheckLink/internal/services"
	"log/slog"
)

type LinksHandler struct {
	log     *slog.Logger
	checker services.LinksChecker
	report  services.ReportService
}

func NewLinksHandler(log *slog.Logger, checker services.LinksChecker, report services.ReportService) *LinksHandler {
	return &LinksHandler{
		log:     log,
		checker: checker,
		report:  report,
	}
}
