package dto

type ReqCheckLink struct {
	Links []string `json:"links"`
}

type ResponseCheckLink struct {
	Links    map[string]string `json:"links"`
	LinksNum int64             `json:"links_num"`
}

type ReqGetReportLinks struct {
	LinksList []int64 `json:"links_list"`
}

type Task struct {
	ID     string `json:"id"`
	Date   string `json:"date"`
	Path   string `json:"path"`
	Method string `json:"method"`
}
