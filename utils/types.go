package utils

type BasicResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Config struct {
	Port    int    `json:"PORT"`
	Secret  string `json:"SECRET"`
	Logging bool   `json:"LOGGING"`
}

type Meta struct {
	Total int64 `json:"total"`
	Links Links `json:"links"`
}

type Links struct {
	Next string `json:"next"`
	Prev string `json:"prev"`
}
