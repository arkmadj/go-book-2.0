package bank

type Withdrawal struct {
	amount  int
	success chan bool
}

var deposits = make(chan int)
var balances = make(chan int)
var withdrawals = make(chan Withdrawal)

func Deposits(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func Withdraw(amount int) bool {
	ch := make(chan bool)
	withdrawals <- Withdrawal{amount, ch}
	return <-ch
}
