package model

type Task struct {
	Id     uint   `json:"id"`
	Title  string `json:"title"`
	Active bool   `json:"active"`
}
