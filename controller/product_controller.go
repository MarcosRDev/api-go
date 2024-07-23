package controller

import (
	"gin-api/model"
	"gin-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productController struct {
	ProductUsecase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		ProductUsecase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {

	products, err := p.ProductUsecase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)

}

func (p *productController) CreatProduct(ctx *gin.Context) {
	var product model.Product
	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := p.ProductUsecase.CreatProduct(product)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}

func (p *productController) GetProductByID(ctx *gin.Context) {

	id := ctx.Param("id")

	if id == "" {
		respose := model.Response{
			Message: "ID do produto não pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, respose)
		return
	}

	ProductId, err := strconv.Atoi(id)
	if err != nil {
		respose := model.Response{
			Message: "ID do produto precisa ser um numero",
		}
		ctx.JSON(http.StatusBadRequest, respose)
		return
	}

	product, err := p.ProductUsecase.GetProductByID(ProductId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		respose := model.Response{
			Message: "produto não foi encontrado",
		}
		ctx.JSON(http.StatusNotFound, respose)
		return
	}

	ctx.JSON(http.StatusOK, product)

}
