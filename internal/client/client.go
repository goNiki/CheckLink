package client

import (
	"context"
	"goNiki/CheckLink/internal/dto"
)

type HTTPClient interface {
	TaskWorker(ctx context.Context, task dto.Task) error
	CheckLink(ctx context.Context, url string) (int, error)
}
