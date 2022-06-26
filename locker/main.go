package main

import (
	"fmt"
	"sync"
	"time"
)

type BankAccountV1 struct {
	id      string
	balance int
}

func (b *BankAccountV1) add(ba int) {
	b.balance = b.balance + ba
}

func (b *BankAccountV1) getBalance() int {
	return b.balance
}

func TestV1Balance() {
	b1 := BankAccountV1{
		id:      "b1",
		balance: 0,
	}

	for i := 0; i < 1000; i++ {
		go b1.add(100)
	}

	time.Sleep(2 * time.Second) // 96600
	fmt.Println("balance:", b1.getBalance())
}

type BankAccountV2 struct {
	id      string
	balance int
	locker  sync.Mutex
}

func (b *BankAccountV2) add(ba int) {
	b.locker.Lock()
	defer b.locker.Unlock()

	b.balance = b.balance + ba
}

func (b *BankAccountV2) getBalance() int {
	b.locker.Lock()
	defer b.locker.Unlock()

	return b.balance
}

func TestV2Balance() {
	b1 := BankAccountV2{
		id:      "b1",
		balance: 0,
	}

	for i := 0; i < 1000; i++ {
		go b1.add(100)
	}

	time.Sleep(2 * time.Second) // 100000
	fmt.Println("balance:", b1.getBalance())
}

type BankAccountV3 struct {
	id      string
	balance int
	locker  sync.RWMutex
}

func (b *BankAccountV3) add(ba int) {
	b.locker.Lock()
	defer b.locker.Unlock()

	b.balance = b.balance + ba
}

func (b *BankAccountV3) getBalance() int {
	b.locker.RLock()
	defer b.locker.RUnlock()

	return b.balance
}

func TestV3Balance() {
	b1 := BankAccountV3{
		id:      "b1",
		balance: 0,
	}

	for i := 0; i < 1000; i++ {
		go b1.add(100)
	}

	time.Sleep(2 * time.Second) // 100000
	fmt.Println("balance:", b1.getBalance())
}

func main() {
	//TestV2Balance()
	TestV3Balance()
}
