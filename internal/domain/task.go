package domain

type Task struct {
	ID     string `json:"id"`
	Date   string `json:"date"`
	Path   string `json:"path"`
	Method string `json:"method"`
}
