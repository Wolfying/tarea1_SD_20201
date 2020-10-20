package logistica

// implementacion servicios
import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

const (
	retail      int = 1
	normal      int = 2 //pyme
	prioritario int = 3
	seed        int = 20
)

type logisticaServer struct {
	listaRegistros []*registro
}

type registro struct {
	timestamp   string
	idPaquete   string
	tipo        int32
	producto    string
	precio      int32
	origen      string
	destino     string
	seguimiento int32
}

func (ls *logisticaServer) IngresarOrden(ctx context.Context, o *Orden) (*Response, error) {

	seg := rand.Int31()
	r := registro{
		timestamp:   time.Now().Format("02-01-2006 15:04"),
		idPaquete:   o.Id,
		tipo:        o.Tipo,
		producto:    o.Producto,
		precio:      o.Precio,
		origen:      o.Origen,
		destino:     o.Destino,
		seguimiento: seg,
	}

	ls.listaRegistros = append(ls.listaRegistros, &r)
	fmt.Printf("asdasdasdasda")
	return &Response{Status: fmt.Sprintf("Se ha ingresado la orden %d", seg), IdSeguimiento: seg}, nil
}

func (ls *logisticaServer) ConsultaSeguimiento(ctx context.Context, r *Response) (*Response, error) {
	return &Response{}, nil
}
