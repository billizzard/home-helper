package config

import (
	"os"
	"strconv"
	"time"
)

var APP = map[string]string{
	"APP_PORT":          "9000",
	"USER_FILES_FOLDER": "./user/files",
	//"APP_ADDR":     "0.0.0.0:80",
	//"LOG_FILE":     "/var/log/mw/mw.log",
	//"FIREBASE_KEY": "/app/config/keys/firebaseAdminKey.json",
	//"SIGN_KEY":     "_IH@#CODIThdFPnxQSYLvJaZrfZVjz-2342990457434165",
	//"VERSION":      "1.0.0", // сделать дамп перед деплоем
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
