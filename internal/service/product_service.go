package service

import (
	"github.com/upalx/Store-API/internal/database"
	"github.com/upalx/Store-API/internal/entity"
)

type ProductService struct {
	ProductDB database.ProductDB
}

func NewProductService(productDB database.ProductDB) *ProductService {
	return &ProductService{ProductDB: productDB}
}

func (ps *ProductService) GetProducts() ([]*entity.Product, error) {
	products, err := ps.ProductDB.GetProducts()

	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ps *ProductService) GetProduct(id string) (*entity.Product, error) {
	product, err := ps.ProductDB.GetProduct(id)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (ps *ProductService) GetProductsByCategoryID(categoryID string) ([]*entity.Product, error) {
	products, err := ps.ProductDB.GetProductsByCategoryID(categoryID)

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (ps *ProductService) CreateProduct(name, description, categoryID, imageURL string, price float64) (*entity.Product, error) {
	product := entity.NewProduct(name, description, categoryID, imageURL, price)

	_, err := ps.ProductDB.CreateProduct(product)

	println("IS HERE ON CREATE PRODUCT SERVICE")

	if err != nil {
		println("HAS AN ERROR ON SERVICE")
		return nil, err
	}

	return product, nil
}
