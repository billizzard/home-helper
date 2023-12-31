package services

import (
	"errors"
	"homeHelper/config"
	"homeHelper/src/server/services/sugar"
	url2 "net/url"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
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

	urlEsc, err := url2.QueryUnescape(url[4:])
	if err != nil {
		return "", errors.New("Path not found")
	}
	path := filepath.Clean(filepath.FromSlash(urlEsc))
	if path == "." {
		path = ""
	}

	res, err := filepath.EvalSymlinks(config.APP["USER_FILES_FOLDER"] + path)
	if err != nil {
		return "", errors.New("Path not found")
	}

	return res, nil
}

func IsDir(path string) (bool, error) {
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

func sanitizeName(name string) string {
	name = regexp.MustCompile(`[^\p{L}0-9 \.\-_\)\(\[\]\|]+`).ReplaceAllString(name, "")
	name = strings.ReplaceAll(name, " ", "_")

	return name
}

func GetPrevPath(path string) string {
	splitted := strings.Split(path, string(filepath.Separator))

	if len(splitted) < 2 {
		return ""
	}

	return strings.Join(splitted[0:len(splitted)-1], string(filepath.Separator))
}

func IsFileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, nil
	}

	if err != nil {
		return true, errors.New("Error checking directory existence")
	}

	return true, nil
}

func GetNewFilePath(name string, pathIn string) (string, error) {
	sanitizedName := sanitizeName(name)
	i := 0

	ext := filepath.Ext(sanitizedName)
	ps := string(filepath.Separator)
	fileName := sanitizedName[:len(sanitizedName)-len(ext)]
	fullPath := pathIn + ps + sanitizedName

	for {
		i++
		isExist, err := IsFileExists(fullPath)
		if err != nil {
			return "", err
		}

		if isExist {
			if i == 1 {
				fullPath = pathIn + ps + fileName + "(1)" + ext
			} else {
				pervPostfix := "(" + strconv.Itoa(i-1) + ")"
				fullPath = pathIn + ps + fileName[0:len(fileName)-len(pervPostfix)] + "(" + strconv.Itoa(i) + ")" + ext
			}
		} else {
			break
		}
	}

	return fullPath, nil
}

func GetPublicLink(filePath string) string {
	if !sugar.InArray(strings.ToLower(filepath.Ext(filePath)), []string{".jpg", ".jpeg", ".png"}) {
		return ""
	}

	return strings.Replace(filePath, config.APP["USER_FILES_FOLDER"], "public", 1)
}
