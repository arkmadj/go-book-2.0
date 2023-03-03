package bank

var (
	sema    = make(chan struct{}, 1)
	balance int
)
