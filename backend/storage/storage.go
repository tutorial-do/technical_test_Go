package storage

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"technical_test_Go/backend/models"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"google.golang.org/grpc"
)

// DatabaseConnection creates the connection with Dgraph
func DatabaseConnection() (*dgo.Dgraph, func()) {
	conn, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Error while trying to dial gRPC")
	}

	dgraphClient := dgo.NewDgraphClient(api.NewDgraphClient(conn))

	closeDB := func() {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}

	return dgraphClient, closeDB
}

// SaveData is a function to save data to the database Dgraph
func SaveData(dgraphClient *dgo.Dgraph, dataJSON []byte) error {
	ctx := context.Background()

	txn := dgraphClient.NewTxn()

	mu := &api.Mutation{
		SetJson: dataJSON,
	}
	mu.CommitNow = true

	_, err := txn.Mutate(ctx, mu)
	if err != nil {
		return err
	}

	return nil
}

// DeleteAll function to delete all data and schema from the database
func DeleteAll() {
	dgraphClient, cancel := DatabaseConnection()

	defer cancel()

	op := api.Operation{DropAll: true}
	ctx := context.Background()

	if err := dgraphClient.Alter(ctx, &op); err != nil {
		log.Fatal("The drop all operation should have succeeded")
	}
}

// ConnectData function to connect the data within the Database creating links between related nodes
func ConnectData(dbClient *dgo.Dgraph, dataBuyers []models.Buyer, dataProducts []models.Product, dataTransactions []models.Transaction) error {
	var buyerIDs []string
	var productIDs []string

	for _, buyerElement := range dataBuyers {
		buyerIDs = append(buyerIDs, buyerElement.ID)
	}
	for _, productElement := range dataProducts {
		productIDs = append(productIDs, productElement.ID)
	}

	fmt.Println("Entre a connect")

	err := linkCreatorBuyers(dbClient, buyerIDs)
	if err != nil {
		return err
	}
	err = linkCreatorProducts(dbClient, productIDs)
	if err != nil {
		return err
	}

	return nil
}

func linkCreatorBuyers(dbClient *dgo.Dgraph, buyerIDs []string) error {
	var upsertQuery bytes.Buffer

	fmt.Println("Entre a Linker de Buyers")

	upsertQuery.WriteString("query {")

	var upsertMutation []*api.Mutation

	for index, value := range buyerIDs {
		q := fmt.Sprintf(`
			predicate1_%d as var(func: eq(buyerID, "%s"))
			predicate2_%d as var(func: eq(id, "%s"))
		`, index, value, index, value)

		upsertQuery.WriteString(q)

		mu := &api.Mutation{
			SetNquads: []byte(fmt.Sprintf(`uid(predicate1_%d) <buyerLinker> uid(predicate2_%d) .`, index, index)),
		}

		upsertMutation = append(upsertMutation, mu)
	}

	req := &api.Request{
		Query:     upsertQuery.String(),
		Mutations: upsertMutation,
		CommitNow: true,
	}
	_, err := dbClient.NewTxn().Do(context.Background(), req)
	if err != nil {
		return err
	}

	return nil
}

func linkCreatorProducts(dbClient *dgo.Dgraph, productIDs []string) error {
	var upsertQuery bytes.Buffer

	fmt.Println("Entre a Linker de Products")

	upsertQuery.WriteString("query {")

	var upsertMutation []*api.Mutation

	for index, value := range productIDs {
		q := fmt.Sprintf(`
			predicate1_%d as var(func: eq(productsIDs, "%s"))
			predicate2_%d as var(func: eq(id, "%s"))
		`, index, value, index, value)

		upsertQuery.WriteString(q)

		mu := &api.Mutation{
			SetNquads: []byte(fmt.Sprintf(`uid(predicate1_%d) <productLinker> uid(predicate2_%d) .`, index, index)),
		}

		upsertMutation = append(upsertMutation, mu)
	}

	req := &api.Request{
		Query:     upsertQuery.String(),
		Mutations: upsertMutation,
		CommitNow: true,
	}
	_, err := dbClient.NewTxn().Do(context.Background(), req)
	if err != nil {
		return err
	}

	return nil
}
