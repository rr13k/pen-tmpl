package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"test2/internal/app/servers"
	"test2/internal/app/servers/user_servers"
)

func GetUserList(w http.ResponseWriter, r *http.Request) {
	type Args struct {
		servers.Pagination
		Name *string
	}
	var (
		args Args
		err  error
	)
	json.NewDecoder(r.Body).Decode(&args)
	defer r.Body.Close()

	var result BaseResponse
	result.Data = make(map[string]interface{})

	userServers := user_servers.User{
		Pagination: args.Pagination,
		Name:       args.Name,
	}

	result.Data, err = userServers.GetAll()
	result.Total, err = userServers.Count()
	if err != nil {
		fmt.Print(err)
	}

	_res, _ := json.Marshal(result)
	w.Write(_res)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	type Args struct {
		Name *string
		Id   *int
	}
	var (
		args Args
		err  error
	)
	json.NewDecoder(r.Body).Decode(&args)

	var result BaseResponse
	result.Data = make(map[string]interface{})

	userServers := user_servers.User{
		Name: args.Name,
		Id:   args.Id,
	}

	result.Data, err = userServers.Get()
	if err != nil {
		fmt.Println(err)
	}

	_res, _ := json.Marshal(result)
	w.Write(_res)
}
