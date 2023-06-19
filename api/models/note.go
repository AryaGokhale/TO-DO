package models

type Note struct {
	ID      uint32 `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}
