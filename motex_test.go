package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	x := 0

	var mutex sync.Mutex
	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock() //! Menghindari atau menghandle rest condition, sehigga ada antrian
				x = x + 1
				mutex.Unlock() //! Menghindari atau menghandle rest condition, sehigga ada antrian
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("counter: ", x)
}

type BankAccount struct { 
	RWMutex sync.RWMutex //! Jika struct di akses oleh beberapa goroutine sekaligus
	Balance int
}

func (account *BankAccount) Addbalance(amout int) {
	account.RWMutex.Lock() //! untuk proeses write
	account.Balance = account.Balance + amout
	account.RWMutex.Unlock() //! untuk proeses write
}

func (account *BankAccount) GetBalance() int  {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i:= 0; i < 100; i++ {
		go func ()  {
			for j:= 0; j < 100; j++ {
				account.Addbalance(1)
				fmt.Println(account.GetBalance())

			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Total balance", account.GetBalance())
}


type UserBalance struct {
	sync.Mutex
	Name string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int)  {
	user.Balance += amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock user1", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock user2", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadLock(t *testing.T)  {
	user1 := UserBalance{
		Name: "eko",
		Balance: 1000000,
	}

	user2 := UserBalance{
		Name: "budi",
		Balance: 1000000,
	}

	go Transfer(&user1, &user2, 100000)
	go Transfer(&user2, &user1, 200000)

	time.Sleep(10 * time.Second)

	fmt.Println("User: ", user1.Name, ", Balance ", user1.Balance)
	fmt.Println("User: ", user2.Name, ", Balance ", user2.Balance)
}

