package util

type Response struct {
	Data    interface{} `json:"data,omitempty"`
	Total   int64       `json:"total,omitempty"`
	Message string      `json:"message,omitempty"`
	Status  string      `json:"status,omitempty"`
}
