package main

import (
	"cloudCar/internal/app/api"
	"cloudCar/internal/app/common"
	"cloudCar/internal/app/toolkit/log"
	"fmt"
	"net/http"
)

func main() {
	log.Info("cloudCar start...", "localhost"+common.PORT)
	api.RegisterApi()
	err := http.ListenAndServe(common.PORT, nil)
	if err != nil {
		fmt.Println(err)
	}
}
