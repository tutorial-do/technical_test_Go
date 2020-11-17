package models

import (
	"errors"
)

var (
	// ErrEmptyTransactionID is used to indicate that the transaction id is empty.
	ErrEmptyTransactionID = errors.New("transaction 'id' cannot be empty")

	// ErrEmptyTransactionBuyerID is used to indicate that the buyerID is empty.
	ErrEmptyTransactionBuyerID = errors.New("transaction 'buyerID' cannot be empty")

	// ErrEmptyTransactionIP is used to indicate that the ip from where the transaction was made is empty.
	ErrEmptyTransactionIP = errors.New("transaction 'ip' cannot be empty")

	// ErrEmptyTransactionDevice is used to indicate that the device from where the transaction was made is empty.
	ErrEmptyTransactionDevice = errors.New("transaction 'device' cannot be empty")

	// ErrEmptyProductsIDs is used to indicate that the list of products in the transaction is empty.
	ErrEmptyProductsIDs = errors.New("'productsIDs' cannot be empty, must have at least one product in the transaction")
)

// Transaction Representation of transactions
type Transaction struct {
	ID          string   `json:"id,omitempty"`
	BuyerID     string   `json:"buyerID,omitempty"`
	IP          string   `json:"ip,omitempty"`
	Device      string   `json:"device,omitempty"`
	ProductsIDs []string `json:"productsIDs,omitempty"`
	DType       []string `json:"dgraph.type,omitempty"`
}

// NewTransaction Constructor For a Transaction
func NewTransaction(id string, buyerID string, ip string, device string, productsIds []string) (*Transaction, error) {
	if id == "" {
		return nil, ErrEmptyTransactionID
	}

	if buyerID == "" {
		return nil, ErrEmptyTransactionBuyerID
	}

	if ip == "" {
		return nil, ErrEmptyTransactionIP
	}

	if device == "" {
		return nil, ErrEmptyTransactionDevice
	}

	if len(productsIds) <= 0 {
		return nil, ErrEmptyProductsIDs
	}

	newTransaction := &Transaction{
		ID:          id,
		BuyerID:     buyerID,
		IP:          ip,
		Device:      device,
		ProductsIDs: productsIds,
		DType:       []string{"Transaction"},
	}

	return newTransaction, nil
}
