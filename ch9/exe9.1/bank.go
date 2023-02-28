package bank

type Withdrawal struct {
	amount  int
	success chan bool
}

var deposits = make(chan int)
