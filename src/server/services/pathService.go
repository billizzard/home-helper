package services

import (
	"errors"
	"homeHelper/config"
	"os"
	"strings"
)

func CheckFilePathInUrl(url string) bool {
	if !strings.HasPrefix(url, "root") {
		return false
	}

	return true
}

func GetFilePathByUrl(url string) string {
	return config.APP["USER_FILES_FOLDER"] + url[4:]
}

func IsDir(url string) (bool, error) {
	path := GetFilePathByUrl(url)
	f, err := os.Stat(path)
	if err != nil {
		return false, errors.New("Wrong path")
	}

	return f.IsDir(), nil
}
