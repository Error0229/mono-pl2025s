package bank

import (
	"time"

	"math/rand"
)

var balance int

func Deposit(amount int) {
	temp := balance
	// Simulate a random delay to allow concurrent access
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	balance = temp + amount
}
func Balance() int { return balance }
func Reset()       { balance = 0 }
