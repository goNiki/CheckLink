package services

import "context"

type Linkchecker interface {
	LinkCheck(ctx context.Context, url string) (string, error)
}
