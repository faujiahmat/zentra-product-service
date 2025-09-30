package handler

import (
	"context"

	"github.com/faujiahmat/zentra-product-service/src/interface/service"
	"github.com/faujiahmat/zentra-product-service/src/model/dto"
	pb "github.com/faujiahmat/zentra-proto/protogen/product"
	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ProductGrpcImpl struct {
	productService service.Product
	pb.UnimplementedProductServiceServer
}

func NewProductGrpc(ps service.Product) pb.ProductServiceServer {
	return &ProductGrpcImpl{
		productService: ps,
	}
}

func (p *ProductGrpcImpl) FindManyByIdsForCart(ctx context.Context, data *pb.ProductIds) (*pb.ProductsCartRes, error) {
	res, err := p.productService.FindManyByIds(ctx, data.Ids)
	if err != nil {
		return nil, err
	}

	return &pb.ProductsCartRes{
		Data: res,
	}, nil
}

func (p *ProductGrpcImpl) ReduceStocks(ctx context.Context, data *pb.ReduceStocksReq) (*emptypb.Empty, error) {
	var req []*dto.ReduceStocksReq
	if err := copier.Copy(&req, data.Data); err != nil {
		return nil, err
	}

	err := p.productService.ReduceStocks(ctx, req)
	return nil, err
}

func (p *ProductGrpcImpl) RollbackStocks(ctx context.Context, data *pb.RollbackStocksReq) (*emptypb.Empty, error) {
	var req []*dto.RollbackStoksReq
	if err := copier.Copy(&req, data.Data); err != nil {
		return nil, err
	}

	err := p.productService.RollbackStoks(ctx, req)
	return nil, err
}
