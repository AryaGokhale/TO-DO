package models

type Note struct {
	ID      uint64 `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}
