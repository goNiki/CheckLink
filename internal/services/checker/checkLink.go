package checker

import (
	"context"
	"goNiki/CheckLink/internal/domain"
	"net/http"
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

	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return link, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return link, err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		link.Status = domain.StatusAvailable
		return link, nil
	}
	return link, err

}
