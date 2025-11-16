package checker

import (
	"context"
	"fmt"
	"goNiki/CheckLink/internal/domain"
	"goNiki/CheckLink/pkg/errorsAPP"
	"strings"
	"time"
)

func (s *service) CheckLink(ctx context.Context, url string) (domain.Link, error) {

	link := domain.Link{
		URL:    url,
		Status: domain.StatusNotAvailable,
	}

	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	status, err := s.client.CheckLink(ctx, url)
	if err != nil {
		return link, fmt.Errorf("%s: %w: %w", link.URL, errorsAPP.ErrCheckLink, err)
	}

	if status >= 200 && status < 400 {
		link.Status = domain.StatusAvailable
		return link, nil
	}
	return link, fmt.Errorf("%s: %w: %v", link.URL, errorsAPP.ErrHTTPStatusInvalid, status)

}
