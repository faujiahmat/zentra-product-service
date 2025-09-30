package restful

import (
	"github.com/faujiahmat/zentra-product-service/src/core/restful/client"
	"github.com/faujiahmat/zentra-product-service/src/core/restful/delivery"
	"github.com/faujiahmat/zentra-product-service/src/core/restful/handler"
	"github.com/faujiahmat/zentra-product-service/src/core/restful/middleware"
	"github.com/faujiahmat/zentra-product-service/src/core/restful/server"
	"github.com/faujiahmat/zentra-product-service/src/interface/service"
)

func InitServer(ps service.Product, rc *client.Restful) *server.Restful {
	productHandler := handler.NewProductRESTful(ps, rc)

	middleware := middleware.New(rc)
	restfulServer := server.New(productHandler, middleware)

	return restfulServer
}

func InitClient() *client.Restful {
	imageKitDelivery := delivery.NewImageKit()
	restfulClient := client.NewRestful(imageKitDelivery)

	return restfulClient
}
