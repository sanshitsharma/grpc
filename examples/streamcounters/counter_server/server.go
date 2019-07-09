package main

import (
	"errors"
	"google.golang.org/grpc"
	pb "github.com/sanshitsharma/grpc/examples/streamcounters/counters"
	"log"
	"net"
	"time"
)

// Config holds the server context
type Config struct{}

func (s *Config) GetCounters(req *pb.CounterReq, stream pb.Status_GetCountersServer) error {
	if req.GetClientId() == `` {
		return errors.New(`invalid client ID`)
	}

	resp := &pb.CounterResp{
		Ok: true,
	}

	for counter := 0; counter < 5; counter++ {
		resp.Counter = int32(counter+1)
		time.Sleep(3*time.Second)
		if err := stream.Send(resp); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", `:50052`)
	if err != nil {
		log.Fatalf("failed to create listener. err: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterStatusServer(s, &Config{})
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to start grpc server. err: %v", err)
	}
}
