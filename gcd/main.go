package main

import (
	"fmt"
	"log"
	"net"

	"github.com/krufyliu/kubernetes-go-grpc-tutorial/pb"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGCDServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func gcd(a, b uint64) uint64 {
	if b%a == 0 {
		return a
	}
	return gcd(b%a, a)
}

func (s *server) Compute(ctx context.Context, r *pb.GCDRequest) (*pb.GCDResponse, error) {
	res := new(pb.GCDResponse)
	if r.A == 0 || r.B == 0 {
		return nil, fmt.Errorf("parameters can not be zero")
	}
	res.Result = gcd(r.A, r.B)
	return res, nil
}
