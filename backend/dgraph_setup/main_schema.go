package main

import (
	"context"
	"fmt"
	"log"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	dgraphClient := dgo.NewDgraphClient(api.NewDgraphClient(conn))

	op := &api.Operation{
		Schema: `
		# Define Types
		
		type Buyer {
			id
			name
			age 
		}
		
		type Product {
			id
			name
			price
		}

		type Transaction {
			id
			buyerID
			ip
			device
			productsIDs
		}
		
		# Define Directives and index
		
		productLinker: [uid] @reverse .
		transactionLinker: [uid] @reverse .
		id: string @index(hash) .
		name: string .
		age: int .
		price: int .
		buyerID: string @index(hash) .
		ip: string @index(hash) .
		device: string @index(hash) .
		productsIDs: [string] @index(hash) .`,
	}

	err = dgraphClient.Alter(context.Background(), op)
	if err != nil {
		log.Fatalln("Error Altering the database, schema couldn't be set.", err)
	} else {
		fmt.Println("Altering database succesful, the schema was set.")
	}
}
