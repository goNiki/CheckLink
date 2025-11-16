package httpclient

import (
	"net/http"
	"time"
)

type client struct {
	client  *http.Client
	baseURL string
}

func NewHttpClient(baseUrl string) *client {
	return &client{
		client: &http.Client{
			Timeout: 10 * time.Second,
			Transport: &http.Transport{
				TLSHandshakeTimeout:   5 * time.Second,
				ResponseHeaderTimeout: 5 * time.Second,
				IdleConnTimeout:       30 * time.Second,
			},
		},
		baseURL: baseUrl,
	}
}
