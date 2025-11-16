package services

import (
	"context"
	"goNiki/CheckLink/internal/domain"
	"io"
)

type LinksChecker interface {
	CheckLink(ctx context.Context, url string) (domain.Link, error)
	CheckBatch(ctx context.Context, urls []string) (domain.LinkBatch, error)
}

type ReportService interface {
	CreateReport(ctx context.Context, numbers []int64) (io.Reader, error)
}

type SaveService interface {
	SaveLinksToFile(ctx context.Context) error
	SaveTasksToFile(ctx context.Context) error
}

type TaskService interface {
	SaveTask(ctx context.Context, task domain.Task) error
	ProcessPendingTasks(ctx context.Context) error
}
