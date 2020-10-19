// implementacion servicios
package logistica

import "context"

type logisticaServer struct {
}

func (ls *logisticaServer) ingresarOrden(ctx context.Context, o *Orden) (*Response, error) {
	return &Response{}, nil
}
