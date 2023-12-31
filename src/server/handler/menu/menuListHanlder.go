package handler

import (
	"homeHelper/src/server/dto/menuPage"
)

func MenuListHandler() (*menuPage.MenuPage, error) {
	dto := menuPage.NewMenuPage()

	return dto, nil
}
