package main

import (
	"fmt"
	"sync"
	"time"
)

type BankAccount struct {
	balance int
	mu      sync.Mutex
}

func (ba *BankAccount) Withdraw(amount int) bool {
	ba.mu.Lock()
	defer ba.mu.Unlock()
	if ba.balance >= amount {
		time.Sleep(time.Microsecond)
		ba.balance -= amount
		return true

	}

	return false
}
func (ba *BankAccount) GetBalance() int {

	return ba.balance
}

func main() {
	account := &BankAccount{balance: 1000}
	var wg sync.WaitGroup

	for i := range 10 {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			success := account.Withdraw(200)
			fmt.Printf("Goroutine %d: attempt to withdraw 200 uah - %t\n", id, success)
		}(i)
	}
	wg.Wait()
	fmt.Printf("Final balance:  %d uah\n", account.GetBalance())
}
