package models

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductModel struct {
	ID    string  `bson:"_id,omitempty" json:"id"`
	Name  string  `json:"name" bson:"name"`
	Price float64 `json:"price" bson:"price"`
}
type ProductModelImpl struct {
	ProductCollection *mongo.Collection
}

func NewProductModelImpl(db *mongo.Database, client *mongo.Client) *ProductModelImpl {
	return &ProductModelImpl{
		ProductCollection: db.Collection("products"),
	}
}

func (p *ProductModelImpl) GetAllProducts() {

}
