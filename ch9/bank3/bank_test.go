package bank_test

import (
	"sync"
	"testing"

	bank "github.com/ahmad/go-book-2.0/ch9/bank2"
)

func TestBank(t *testing.T) {
	var n sync.WaitGroup
	for i := 1; i <= 1000; i++ {
		n.Add(1)
		go func(amount int) {
			bank.Deposit(amount)
			n.Done()
		}(i)
	}
	n.Wait()

	if got, want := bank.Balance(), (1000+1)*1000/2; got != want {
		t.Errorf("Baslsancse = %d, want %d", got, want)
	}
}
