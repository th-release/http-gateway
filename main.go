package main

import (
	"fmt"
	"log"
	"time"

	"cth.release/http-gateway/api"
	"cth.release/http-gateway/utils"
)

func main() {
	config := utils.GetConfig()

	if config == nil {
		log.Fatalln("Not Found Config")
		return
	}

	app := api.InitServer(config)

	if app == nil {
		log.Fatalln("Init Server Error")
		return
	}

	log.Println(time.Now())

	app.App.Listen(fmt.Sprintf(":%s", config.Port))
}
