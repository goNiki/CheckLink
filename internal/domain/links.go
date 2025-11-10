package domain

import "time"

type LinkStatus string

const (
	StatusAvailable    LinkStatus = "available"
	StatusNotAvailable LinkStatus = "not available"
)

type Link struct {
	URL       string
	Status    LinkStatus
	CheckedAt time.Time
}
