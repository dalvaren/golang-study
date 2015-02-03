package product

import (
	"errors"
)

type Product struct {
	Id string
	AccountId int
	Value float64
}

func NewProduct (productId string, accountId int) (*Product, error) {
	if len(productId) == 0 {
		return nil, errors.New("Empty App Id")
	}
	var createdProduct = new(Product)
	createdProduct.Id = productId
	createdProduct.AccountId = accountId	
	createdProduct.Value = 0.0
	return createdProduct, nil
}