package models

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type ProductModel struct {
	ID    string  `bson:"_id,omitempty" json:"id"`
	Name  string  `json:"name" bson:"name"`
	Price float64 `json:"price" bson:"price"`
}
type ClassProduct struct {
	ProductCollection *mongo.Collection
}

func NewClassProduct(client *mongo.Client) *ClassProduct {
	return &ClassProduct{
		ProductCollection: client.Database("SHOP").Collection("products"),
	}
}

func (p *ClassProduct) CreateProduct(product *ProductModel) (*mongo.InsertOneResult, error) {
	res, err := p.ProductCollection.InsertOne(context.Background(), product)
	return res, err
}
