package httpclient

import (
	"bytes"
	"context"
	"goNiki/CheckLink/internal/dto"
	"net/http"
)

func (c *client) TaskWorker(ctx context.Context, task dto.Task) error {

	url := c.baseURL + task.Path

	req, err := http.NewRequestWithContext(ctx, task.Method, url, bytes.NewBufferString(task.Date))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	_, err = c.client.Do(req)
	if err != nil {
		return err
	}

	return nil

}
