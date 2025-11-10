package dto

type Req struct {
	Links []string `json:"links"`
}

type Response struct {
	Links    map[string]string `json:"links"`
	LinksNum int64             `json:"links_num"`
}
