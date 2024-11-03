package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type ProductUseCase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	return ProductUseCase{
		repository: repo,
	}
}

func (pu *ProductUseCase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUseCase) GetProductById(id int) (*model.Product, error) {
	return pu.repository.GetProductById(id)
}

func (pu *ProductUseCase) CreateProduct(product model.Product) (model.Product, error) {
	productId, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	product.ID = productId

	return product, nil
}

func (pu *ProductUseCase) UpdateProduct(id int, product model.Product) (model.Product, error) {
	productId, err := pu.repository.UpdateProduct(id, product)
	if err != nil {
		return model.Product{}, err
	}

	product.ID = productId

	return product, nil
}

func (pu *ProductUseCase) DeleteProduct(id int) error {
	err := pu.repository.DeleteProduct(id)
	if err != nil {
		return err
	}

	return nil
}