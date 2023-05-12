package handlers

import "net/http"

func Ping(w http.ResponseWriter, r *http.Request) {
	data := []byte("pong")
	w.Write(data)
}
