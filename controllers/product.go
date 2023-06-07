package controllers

import (
	"gomongo/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductControllerImpl struct {
	db           *mongo.Database
	productModel *models.ProductModelImpl
	client       *mongo.Client
}

func NewProductController(db *mongo.Database, client *mongo.Client) *ProductControllerImpl {
	productModel := models.NewProductModelImpl(db, client)

	return &ProductControllerImpl{
		db:           db,
		productModel: productModel,
		client:       client,
	}

}

func (p *ProductControllerImpl) GetAllProducts(ctx *gin.Context) {
		p.productModel.GetAllProducts()
}
