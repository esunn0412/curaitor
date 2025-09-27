package model

type Question struct {
	Question string `json:"question"`
	Choices []string `json:"choices"`
	Answer int `json:"answer"`
} 
