package main

import (
	"test2/internal/app/api"
	"test2/internal/app/common"
	"test2/internal/app/toolkit/log"
	"fmt"
	"net/http"
)

func main() {
	log.Info("test2 start...", "localhost"+common.PORT)
	api.RegisterApi()
	err := http.ListenAndServe(common.PORT, nil)
	if err != nil {
		fmt.Println(err)
	}
}
