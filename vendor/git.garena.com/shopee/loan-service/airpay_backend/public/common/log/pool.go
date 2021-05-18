package log

import (
	"go.uber.org/zap/buffer"
)

var (
	_pool = buffer.NewPool()
	// Get retrieves a buffer from the pool, creating one if necessary.
	GetBuffer = _pool.Get
)
