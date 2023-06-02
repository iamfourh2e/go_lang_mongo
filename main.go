package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProductModel struct {
	ID    string  `json:"_id" bson:"_id"`
	Name  string  `json:"name" bson:"name"`
	Price float64 `json:"price" bson:"price"`
}

func main() {
	r := gin.Default()
	client, err := loadMongodbConfig()
	if err != nil {
		panic(err.Error())
	}
	//connect to mongodb client
	db := client.Database("test")
	col := db.Collection("products")
	//creata a gateway for get
	//get all products from mongodb with criteria (filter)
	//all produts return from mongodb is a cursor ([element])
	//delcare variable to retrieve that
	r.GET("/products", func(ctx *gin.Context) {
		var filter = bson.M{}
		var products []ProductModel

		cur, err := col.Find(context.Background(), filter)

		if err = cur.All(context.Background(), &products); err != nil {
			ctx.JSON(http.StatusBadRequest, bson.M{
				"message": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, bson.M{
			"data":    products,
			"message": "success",
		})

	})

	r.POST("/product/add", func(ctx *gin.Context) {
		var product ProductModel
		ctx.BindJSON(&product)
		product.ID = primitive.NewObjectID().Hex()
		res, err := col.InsertOne(context.Background(), &product)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, bson.M{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, bson.M{
			"data":    res.InsertedID,
			"message": "success",
		})

	})
	r.Run(":6969")

}

func loadMongodbConfig() (*mongo.Client, error) {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}

	return client, nil
}
