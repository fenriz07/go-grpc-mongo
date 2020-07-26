package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/fenriz07/go-grpc-mongo/server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/*CreateProduct inserta un producto en la base de datos*/
func CreateProduct(p models.ProductModel) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("grpc")
	col := db.Collection("products")

	result, err := col.InsertOne(ctx, p)

	if err != nil {
		return "", status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	ObjID, ok := result.InsertedID.(primitive.ObjectID)

	if !ok {
		return "", status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot convert OID: %v", err),
		)
	}

	return ObjID.Hex(), nil
}
