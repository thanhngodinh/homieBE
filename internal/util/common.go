package util

type Response struct {
	Data    interface{} `json:"data,omitempty"`
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Status  string      `json:"status,omitempty"`
}

type Pagination struct {
	Total    int64 `json:"total"`
	PageIdx  int   `json:"pageIdx"`
	PageSize int   `json:"pageSize"`
}
