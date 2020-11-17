package models

import (
	"errors"
)

var (
	// ErrEmptyProductID is used to indicate that the id is empty.
	ErrEmptyProductID = errors.New("product 'id' cannot be empty")

	// ErrEmptyProductName is used to indicate that the name is empty.
	ErrEmptyProductName = errors.New("product 'name' cannot be empty")

	// ErrEmptyProductPrice is used to indicate that the price is empty.
	ErrEmptyProductPrice = errors.New("product 'price' cannot be empty")
)

// Product Representation of products
type Product struct {
	ID    string   `json:"id,omitempty"`
	Name  string   `json:"name,omitempty"`
	Price int      `json:"price,omitempty"`
	DType []string `json:"dgraph.type,omitempty"`
}

// NewProduct Constructor For a Product
func NewProduct(id string, name string, price int) (*Product, error) {
	if id == "" {
		return nil, ErrEmptyProductID
	}

	if name == "" {
		return nil, ErrEmptyProductName
	}

	if price <= 0 {
		return nil, ErrEmptyProductPrice
	}

	newProduct := &Product{
		ID:    id,
		Name:  name,
		Price: price,
		DType: []string{"Product"},
	}

	return newProduct, nil
}
