package config

import (
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var APP = map[string]string{
	"APP_PORT":          "9030",
	"USER_FILES_FOLDER": filepath.FromSlash("user/files"),
	"TEMP_FOLDER":       filepath.FromSlash("user/temp"),
}

var DB = map[string]string{
	"USER": os.Getenv("MYSQL_USER"),
	"PASS": os.Getenv("MYSQL_PASS"),
	"HOST": os.Getenv("MYSQL_HOST"),
	"PORT": os.Getenv("MYSQL_PORT"),
	"NAME": os.Getenv("MYSQL_DBNAME"),
}

var JWT = map[string]string{
	"SECRET": os.Getenv("JWT_SECRET"),
	"EXP":    strconv.FormatInt(time.Now().Add(time.Minute*10).Unix(), 10),
	"AUD":    "mw.com",
}
