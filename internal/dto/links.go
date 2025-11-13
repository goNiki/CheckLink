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
