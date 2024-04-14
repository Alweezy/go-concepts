package main

import (
	"encoding/json"
	"fmt"
	"go-refresher/json/web"
	"log"
	"os"
)

type Config struct {
	ServerPort string `json:"server_port"`
}

func main() {
	currentDir, err := os.Getwd()
	filePath := fmt.Sprintf(currentDir + "/json/config.json")
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatalf(err.Error())
	}

	config := new(Config)

	err = json.NewDecoder(file).Decode(config)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("server listening on  port %s ...", config.ServerPort)
	web.RunWeb(config.ServerPort)
}
