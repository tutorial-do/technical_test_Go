package models

import (
	"technical_test_Go/backend/models"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewProduct(t *testing.T) {
	c := require.New(t)
	product, err := models.NewProduct("PRD1", "Noodles", 3045)
	c.NotEmpty(product.ID)
	c.NotEmpty(product.Name)
	c.Equal("Noodles", product.Name)
	c.Equal(3045, product.Price)
	c.NoError(err)
}

func TestNewProductErrors(t *testing.T) {
	c := require.New(t)

	_, err := models.NewProduct("", "Noodles", 3045)
	c.Equal(models.ErrEmptyProductID, err)
	_, err = models.NewProduct("PRD1", "", 3045)
	c.Equal(models.ErrEmptyProductName, err)
	_, err = models.NewProduct("PRD1", "Noodles", -1)
	c.Equal(models.ErrEmptyProductPrice, err)
}
