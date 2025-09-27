package model

type CourseInfo struct {
	Code string `json:"code"`
	Desc string `json:"desc"`
}

type FileInfo struct {
	CourseInfo
	FileType   string `json:"file_type"`
	Title      string `json:"title"`
}
