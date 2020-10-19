// implementacion servicios
package logistica

import "context"

type logisticaServer struct {
}

const (
	retail      int = 1
	normal      int = 2 //pyme
	prioritario int = 3
)

type registro struct {
	timestamp string
	idPaquete int
}

func (ls *logisticaServer) ingresarOrden(ctx context.Context, o *Orden) (*Response, error) {
	return &Response{}, nil
}

func (ls *logisticaServer) consultaSeguimiento(ctx context.Context, r *Response) (*Response, error) {
	return &Response{}, nil
}
