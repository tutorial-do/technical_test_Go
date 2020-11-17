package storage

import (
	"context"
	"log"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"google.golang.org/grpc"
)

// DatabaseConnection creates the connection with Dgraph
func DatabaseConnection() (*dgo.Dgraph, func()) {
	conn, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
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

func SaveData(dgraphClient *dgo.Dgraph, dataJSON []byte) error {
	ctx := context.Background()

	txn := dgraphClient.NewTxn()

	defer txn.Discard(ctx)

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
