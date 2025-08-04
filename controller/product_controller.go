package controller

import (
	"net/http"
	"products-api/model"
	"products-api/usecase"

	"github.com/gin-gonic/gin"
)

type productController struct {
	// Usecase
	productUseCase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		productUseCase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products := []model.Product{
		{
			ID:    1,
			Name:  "Batata frita",
			Price: 20,
		},
	}

	ctx.JSON(http.StatusOK, products)
}
