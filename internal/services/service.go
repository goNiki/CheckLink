package services

import (
	"context"
	"goNiki/CheckLink/internal/domain"
	"io"
)

type Checker interface {
	CheckLink(ctx context.Context, url string) (domain.Link, error)
	CheckBatch(ctx context.Context, urls []string) (domain.LinkBatch, error)
}

type ReportService interface {
	CreateReport(ctx context.Context, numbers []int64) (io.Reader, error)
}
