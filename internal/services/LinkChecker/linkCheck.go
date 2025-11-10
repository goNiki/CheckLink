package linkchecker

import (
	"context"
	"net/http"
	"strings"
	"time"
)

func (s *service) LinkCheck(ctx context.Context, url string) (string, error) {

	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "HEAD", url, nil)
	if err != nil {
		return "not available", err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return "not available", err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		return "available", nil
	}
	return "not available", err

}
