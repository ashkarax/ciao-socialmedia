package main

import (
	"log"

	"github.com/ashkarax/ciao-socialmedia/internal/config"
	"github.com/ashkarax/ciao-socialmedia/internal/di"
)

func main() {

	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("error at loading the env file using viper")
	}
	server, diErr := di.InitializeAPI(config)
	if diErr != nil {
		log.Fatal("error for server creation")
	}
	server.Start()
}
