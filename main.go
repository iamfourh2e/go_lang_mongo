package main

import (
	"context"

	"gomongo/models"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	route := gin.Default()
	client, err := loadMongoConfig()
	if err != nil {
		log.Fatal(err)
	}

	pImpl := models.NewClassProduct(client)
	//Classic rest POST, GET, PUT, DELETE
	route.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	route.GET("/products", func(c *gin.Context) {
		p := &models.ProductModel{
			ID:    primitive.NewObjectID().Hex(),
			Name:  "Khe",
			Price: 4352222,
		}
		if err != nil {

			log.Fatal(err.Error())
		}
		res, err := pImpl.CreateProduct(p)
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{

			"message": res,
		})

	})
	route.Run(":6969")

}

func loadMongoConfig() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	return client, nil
}
