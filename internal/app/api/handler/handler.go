package handler

type BaseResponse struct {
	Success bool
	Code    int
	Message string
	Data    interface{}
}
