package main

import (
	"user/config"
	"user/controller"
	"user/repository"
)

func main() {
	config.Init()
	repository.Init()
	controller.Run(config.Configs.User.HttpPort)
}
