package api

import (
	"test2/internal/app/api/handler"
	"test2/internal/app/toolkit/router"
)

func RegisterApi() {

	router.UrlGroup("/user",
		router.Url("/list", handler.GetUserList),
		router.Url("/find", handler.GetUser),
	)
	router.Init()
}
