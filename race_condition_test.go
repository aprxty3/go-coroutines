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

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Deposit(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Locking user1", user1.Name)
	user1.Deposit(-amount)
	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Locking user2", user2.Name)
	user2.Deposit(amount)
	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlock(t *testing.T) {
	user1 := &UserBalance{Name: "user1", Balance: 1000}
	user2 := &UserBalance{Name: "user2", Balance: 1000}

	go Transfer(user1, user2, 100)
	go Transfer(user2, user1, 200)

	time.Sleep(5 * time.Second)

	fmt.Println("user", user1.Name, "balance:", user1.Balance)
	fmt.Println("user", user2.Name, "balance:", user2.Balance)

}
