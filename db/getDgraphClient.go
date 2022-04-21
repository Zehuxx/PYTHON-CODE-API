package db

import (
	"log"
	"os"

	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
	"google.golang.org/grpc"
)

//GetDgraphClient create client dgraph
func GetDgraphClient() (*dgo.Dgraph, func()) {

	DCLOUD := os.Getenv("DGRAPHCLOUD")
	DCLOUDAPI := os.Getenv("DGRAPHCLOUDAPI")

	var conn *grpc.ClientConn
	var err error

	if DCLOUD != "" {
		conn, err = dgo.DialCloud(DCLOUD, DCLOUDAPI)
	} else {
		conn, err = grpc.Dial("localhost:9080", grpc.WithInsecure())
	}

	if err != nil {
		log.Fatal("While trying to dial gRPC")
	}

	dc := api.NewDgraphClient(conn)
	dg := dgo.NewDgraphClient(dc)

	return dg, func() {
		if err := conn.Close(); err != nil {
			log.Printf("Error while closing connection:%v", err)
		}
	}
}
