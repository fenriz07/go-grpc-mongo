package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//ProductModel model colecction
type ProductModel struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name"`
	Price float64            `bson:"price"`
}
