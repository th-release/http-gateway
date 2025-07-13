package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"cth.release/http-gateway/api"
	"cth.release/http-gateway/utils"
)

func main() {
	port := flag.Int("PORT", 8080, "PORT Integer Default: 8080")
	secret := flag.String("SECRET", "XOR", "SECRET String Default: XOR")
	logging := flag.Bool("LOGGING", true, "LOGGING Boolean Default: true")

	flag.Parse()

	err := utils.SetConfig(&utils.Config{
		Port:    *port,
		Secret:  *secret,
		Logging: *logging,
	})

	if err != nil {
		log.Fatalf("Config 저장중 에러 발생 %s", err.Error())
	}

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

	app.App.Listen(fmt.Sprintf(":%d", config.Port))
}
