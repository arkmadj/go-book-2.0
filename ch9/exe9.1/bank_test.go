package bank

import (
	"fmt"
	"testing"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	go func() {
		Deposits(200)
		Withdraw(200)
		fmt.Println("=", Balance())
		done <- struct{}{}
	}()

	go func() {
		Deposits(50)
		Withdraw(50)
		Deposits(100)
		done <- struct{}{}
	}()

	<-done
	<-done

	if got, want := Balance(), 100; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}

func TestWithdrawal(t *testing.T) {
	b1 := Balance()
	ok := Withdraw(50)
	if !ok {
		t.Errorf("ok = false, want true. balance = %d", Balance())
	}
	expected := b1 - 50
	if b2 := Balance(); b2 != expected {
		t.Errorf("balance = %d, want %d", b2, expected)
	}
}
