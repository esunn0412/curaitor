package model

type CourseInfo struct {
	Code        string `json:"course_code"`
	CourseTitle string `json:"course_title"`
	Desc        string `json:"desc"`
}

type FileInfo struct {
	CourseInfo
	FileType string `json:"file_type"`
	Title    string `json:"title"`
}
