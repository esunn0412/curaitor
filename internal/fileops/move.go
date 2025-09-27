package fileops

import (
	"os"
	"path/filepath"
)

// logic for moving file to another location
// e.g. func Move(src string, dest string) error

func MoveFile(srcFile string, destPath string, destFileName string) error {
	destFile := filepath.Join(destPath, destFileName)
	return os.Rename(srcFile, destFile)
}
