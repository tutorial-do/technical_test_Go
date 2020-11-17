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

		buyerJSON, err := json.Marshal(dataBuyers)
		if err != nil {
			log.Fatal(err)
		}
		err = storage.SaveData(dbConnection, buyerJSON)
		if err != nil {
			log.Fatal(err)
		}

		productJSON, err := json.Marshal(dataProducts)
		if err != nil {
			log.Fatal(err)
		}

		err = storage.SaveData(dbConnection, productJSON)
		if err != nil {
			log.Fatal(err)
		}

		transactionJSON, err := json.Marshal(dataTransactions)
		if err != nil {
			log.Fatal(err)
		}

		err = storage.SaveData(dbConnection, transactionJSON)
		if err != nil {
			log.Fatal(err)
		}

		// w.Write(buyerJSON)
		// w.Write(productJSON)
		// w.Write(transactionJSON)
	}
}

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
		// w.Write([]byte("welcome"))
		w.Write(res.Json)
	}
}
