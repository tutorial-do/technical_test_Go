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

// func TestNewChannelErrors(t *testing.T) {
// 	c := require.New(t)

// 	_, err := NewBuyer("BYR1", "Julian", 28)
// 	c.Equal(ErrEmptyDescription, err)
// 	_, err = NewBuyer("BYR1", "Julian", 28)
// 	c.Equal(ErrEmptyTitle, err)
// 	_, err = NewBuyer("BYR1", "Julian", 28)
// 	c.Equal(ErrEmptyCreatorID, err)
// }
