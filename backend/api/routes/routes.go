package routes

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"technical_test_Go/backend/api/handlers"
	"technical_test_Go/backend/storage"

	"github.com/dgraph-io/dgo"
	"github.com/go-chi/chi"
)

// DataLoader function to load the data from the endpoints
func DataLoader(dbClient *dgo.Dgraph) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		date := chi.URLParam(r, "date")
		dataBuyers, dataProducts, dataTransactions := handlers.FetchData(date)

		uniqueBuyers, uniqueProducts, uniqueTransactions, err := handlers.DataFilter(dbClient, dataBuyers, dataProducts, dataTransactions)
		if err != nil {
			log.Fatal(err)
		}

		// buyerJSON, err := json.Marshal(dataBuyers)
		buyerJSON, err := json.Marshal(uniqueBuyers)
		if err != nil {
			log.Fatal(err)
		}
		err = storage.SaveData(dbClient, buyerJSON)
		if err != nil {
			log.Fatal(err)
		}

		// productJSON, err := json.Marshal(dataProducts)
		productJSON, err := json.Marshal(uniqueProducts)
		if err != nil {
			log.Fatal(err)
		}

		err = storage.SaveData(dbClient, productJSON)
		if err != nil {
			log.Fatal(err)
		}

		// transactionJSON, err := json.Marshal(dataTransactions)
		transactionJSON, err := json.Marshal(uniqueTransactions)
		if err != nil {
			log.Fatal(err)
		}

		err = storage.SaveData(dbClient, transactionJSON)
		if err != nil {
			log.Fatal(err)
		}
		// function to connect the data
		// storage.ConnectData(dbClient, dataBuyers, dataProducts, dataTransactions)
		storage.ConnectData(dbClient, uniqueBuyers, uniqueProducts, uniqueTransactions)
		if err != nil {
			log.Fatal(err)
		}
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

// GetDataBuyerbyID function to find the information of a buyer by ID
func GetDataBuyerbyID(dgraphClient *dgo.Dgraph) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		ctx := context.Background()
		q := `query BuyerData($id: string){
			buyerInformation(func: eq(id, $id)){
				~buyerLinker{
					buyerID
					name
					id
					device
					sameIP as ip
					productLinker{
						name
						price
						id
					}
				}
		},
			{
				sameIPBuyers(func: eq(ip, val(sameIP))){
					ip
					buyerLinker{
						id
						name
						age
					}
					productLinker{
						name
					}
				}
			},
			{
				recomendedProducts(func: type(Product)){
					id
					name
					price
				}
			}  
		}`

		resp, err := dgraphClient.NewTxn().QueryWithVars(ctx, q, map[string]string{"$id": id})
		if err != nil {
			log.Fatal(err)
		}

		w.Write(resp.Json)
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
