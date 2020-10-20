package main

import (
	"context"
	"fmt"
	"log"
	"tarea1/Logistica/logistica"

	"google.golang.org/grpc"
)

const (
	address = "localhost:8888"
)

const (
	retail      int = 1
	normal      int = 2 //pyme
	prioritario int = 3
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := logistica.NewLogisticaOrdenesClient(conn)
	orden := logistica.Orden{
		Id:       "producto1",
		Producto: "leche de palbo",
		Precio:   9999,
		Origen:   "Palbito",
		Destino:  "Jorgito",
		Tipo:     1,
	}

	response, err := c.IngresarOrden(context.Background(), &orden)

	if err != nil {
		fmt.Printf("Oopsie Dupsy %s", err)
	} else {
		fmt.Printf("%s", response)
	}

}
