package services

import (
	"homeHelper/config"
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
