package controller

import (
	"github.com/kataras/iris/v12"
	"homeHelper/src/server/dto"
)

func BookList(ctx iris.Context) {
	books := []dto.Book{
		{"Mastering Concurrency in Go"},
		{"Go Design Patterns"},
		{"Black Hat Go"},
	}

	ctx.JSON(books)
	// TIP: negotiate the response between server's prioritizes
	// and client's requirements, instead of ctx.JSON:
	// ctx.Negotiation().JSON().MsgPack().Protobuf()
	// ctx.Negotiate(books)
}
