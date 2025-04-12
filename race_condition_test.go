package go_coroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRaceConditon(t *testing.T) {
	x := 0

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				x += 1
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("count:", x)
}

func TestMutext(t *testing.T) {
	x := 0
	var mutext sync.Mutex

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				mutext.Lock()
				x += 1
				mutext.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("count:", x)
}

type BankAccount struct {
	RWMutext sync.RWMutex
	Balance  int
}

func (account *BankAccount) Deposit(amount int) {
	account.RWMutext.Lock()
	account.Balance += amount
	defer account.RWMutext.Unlock()

}

func (account *BankAccount) ReadBalance() int {
	account.RWMutext.RLock()
	defer account.RWMutext.RUnlock()
	return account.Balance

}

func TestReadWriteMutext(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.Deposit(1)
				fmt.Println(account.ReadBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("count:", account.ReadBalance())
}
