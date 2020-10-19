package main

import (
	"log"
	"tarea1/Logistica/logistica"

	"google.golang.org/grpc"
)

const (
	address = "localhost:8888"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := logistica.NewLogisticaOrdenesClient(conn)

}
