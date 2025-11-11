package domain

type LinkStatus string

const (
	StatusAvailable    LinkStatus = "available"
	StatusNotAvailable LinkStatus = "not available"
)

type Link struct {
	URL    string
	Status LinkStatus
}

type LinkBatch struct {
	Links  map[string]LinkStatus
	Number int64
}

func (l *LinkStatus) String() string {
	return string(*l)
}
