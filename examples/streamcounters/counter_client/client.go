package main

import (
	"context"
	"flag"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "github.com/sanshitsharma/grpc/examples/streamcounters/counters"
)

var (
	serverAddress = flag.String(`server_addr`, `localhost:50052`, `The server address in the format host:port`)
)

func main() {
	log.Println(`Welcome to streaming client test..`)
	flag.Parse()

	connection, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf(`failed to dial: %v`, err)
	}
	defer connection.Close()

	client := pb.NewStatusClient(connection)

	ctx, cancel := context.WithTimeout(context.TODO(), 20*time.Second)
	defer cancel()

	// Test 1: Invalid client ID
	log.Println(`******* TEST 1 *******`)
	req1 := &pb.CounterReq{
		ClientId: ``,
	}
	stream, err := client.GetCounters(ctx, req1)
	if err != nil {
		log.Println(`Expected Failure: %v.GetServiceState(_) = _, %v`, client, err)
	}
	log.Printf(`req: %v err: %v`, req1, err)
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf(`T1 stream %v.GetServiceState(_) = _, %v`, client, err)
		}

		log.Println(resp)
	}

	// Test 2: Valid I/P
	log.Println(`******* TEST 2 *******`)
	req2 := &pb.CounterReq{
		ClientId: `sanshit`,
	}
	stream, err = client.GetCounters(ctx, req2)
	if err != nil {
		log.Fatalf(`T2 %v.GetCounters(_) = _, %v`, client, err)
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf(`T2 stream %v.GetCounters(_) = _, %v`, client, err)
		}

		log.Println(resp)
	}

	// Test 3: Deadline exceed
	log.Println(`******* TEST 3 *******`)
	ctx1, cancel1 := context.WithTimeout(context.TODO(), 10*time.Second)
	req3 := &pb.CounterReq{
		ClientId: `sansshar`,
	}
	stream, err = client.GetCounters(ctx1, req3)
	if err != nil {
		log.Fatalf(`T3 %v.GetCounters(_) = _, %v`, client, err)
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf(`T3 stream %v.GetCounters(_) = _, %v`, client, err)
		}

		log.Println(resp)
	}
	cancel1()
}
