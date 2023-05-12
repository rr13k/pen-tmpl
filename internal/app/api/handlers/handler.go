package handlers

import "net/http"

type BaseResponse struct {
	Success bool
	Code    int
	Message string
	Data    interface{}
	Total   int `json:"total"`
}

func Ping(w http.ResponseWriter, r *http.Request) {
	data := []byte("pong")
	w.Write(data)
}
