package models

type Note struct {
	ID      uint64 `json:"id"`
	Content string `json:"name"`
	Author  string `json:"author"`
}
