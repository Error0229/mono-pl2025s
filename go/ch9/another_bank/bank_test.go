// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package bank_test

import (
	"fmt"
	"testing"

	bank "gopl.io/ch9/another_bank"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})
	for i := 0; i < 10; i++ {

		// Alice
		go func() {
			bank.Deposit(200)
			fmt.Println("=", bank.Balance())
			done <- struct{}{}
		}()

		// Bob
		go func() {
			bank.Deposit(100)
			done <- struct{}{}
		}()

		// Wait for both transactions.
		<-done
		<-done

		if got, want := bank.Balance(), 300; got != want {
			t.Errorf("Balance = %d, want %d", got, want)
		}
		bank.Reset()
	}
}
