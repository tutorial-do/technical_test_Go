package models

import (
	"technical_test_Go/backend/models"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewTransaction(t *testing.T) {
	c := require.New(t)
	buyer, err := models.NewBuyer("BYR1", "Julian", 28)
	product1, err := models.NewProduct("PRD1", "Noodles", 3045)
	product2, err := models.NewProduct("PRD2", "lemonade", 1000)
	productsIDs := make([]string, 2)
	productsIDs[0] = product1.ID
	productsIDs[1] = product2.ID
	transaction, err := models.NewTransaction("TRX1", buyer.ID, "194.35.118.92", "linux", productsIDs)
	c.NotEmpty(transaction.ID)
	c.NotEmpty(transaction.BuyerID)
	c.Equal("linux", transaction.Device)
	c.Equal([]string{"PRD1", "PRD2"}, transaction.ProductsIDs)
	c.NoError(err)
}
