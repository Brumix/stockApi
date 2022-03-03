package main

import (
	"fmt"
	"os"

	_ "stockApi/config"
	"stockApi/server"
	_ "stockApi/services"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.Info("The app is starting ...")
}

func main() {
	port := os.Getenv("PORT")
	log.Info("Router started at port ", port)
	port = fmt.Sprintf(":%v", port)
	_ = server.New().Run(port)
}
