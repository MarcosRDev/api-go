package usecase

import (
	"gin-api/model"
	"gin-api/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()

}

func (pu *ProductUsecase) CreatProduct(Product model.Product) (model.Product, error) {
	ProductID, err := pu.repository.CreatProduct(Product)
	if err != nil {
		return model.Product{}, err
	}

	Product.ID = ProductID
	return Product, nil
}

func (pu *ProductUsecase) GetProductByID(id int) (*model.Product, error) {

	product, err := pu.repository.GetProductByID(id)

	if err != nil {
		return nil, err

	}

	return product, nil
}
