package model

type Node struct {
	Id    int    `json:"id"`
	Label string `json:"label"`
}

type Edge struct {
	From int `json:"from"`
	To   int `json:"to"`
}
