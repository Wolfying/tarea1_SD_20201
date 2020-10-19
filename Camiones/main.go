package main

import (
	"fmt"
	"time"
)

// import (
// 	"fmt"
// 	"math"
// )

const (
	retail      int = 1
	normal      int = 2 //pyme
	prioritario int = 3
)

type registro struct {
	id           int
	tipo         int
	valor        int
	origen       string
	destino      string
	intenteos    int
	fechaEntrega string
}

type camion struct {
	carga   int
	tipo    int
	informe *[]registro
}

func (c camion) detalles() {
	registro := c.informe
	fmt.Println(*registro, " Camión")
}

func generarCamiones() []camion {
	var registros1 []registro
	camion1 := camion{0, 1, &registros1}
	var registros2 []registro
	camion2 := camion{0, 1, &registros2}
	var registros3 []registro
	camion3 := camion{0, 2, &registros3}

	var camiones = []camion{camion1, camion2, camion3}
	return camiones

}

func main() {
	var registros []registro

	carga := registro{1, 1, 1, "asda", "asda", 1, time.Now().Format("02-01-2006 15:04")}

	registros = append(registros, carga)
	var camiones = generarCamiones()

	for _, camione := range camiones {
		camione.detalles()
	}

	camiones[1].detalles()

	fmt.Printf(" %v \n", registros)
}
