package routes

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

//product model

type ProductModel struct {
	ID    string  `bson:"_id,omitempty" json:"id"`
	Name  string  `json:"name" bson:"name"`
	Price float64 `json:"price" bson:"price"`
}

// controller
type ProductControllerImpl struct {
	collection *mongo.Collection
}

func (p *ProductControllerImpl) GetProduct(ctx *gin.Context) {
	//get all products from mongodb with criteria (filter)
	//all produts return from mongodb is a cursor ([element])
	//delcare variable to retrieve that
	cursor, err := p.collection.Find(ctx, nil)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	//declare variable to store all products
	var products []ProductModel
	//loop through cursor and append each element to products
	for cursor.Next(ctx) {
		var product ProductModel
		cursor.Decode(&product)
		products = append(products, product)
	}
	//return products
	ctx.JSON(200, products)

}

func ProductRoutes(route *gin.Engine, db *mongo.Database, client *mongo.Client) {
	productObj := ProductControllerImpl{
		collection: db.Collection("products"),
	}
	r := route.Group("/products")
	r.GET("/", productObj.GetProduct)

}
