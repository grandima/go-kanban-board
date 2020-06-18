package entity

type Comment struct {
	Id int `json:"id"`
	Text string `json:"text"`
	Timestamp int `json:"timestamp"`
}
