package connpool

import (
	"context"
)

type ConnPool interface {
	// Get return a connection TODO:connection struct
	Get(ctx context.Context, address string) error
	Put() error
	Ping() error
	Discard() error
	Close() error
}
