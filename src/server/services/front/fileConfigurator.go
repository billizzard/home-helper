package front

import (
	"errors"
	"fmt"
	"homeHelper/src/server/services/http"
	"os"
)

func CreateConfigFile() error {
	addr, err := getBackendAddress()
	err = fillFile(addr)
	if err != nil {
		return err
	}
	return nil
}

func fillFile(backendUrl string) error {
	d1 := []byte(fmt.Sprintf("window.appConfig = {\n    backendUrl: \"%s\"\n}", backendUrl))
	err := os.WriteFile("src/public/assets/js/config.js", d1, 0766)
	if err != nil {
		return err
	}

	return nil
}

func getBackendAddress() (string, error) {
	addresses, err := http.FindPublicAddresses()
	if len(addresses) == 0 {
		return "", errors.New("Can not found local routes")
	}
	if err != nil {
		return "", err
	}

	return addresses[0], nil
}
