package domain

type Payload struct {
	Id       string `json:"id"`
	ToId     string `json:"to_id"`
	SendTime string `json:"send_time"`
	Msg      string `json:"msg"`
}
