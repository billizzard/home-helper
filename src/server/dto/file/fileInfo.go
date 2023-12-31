package file

import (
	"errors"
	"homeHelper/src/server/services"
	"homeHelper/src/server/services/converter"
	"os"
)

type FileInfo struct {
	Title      string
	Path       string
	PrevPath   string
	realPath   string
	PublicLink string
	Size       string
	Icon       string
	IsFileInfo bool
}

func NewFileInfo(path string, realPath string) (*FileInfo, error) {
	info, err := os.Stat(realPath)

	if err != nil {
		return nil, errors.New("File not found")
	}

	model := FileInfo{}
	model.Title = info.Name()
	model.Path = path
	model.Size = converter.ByteToHuman(info.Size())
	model.Icon = converter.FileToIcon(info)
	model.IsFileInfo = true
	model.realPath = realPath
	model.PublicLink = model.getPublicLink()
	model.PrevPath = services.GetPrevPath(path)

	return &model, nil
}

func (fl *FileInfo) getPublicLink() string {
	return services.GetPublicLink(fl.realPath)
}
