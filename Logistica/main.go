package main

import (
	"log"
	"math/rand"
	"net"
	"tarea1/Logistica/logistica"

	"google.golang.org/grpc"
)

const (
	port     = ":8888"
	camiones = "dist"
)

type server struct {
	logistica.UnimplementedLogisticaOrdenesServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	rand.Seed(20)
	s := grpc.NewServer()
	logistica.RegisterLogisticaOrdenesServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
