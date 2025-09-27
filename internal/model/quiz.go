package model

type QuizInfo struct {
	Code      string     `json:"course_code"`
	// NumFiles  int        `json:"numFiles"`
	Questions []Question `json:"questions"`
}
