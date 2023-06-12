package routes

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// product model
type ProductModelImpl struct {
	db         *mongo.Database
	collection *mongo.Collection
}
type ProductModel struct {
	ID    string  `bson:"_id,omitempty" json:"id"`
	Name  string  `json:"name" bson:"name"`
	Price float64 `json:"price" bson:"price"`
}

func (p *ProductModelImpl) FindProducts() ([]ProductModel, error) {

	cursor, err := p.collection.Find(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	//declare variable to store all products
	var products []ProductModel
	//loop through cursor and append each element to products
	for cursor.Next(context.Background()) {
		var product ProductModel
		cursor.Decode(&product)
		products = append(products, product)
	}
	//return products
	return products, nil
}

func NewProductModelImpl(db *mongo.Database) *ProductModelImpl {
	return &ProductModelImpl{
		db:         db,
		collection: db.Collection("product"),
	}

}

// controller
type ProductControllerImpl struct {
	ProductModelImpl *ProductModelImpl
}

func (p *ProductControllerImpl) GetProduct(ctx *gin.Context) {
	//get all products from mongodb with criteria (filter)
	//all produts return from mongodb is a cursor ([element])
	//delcare variable to retrieve that
	pro, err := p.ProductModelImpl.FindProducts()

}

func ProductRoutes(route *gin.Engine, db *mongo.Database, client *mongo.Client) {
	ProductControllerImpl := ProductControllerImpl{
		ProductModelImpl: NewProductModelImpl(db),
	}
	r := route.Group("/products")
	r.GET("/", ProductControllerImpl.GetProduct)

}
