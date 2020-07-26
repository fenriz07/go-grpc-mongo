package main

import (
	"context"
	"fmt"
	"log"

	"github.com/fenriz07/go-grpc-mongo/client/product"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Go client run")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatal(err.Error())
	}

	defer cc.Close()

	c := product.NewProductServiceClient(cc)

	productModel := &product.Product{
		Name:  "Smarthphone YY",
		Price: 20500.50,
	}

	createProduct, err := c.CreateProduct(context.Background(), &product.CreateProductRequest{
		Product: productModel,
	})

	if err != nil {
		log.Fatalf("Error creando el producto %v", err)
	}

	log.Printf("Producto creado %v \n", createProduct)

}
