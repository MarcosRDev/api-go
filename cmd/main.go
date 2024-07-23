package main

import (
	"gin-api/controller"
	"gin-api/db"
	"gin-api/repository"
	"gin-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	// camada repository
	ProductRepository := repository.NewProductRepository(dbConnection)
	// canada de useCase
	ProductUsecase := usecase.NewProductUsecase(ProductRepository)
	// camada de controllers
	productController := controller.NewProductController(ProductUsecase)

	LoginRepository := repository.NewLoginRepository(dbConnection)

	LoginUsecase := usecase.NewLoginUsecase(LoginRepository)

	loginController := controller.NewLoginController(LoginUsecase)

	server.GET("/ping", func(ctx *gin.Context) {

		ctx.JSON(200, gin.H{
			"message": "pong",
		})

	})

	server.GET("/products", productController.GetProducts)

	server.GET("/product/:id", productController.GetProductByID)

	server.POST("product", productController.CreatProduct)

	server.POST("login", loginController.LoginUser)

	server.Run(":8000")

}
