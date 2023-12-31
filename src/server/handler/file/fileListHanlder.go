package handler

import (
	"archive/zip"
	"errors"
	"fmt"
	"github.com/kataras/iris/v12"
	"homeHelper/config"
	"homeHelper/src/server/dto/file"
	"homeHelper/src/server/services"
	"io"
	"io/fs"
	"mime/multipart"
	"os"
	"path/filepath"
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

	dto := file.NewFileList("Files", path, realPath)

	for _, fileInfo := range files {
		dto.AddFile(fileInfo)
	}

	return dto, nil
}

func FileUploadHandler(path string, ctx iris.Context) error {
	realPath, err := services.GetFilePathByUrl(path)
	if err != nil {
		return errors.New("Path is not valid")
	}

	isDir, err := services.IsDir(realPath)
	if !isDir {
		return errors.New("Path is not valid")
	}

	_, _, err = ctx.UploadFormFiles(realPath, beforeSave)
	if err != nil {
		fmt.Println(err)
		return errors.New("Can not upload file")
	}

	return nil
}

func beforeSave(ctx iris.Context, file *multipart.FileHeader) bool {
	path := ctx.Request().FormValue("path")
	realPath, err := services.GetFilePathByUrl(path)
	if err != nil {
		return false
	}

	newFilePath, err := services.GetNewFilePath(file.Filename, realPath)
	if err != nil {
		return false
	}

	file.Filename = filepath.Base(newFilePath)

	return true
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

	isDir, err := services.IsDir(realPath)
	if !isDir {
		return errors.New("Path is not directory")
	}

	newDirPath, err := services.GetNewFilePath(name, realPath)
	if err != nil {
		return err
	}

	if err := os.Mkdir(newDirPath, os.ModePerm); err != nil {
		return errors.New("Can not create folder")
	}

	return nil
}

func FileRenameHandler(path string, name string, oldName string) error {
	err := services.CheckFolderName(name)
	if err != nil {
		return err
	}

	realPath, err := services.GetFilePathByUrl(path)
	oldFilePath := realPath + string(filepath.Separator) + oldName

	if err != nil {
		return errors.New("Path is not valid")
	}

	isExist, err := services.IsFileExists(oldFilePath)
	if err != nil || !isExist {
		return errors.New("Old file not exists")
	}

	newFilePath, err := services.GetNewFilePath(name, realPath)
	if err != nil {
		return err
	}

	err = os.Rename(oldFilePath, newFilePath)
	if err != nil {
		return errors.New("Cannot rename file")
	}

	return nil
}

func FileDeleteHandler(url string, items []string) error {
	realPath, err := services.GetFilePathByUrl(url)

	if err != nil {
		return errors.New("Path is not valid")
	}

	for _, item := range items {
		fullPath := realPath + string(filepath.Separator) + item
		isDir, err := services.IsDir(fullPath)
		if err != nil {
			return errors.New("Failed to delete some files")
		}

		if isDir {
			err := os.RemoveAll(fullPath)
			if err != nil {
				return errors.New("Failed to delete directory: " + item)
			}
		} else {
			e := os.Remove(fullPath)
			if e != nil {
				return errors.New("Failed to delete file: " + item)
			}
		}
	}

	return nil
}

func FileDownloadHandler(url string, items []string) (string, error) {
	realPath, err := services.GetFilePathByUrl(url)

	if err != nil {
		return "", errors.New("Path is not valid")
	}

	if len(items) == 1 {
		fullPath := realPath + string(filepath.Separator) + items[0]
		isDir, err := services.IsDir(fullPath)
		if err != nil {
			return "", errors.New("Failed to add some files to archive")
		}

		if !isDir {
			return realPath + string(filepath.Separator) + items[0], nil
		}
	}

	archive, err := os.Create(config.APP["TEMP_FOLDER"] + string(filepath.Separator) + "archive.zip")
	if err != nil {
		return "", errors.New("Can not create archive")
	}
	defer archive.Close()
	zipWriter := zip.NewWriter(archive)

	for _, item := range items {
		fullPath := realPath + string(filepath.Separator) + item
		zipPath := item
		isDir, err := services.IsDir(fullPath)
		if err != nil {
			return "", errors.New("Failed to add some files to archive")
		}

		if isDir {
			err = filepath.WalkDir(fullPath, func(path string, di fs.DirEntry, err error) error {
				if di.IsDir() {
					return nil
				}
				err = addToArchive(zipWriter, path[len(realPath):], path)

				return nil
			})

			continue
		} else {
			err = addToArchive(zipWriter, zipPath, fullPath)
			if err != nil {
				return "", errors.New("Failed to add some files to archive")
			}
		}
	}

	zipWriter.Close()

	return config.APP["TEMP_FOLDER"] + string(filepath.Separator) + "archive.zip", nil
}

func addToArchive(zw *zip.Writer, zipPath string, filePath string) error {
	w, err := zw.Create(zipPath)
	if err != nil {
		return errors.New("Failed to add some files to archive")
	}

	openedFile, err := os.Open(filePath)
	if err != nil {
		return errors.New("Failed to add some files to archive")
	}
	defer openedFile.Close()

	if _, err = io.Copy(w, openedFile); err != nil {
		return errors.New("Failed to add some files to archive")
	}

	return nil
}

func ShowFileHandler(url string) (*file.FileInfo, error) {
	realPath, err := services.GetFilePathByUrl(url)
	if err != nil {
		return nil, errors.New("File not found")
	}

	fileInfo, err := file.NewFileInfo(url, realPath)
	if err != nil {
		return nil, errors.New("File not found")
	}

	return fileInfo, nil
}
