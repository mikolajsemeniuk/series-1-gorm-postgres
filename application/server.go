package application

import (
	"fmt"
	"gorm-postgres/controllers"
	"gorm-postgres/settings"

	"github.com/gin-gonic/gin"
)

type Server interface {
	Listen()
}

type server struct {
	configuration settings.Configuration
	router        *gin.Engine
	// TODO: think what to do with multiple controllers
	order controllers.Order
}

func (server *server) Listen() {
	v1 := server.router.Group(server.configuration.GetServerBasePath())
	{
		order := v1.Group("order")
		{
			order.GET("", server.order.List)
			order.GET(":id", server.order.Find)
			order.POST("", server.order.Add)
			order.PATCH(":id", server.order.Update)
			order.DELETE(":id", server.order.Remove)
		}
	}

	port := fmt.Sprintf(":%s", server.configuration.GetServerPort())
	server.router.Run(port)
}

func New() (Server, error) {
	configuration, err := settings.NewConfiguration()
	if err != nil {
		return nil, err
	}

	server := &server{
		configuration: configuration,
		router:        gin.Default(),
		order:         controllers.NewOrder(),
	}

	return server, err
}
