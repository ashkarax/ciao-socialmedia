package server

import (
	"fmt"
	"log"

	"github.com/ashkarax/ciao-socialmedia/internal/config"
	"github.com/ashkarax/ciao-socialmedia/internal/infrastructure/handler"
	"github.com/ashkarax/ciao-socialmedia/internal/infrastructure/routes"
	"github.com/gin-gonic/gin"
)

type ServerHttp struct {
	engin  *gin.Engine
	config *config.PortManager
}

func NewServerHttp(config *config.PortManager, userHandler *handler.UserHandler) *ServerHttp {

	engin := gin.Default()

	//routes.AdminRoutes(engin.Group("/admin"), adminHandler)
	routes.UserRoutes(engin.Group(""), userHandler)

	return &ServerHttp{engin: engin, config: config}

}

func (server *ServerHttp) Start() {
	port_with_colon := ":" + server.config.RunnerPort
	err := server.engin.Run(port_with_colon)
	if err != nil {
		log.Fatal("gin engin couldn't start")
	}
	fmt.Printf("\ngin engin start:%s", server.config.RunnerPort)
}
