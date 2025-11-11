package services

import (
	"context"
	"goNiki/CheckLink/internal/domain"
)

type Checker interface {
	CheckLink(ctx context.Context, url string) (domain.Link, error)
	CheckBatch(ctx context.Context, urls []string) (domain.LinkBatch, error)
}
