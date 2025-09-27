package model

type QuizInfo struct {
	Name string `json:name`
	Code string `json:"code"`
	NumFiles int `json:"numFiles"`
	QaPairs []Question
}