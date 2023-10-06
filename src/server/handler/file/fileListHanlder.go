package handler

import (
	"errors"
	"fmt"
	"github.com/kataras/iris/v12"
	"homeHelper/src/server/dto/file"
	"homeHelper/src/server/services"
	"os"
)

func FileListHandler(path string) (*file.FileList, error) {
	realPath, err := services.GetFilePathByUrl(path)
	if err != nil {
		return nil, errors.New("Path not found")
	}

	// проверить будет ли идти на уровень выше если передать ..
	files, err := os.ReadDir(realPath)
	if err != nil {
		// todo: log
		return nil, errors.New("Cannot read directory")
	}

	dto := file.NewFileList("Files", path)

	for _, fileInfo := range files {
		dto.AddFile(fileInfo)
	}

	return dto, nil
}

func FileUploadHandler(path string, ctx *iris.Context) error {
	realPath, err := services.GetFilePathByUrl(path)
	if err != nil {
		return errors.New("Path is not valid")
	}

	isDir, err := services.IsDir(path)
	if !isDir {
		return errors.New("Path is not valid")
	}

	_, _, err = (*ctx).UploadFormFiles(realPath)
	if err != nil {
		return errors.New("Can not upload file")
	}

	return nil
}

func FolderCreateHandler(path string, name string) error {
	err := services.CheckFolderName(name)
	if err != nil {
		return err
	}

	realPath, err := services.GetFilePathByUrl(path)

	if err != nil {
		return errors.New("Path is not valid")
	}

	isDir, err := services.IsDir(path)
	if !isDir {
		return errors.New("Path is not directory")
	}

	// todo: sanitize and check name
	name = services.SanitizeName(name)

	if err := os.Mkdir(realPath+"/"+name, os.ModePerm); err != nil {
		return errors.New("Can not create folder")
	}

	return nil
}

func ShowFileHandler(path string) *file.FileInfo {
	fileInfo := file.NewFileInfo(path, path)
	fmt.Println("SHOW FILE HANDLER--------")
	fmt.Println(path)
	fmt.Println(fileInfo)

	return fileInfo
}
