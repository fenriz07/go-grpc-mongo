package handler

import (
	"context"

	"github.com/fenriz07/go-grpc-mongo/server/bd"
	"github.com/fenriz07/go-grpc-mongo/server/models"
	"github.com/fenriz07/go-grpc-mongo/server/product"
)

type Server struct {
	product.UnimplementedProductServiceServer
}

func (*Server) CreateProduct(ctx context.Context, req *product.CreateProductRequest) (*product.CreateProductResponse, error) {

	//Parsear contenido
	prod := req.GetProduct()

	data := models.ProductModel{
		Name:  prod.GetName(),
		Price: prod.GetPrice(),
	}

	id, err := bd.CreateProduct(data)

	if err != nil {
		return nil, err
	}

	return &product.CreateProductResponse{
		Product: &product.Product{
			Id:    id,
			Name:  prod.GetName(),
			Price: prod.GetPrice(),
		},
	}, nil

}
