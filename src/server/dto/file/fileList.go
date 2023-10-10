package file

import (
	"homeHelper/src/server/services"
	"homeHelper/src/server/services/converter"
	"net/url"
	"os"
	"path/filepath"
)

type FileList struct {
	Title      string
	Path       string
	PrevPath   string
	realPath   string
	FilesCount int
	Folders    []File
	Files      []File
}

type File struct {
	Name       string
	IsDir      bool
	Icon       string
	Link       string
	PublicLink string
}

func NewFileList(title string, path string, realPath string) *FileList {
	model := FileList{}
	model.Title = title
	model.Path = path
	model.realPath = realPath
	model.FilesCount = 0
	model.Files = []File{}
	model.Folders = []File{}
	model.PrevPath = services.GetPrevPath(path)

	return &model
}

func (fl *FileList) AddFile(file os.DirEntry) {
	publicLink := ""

	if !file.IsDir() {
		publicLink = fl.getPublicLink(file.Name())
	}

	info, _ := file.Info()

	f := File{
		Name:       file.Name(),
		IsDir:      file.IsDir(),
		Icon:       converter.FileToIcon(info),
		Link:       fl.Path + string(filepath.Separator) + url.QueryEscape(file.Name()),
		PublicLink: publicLink,
	}

	fl.FilesCount++

	if file.IsDir() {
		fl.Folders = append(fl.Folders, f)
	} else {
		fl.Files = append(fl.Files, f)
	}
}

func (fl *FileList) getPublicLink(fileName string) string {
	return services.GetPublicLink(fl.realPath + string(filepath.Separator) + fileName)
	//if !sugar.InArray(strings.ToLower(filepath.Ext(fileName)), []string{".jpg", ".jpeg", ".png"}) {
	//	return ""
	//}
	//
	//return strings.Replace(fl.realPath, config.APP["USER_FILES_FOLDER"], "public", 1) + "/" + fileName
}
