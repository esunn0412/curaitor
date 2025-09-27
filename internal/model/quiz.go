package model

type QuizInfo struct {
	Code      string     `json:"code"`
	NumFiles  int        `json:"numFiles"`
	Questions []Question `json:"questions"`
}
