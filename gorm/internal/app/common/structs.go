package common

import "test2/internal/app/toolkit/log"

type Config struct {
	Mysql  *MysqlConfig                `yaml:"mysql"`
	Mode   string                      `yaml:"mode"`
	Logs   map[string]log.LoggerConfig `yaml:"logs"`
	Server ServerType                  `yaml:"server"`
}

type ServerType struct {
	Port int `yaml:"port"`
}

type MysqlConfig struct {
	User     string `yaml:"handler"`
	Password string `yaml:"password"`
	Port     int    `yaml:"port"`
	Db       string `yaml:"db"`
	Host     string `yaml:"host"`
	Charset  string `yaml:"charset"`
}

type BaseResponse struct {
	Msg   string      `json:"msg"`
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Total int         `json:"total"`
}

type RunReport struct {
	Success int    `json:"success"`
	Error   int    `json:"error"`
	Date    string `json:"date"`
}

type SessionUser struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Role     int    `json:"role"`
	Email    string `json:"email"`
	RealName string `json:"real_name"`
}

type AddCaseMap struct {
	Total int    `json:"total"`
	Name  string `json:"name"`
}
