package httpclient

import (
	"context"
	"fmt"
	"goNiki/CheckLink/pkg/errorsAPP"
	"net/http"
)

func (c *client) CheckLink(ctx context.Context, url string) (int, error) {

	statusCode, err := c.doRequest(ctx, http.MethodHead, url)
	if err != nil {
		return 0, err
	}

	if c.shouldFallbackToGet(statusCode) {
		return c.doRequest(ctx, http.MethodGet, url)
	}

	return statusCode, nil

}

func (c *client) doRequest(ctx context.Context, method string, url string) (int, error) {
	const op = "client.http.dorequest"

	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return 0, fmt.Errorf("%s: %w: %v", op, errorsAPP.ErrFailedCreateRequest, err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")

	resp, err := c.client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("%s : %w: %v", op, errorsAPP.ErrFailedRequest, err)
	}

	defer resp.Body.Close()

	return resp.StatusCode, nil
}

func (c *client) shouldFallbackToGet(statusCode int) bool {
	switch statusCode {
	case http.StatusMethodNotAllowed, http.StatusNotImplemented:
		return true
	default:
		return false
	}
}
