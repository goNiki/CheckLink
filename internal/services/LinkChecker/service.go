package linkchecker

import "net/http"

type service struct {
	client *http.Client
}

func NewLinkChecker(client *http.Client) *service {
	return &service{
		client: client,
	}
}
