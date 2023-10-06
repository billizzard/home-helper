package file

import (
	"homeHelper/src/server/services"
)

type FileInfo struct {
	Title    string
	Path     string
	PrevPath string
}

func NewFileInfo(title string, path string) *FileInfo {
	model := FileInfo{}
	model.Title = title
	model.Path = path
	model.PrevPath = services.GetPrevPath(path)

	return &model
}
