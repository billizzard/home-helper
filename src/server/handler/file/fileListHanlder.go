package handler

import (
	"errors"
	"homeHelper/src/server/dto/file"
	"homeHelper/src/server/services"
	"log"
	"os"
)

func FileListHandler(path string) (*file.FileList, error) {
	if !services.CheckFilePathInUrl(path) {
		return nil, errors.New("Path not found")
	}

	// проверить будет ли идти на уровень выше если передать ..
	files, err := os.ReadDir(services.GetFilePathByUrl(path))
	if err != nil {
		log.Fatal(err)
	}

	dto := file.NewFileList("Files", path)

	for _, fileInfo := range files {
		dto.AddFile(fileInfo)
	}

	return dto, nil
}