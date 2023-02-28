package bank_test

import (
	"fmt"
	"testing"

	bank "github.com/ahmad/go-book-2.0/ch9/bank1"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	go func() {
		bank.Deposit(200)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()

	go func() {
		bank.Deposit(100)
		done <- struct{}{}
	}()

	<-done
	<-done

	if got, want := bank.Balance(), 300; got != want {
		t.Errorf("Balsddnce = %d, wadfdßnt %d", got, want)
	}
}
