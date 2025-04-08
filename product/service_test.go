package product_test

import (
	"dummyjson/product"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddProductSuccess(t *testing.T) {
	mock := &product.MockProductService{
		AddProductFunc: func(p product.Product) (product.Product, error) {
			p.Title = "Mocked Product"
			return p, nil
		},
	}

	input := product.Product{
		Title:       "Test",
		Description: "Desc",
		Price:       10,
		Brand:       "Brand",
		Category:    "Category",
	}

	output, err := mock.AddProduct(input)

	assert.NoError(t, err)
	assert.Equal(t, "Mocked Product", output.Title)
}

func TestAddProductFailure(t *testing.T) {
	mock := &product.MockProductService{
		AddProductFunc: func(p product.Product) (product.Product, error) {
			return product.Product{}, errors.New("mock failure")
		},
	}

	_, err := mock.AddProduct(product.Product{})
	assert.Error(t, err)
	assert.EqualError(t, err, "mock failure")
}
