package routes

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"technical_test_Go/backend/api/handlers"
	"technical_test_Go/backend/storage"

	"github.com/dgraph-io/dgo"
)

// DataLoader function to load the data from the endpoints
func DataLoader(dbConnection *dgo.Dgraph) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		date := ""
		dataBuyers, dataProducts, dataTransactions := handlers.FetchData(date)

		// UniqueBuyers, UniqueProducts, UniqueTransactions, err := handlers.Strainer(dbConnection, dataBuyers, dataProducts, dataTransactions)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		buyerJSON, err := json.Marshal(dataBuyers)
		if err != nil {
			log.Fatal(err)
		}
		err = storage.SaveData(dbConnection, buyerJSON)
		if err != nil {
			log.Fatal(err)
		}

		// productJSON, err := json.Marshal(UniqueProducts)
		productJSON, err := json.Marshal(dataProducts)
		if err != nil {
			log.Fatal(err)
		}

		err = storage.SaveData(dbConnection, productJSON)
		// err = storage.SaveData(dbConnection, productJSON)
		if err != nil {
			log.Fatal(err)
		}

		transactionJSON, err := json.Marshal(dataTransactions)
		// transactionJSON, err := json.Marshal(UniqueTransactions)
		if err != nil {
			log.Fatal(err)
		}

		err = storage.SaveData(dbConnection, transactionJSON)
		if err != nil {
			log.Fatal(err)
		}

		// function to connect the data
		// ConnectData(dbConnection, dataBuyers, dataProducts, dataTransactions)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// w.Write(buyerJSON)
		// w.Write(productJSON)
		// w.Write(transactionJSON)
		w.Write([]byte("Data succesfully loaded"))
	}
}

// GetAllBuyers functions to retrieve all buyers from the database
func GetAllBuyers(dgraphClient *dgo.Dgraph) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		q := `{
			allBuyers(func: type(Buyer)){
				id
				name
				age
			}
		}`

		txn := dgraphClient.NewTxn()

		defer txn.Discard(ctx)

		res, err := txn.Query(ctx, q)
		if err != nil {
			log.Fatal(err)
		}

		w.Write(res.Json)
	}
}

// GetAllProducts functions to retrieve all products from the database
func GetAllProducts(dgraphClient *dgo.Dgraph) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		q := `{
			allProducts(func: type(Product)){
				id
				name
				price
			}
		}`

		txn := dgraphClient.NewTxn()

		defer txn.Discard(ctx)

		res, err := txn.Query(ctx, q)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(res.Json)
	}
}

// GetAllProducts functions to retrieve all products from the database
func GetAllTransactions(dgraphClient *dgo.Dgraph) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		q := `{
			allTransactions(func: type(Transaction)){
				id
    		buyerID
    		ip
    		device
    		productsIDs
			}
		}`

		txn := dgraphClient.NewTxn()

		defer txn.Discard(ctx)

		res, err := txn.Query(ctx, q)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(res.Json)
	}
}
