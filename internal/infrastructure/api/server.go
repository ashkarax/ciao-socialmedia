package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ashkarax/ciao-socialmedia/internal/config"
	"github.com/ashkarax/ciao-socialmedia/internal/infrastructure/handler"
	"github.com/ashkarax/ciao-socialmedia/internal/infrastructure/routes"
	"github.com/gin-gonic/gin"
)

type ServerHttp struct {
	engin  *gin.Engine
	config *config.PortManager
}

func NewServerHttp(apikey *config.ApiKey, config *config.PortManager, userHandler *handler.UserHandler) *ServerHttp {

	engin := gin.Default()

	engin.Use(func(c *gin.Context) {
		apiKey := c.GetHeader("x-api-Key")
		if apiKey != apikey.Key {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API key"})
			c.Abort()
			return
		}
		c.Next()
	})

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
