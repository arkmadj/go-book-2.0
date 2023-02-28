package bank

type Withdrawal struct {
	amount  int
	success chan bool
}

var deposits = make(chan int)
var balances = make(chan int)
var withdrawals = make(chan Withdrawal)
