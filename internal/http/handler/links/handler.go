package links

import (
	"goNiki/CheckLink/internal/services"
	"log/slog"
)

type handler struct {
	log     *slog.Logger
	checker services.Checker
	report  services.ReportService
}

func NewLinksHandler(log *slog.Logger, checker services.Checker, report services.ReportService) *handler {
	return &handler{
		log:     log,
		checker: checker,
		report:  report,
	}
}
