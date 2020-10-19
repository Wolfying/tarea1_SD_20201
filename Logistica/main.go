package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port     = ":8888"
	camiones = "dist"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
