package bank

import "sync"

var (
	mu      sync.Mutex
	balance int
)
