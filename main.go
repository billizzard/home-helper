package main

import (
	"fmt"
	"homeHelper/config"
	"homeHelper/src/server"
)

func main() {
	err := server.RunApi(config.APP["APP_PORT"])
	if err != nil {
		fmt.Println("error", err)
	}
}
