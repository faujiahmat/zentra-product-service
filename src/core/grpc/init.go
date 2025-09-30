package grpc

import (
	"github.com/faujiahmat/zentra-product-service/src/core/grpc/handler"
	"github.com/faujiahmat/zentra-product-service/src/core/grpc/interceptor"
	"github.com/faujiahmat/zentra-product-service/src/core/grpc/server"
	"github.com/faujiahmat/zentra-product-service/src/infrastructure/config"
	"github.com/faujiahmat/zentra-product-service/src/interface/service"
)

func InitServer(ps service.Product) *server.Grpc {
	productHandler := handler.NewProductGrpc(ps)
	unaryResponseInterceptor := interceptor.NewUnaryResponse()

	grpcServer := server.NewGrpc(config.Conf.CurrentApp.GrpcPort, productHandler, unaryResponseInterceptor)
	return grpcServer
}
