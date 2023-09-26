package file

import (
	"net/url"
	"os"
	"strings"
)

type FileList struct {
	Title      string
	Path       string
	PrevPath   string
	FilesCount int
	Folders    []File
	Files      []File
}

type File struct {
	Name  string
	IsDir bool
	Icon  string
	Link  string
}

func NewFileList(title string, path string) *FileList {
	model := FileList{}
	model.Title = title
	model.Path = path
	model.FilesCount = 0
	model.PrevPath = model.getPrevPath(path)

	return &model
}

func (fl *FileList) AddFile(file os.DirEntry) {
	f := File{
		Name:  file.Name(),
		IsDir: file.IsDir(),
		Icon:  fl.getIconByName(file),
		Link:  fl.Path + "/" + url.QueryEscape(file.Name()),
	}

	fl.FilesCount++

	if file.IsDir() {
		fl.Folders = append(fl.Folders, f)
	} else {
		fl.Files = append(fl.Files, f)
	}
}

func (fl *FileList) getIconByName(file os.DirEntry) string {
	if file.IsDir() {
		return "folder"
	}

	return "file-earmark"
}

func (fl *FileList) getPrevPath(path string) string {
	splitted := strings.Split(path, "/")

	if len(splitted) < 2 {
		return ""
	}

	return strings.Join(splitted[0:len(splitted)-1], "/")
}
