package controller

import (
	"errors"
	"github.com/kataras/iris/v12"
	handler "homeHelper/src/server/handler/menu"
	"homeHelper/src/server/services/http"
)

func MenuList(ctx iris.Context) {
	dto, err := handler.MenuListHandler()

	if err != nil {
		http.HttpBadRequest(ctx, errors.New("Can not get addresses"))

		return
	}

	ctx.JSON(dto)
}
