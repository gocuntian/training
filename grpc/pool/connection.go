package pool

import (
	"google.golang.org/grpc"
)

// Connection grpc connection inerface
type Connection interface {
	Value() *grpc.ClientConn
	Close() error
}

type connection struct {
	clientConn *grpc.ClientConn
	pool       *pool
	once       bool
}

// Value see Conn interface.
func (c *connection) Value() *grpc.ClientConn {
	return c.clientConn
}

func (c *connection) Close() error {
	c.pool.decrRef()
	if c.once {
		return c.reset()
	}
	return nil
}

func (c *connection) reset() error {
	cc := c.clientConn
	c.clientConn = nil
	c.pool = nil
	c.once = false
	if cc != nil {
		return cc.Close()
	}
	return nil
}

func (p *pool) wrapConn(clientConn *grpc.ClientConn, once bool) *connection {
	return &connection{
		clientConn: clientConn,
		pool:       p,
		once:       once,
	}
}
