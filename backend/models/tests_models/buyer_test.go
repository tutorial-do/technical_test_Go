package models

import (
	"technical_test_Go/backend/models"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewBuyer(t *testing.T) {
	c := require.New(t)
	buyer, err := models.NewBuyer("BYR1", "Julian", 28)
	c.NotEmpty(buyer.ID)
	c.NotEmpty(buyer.Name)
	c.Equal("Julian", buyer.Name)
	c.Equal(28, buyer.Age)
	c.NoError(err)
}

func TestNewBuyerErrors(t *testing.T) {
	c := require.New(t)

	_, err := models.NewBuyer("", "Julian", 28)
	c.Equal(models.ErrEmptyBuyerID, err)
	_, err = models.NewBuyer("BYR1", "", 28)
	c.Equal(models.ErrEmptyBuyerName, err)
	_, err = models.NewBuyer("BYR1", "Julian", -1)
	c.Equal(models.ErrEmptyBuyerAge, err)
}
