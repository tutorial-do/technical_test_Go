package models

import (
	"errors"
)

var (
	// ErrEmptyBuyerID is used to indicate that the id is empty.
	ErrEmptyBuyerID = errors.New("buyer 'id' cannot be empty")

	// ErrEmptyBuyerName is used to indicate that the name is empty.
	ErrEmptyBuyerName = errors.New("buyer 'name' cannot be empty")

	// ErrEmptyBuyerAge is used to indicate that the age is empty or negative.
	ErrEmptyBuyerAge = errors.New("buyer 'age' cannot be empty or negative")
)

// Buyer Representation of buyer
type Buyer struct {
	ID    string   `json:"id,omitempty"`
	Name  string   `json:"name,omitempty"`
	Age   int      `json:"age,omitempty"`
	DType []string `json:"dgraph.type,omitempty"`
}

// NewBuyer Constructor For a Buyer
func NewBuyer(id string, name string, age int) (*Buyer, error) {
	if id == "" {
		return nil, ErrEmptyBuyerID
	}

	if name == "" {
		return nil, ErrEmptyBuyerName
	}

	if age <= 0 {
		return nil, ErrEmptyBuyerAge
	}

	newBuyer := &Buyer{
		ID:    id,
		Name:  name,
		Age:   age,
		DType: []string{"Buyer"},
	}

	return newBuyer, nil
}
