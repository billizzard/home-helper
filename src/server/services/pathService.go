package services

import (
	"errors"
	"homeHelper/config"
	"os"
	"strings"
)

func checkFilePathInUrl(url string) bool {
	if !strings.HasPrefix(url, "root") {
		return false
	}

	return true
}

func GetFilePathByUrl(url string) (string, error) {
	if !checkFilePathInUrl(url) {
		return "", errors.New("Path not found")
	}
	return config.APP["USER_FILES_FOLDER"] + url[4:], nil
}

func IsDir(url string) (bool, error) {
	path, err := GetFilePathByUrl(url)
	if err != nil {
		return false, err
	}

	f, err := os.Stat(path)
	if err != nil {
		return false, errors.New("Wrong path")
	}

	return f.IsDir(), nil
}

func CheckFolderName(name string) error {
	if len(name) > 100 {
		return errors.New("Name must be less than 100 characters")
	}

	return nil
}

func SanitizeName(name string) string {
	return name
}

func GetPrevPath(path string) string {
	splitted := strings.Split(path, "/")

	if len(splitted) < 2 {
		return ""
	}

	return strings.Join(splitted[0:len(splitted)-1], "/")
}

func GetA() {

}
