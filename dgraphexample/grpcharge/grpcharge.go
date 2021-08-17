package grpcharge

import (
	"context"
	"dgraphexample/importendpoint"
	"encoding/json"
	"fmt"
	"log"

	//"strings"

	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
	"google.golang.org/grpc"
)

type CancelFunc func()

func getDgraphClient() (*dgo.Dgraph, CancelFunc) {
	conn, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("While trying to dial gRPC")
	}

	dc := api.NewDgraphClient(conn)
	dg := dgo.NewDgraphClient(dc)
	//ctx := context.Background()

	// Perform login call. If the Dgraph cluster does not have ACL and
	// enterprise features enabled, this call should be skipped.
	/*
		for {
			// Keep retrying until we succeed or receive a non-retriable error.
			err = dg.Login(ctx, "groot", "password")
			if err == nil || !strings.Contains(err.Error(), "Please retry") {
				break
			}
			time.Sleep(time.Second)
		}
	*/
	if err != nil {
		log.Fatalf("While trying to login %v", err.Error())
	}

	return dg, func() {
		if err := conn.Close(); err != nil {
			log.Printf("Error while closing connection:%v", err)
		}
	}
}

func MutatedGraph() {

	p := importendpoint.GenerateShop()
	fmt.Print("Archivo Generado")

	dg, cancel := getDgraphClient()
	defer cancel()

	ctx := context.Background()
	mu := &api.Mutation{
		CommitNow: true,
	}

	pb, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	//println(string(pb))

	mu.SetJson = pb
	response, err := dg.NewTxn().Mutate(ctx, mu)
	if err != nil {
		log.Fatal(err)
	}

	println(response)

}

func QueryListBuyer() string {

	dg, cancel := getDgraphClient()
	defer cancel()

	q := `{
			var(func:has(Transaction.idTran)){			
			Transaction.Buyer{
			b as Buyer.idBuyer
			#Buyer.name
			#Buyer.age  
			}  			  
			}			
			q(func:has(Buyer.idBuyer)) @filter(eq(Buyer.idBuyer,val(b))){			
			idBuyer:Buyer.idBuyer
			name:Buyer.name
			age:Buyer.age
			}			  
			}
		   `

	resp, err := dg.NewTxn().Query(context.Background(), q)

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("Response: %s\n", resp.Json)

	return string(resp.Json)
}

func QueryListHistory(sidBuyer string) string {

	dg, cancel := getDgraphClient()
	defer cancel()

	q := `{
		q(func:has(Transaction.idTran))
		  @filter(eq(Transaction.idBuyer,"` + sidBuyer + `"))
		{
		idTran:Transaction.idTran
		ip    :Transaction.ip  
		
		Products:Transaction.Products{
		idProduct:Product.idProduct
		name     :Product.name
		price    :Product.price  
		}  
		  
		}
		}`
	//fmt.Printf("Response: %s\n", q)
	resp, err := dg.NewTxn().Query(context.Background(), q)

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("Response: %s\n", resp.Json)

	return string(resp.Json)
}

func QueryBuyersGraph() {

	p := importendpoint.GenerateShop()
	fmt.Print("Archivo Generado")

	dg, cancel := getDgraphClient()
	defer cancel()

	ctx := context.Background()
	mu := &api.Mutation{
		CommitNow: true,
	}

	pb, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	//println(string(pb))

	mu.SetJson = pb
	response, err := dg.NewTxn().Mutate(ctx, mu)
	if err != nil {
		log.Fatal(err)
	}

	println(response)

}
