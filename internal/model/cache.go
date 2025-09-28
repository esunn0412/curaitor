package model

type CachedFile struct {
	FilePath string `json:"file_path"`
	Content  string `json:"content"`
}
