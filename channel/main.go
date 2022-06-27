package main

import (
	"fmt"
	"time"
)

func testNonBufChan() {
	var ic chan int
	fmt.Println(ic) // <nil>

	ic = make(chan int)
	fmt.Println(ic) //0xc000068060

	go func() {
		res := <-ic
		fmt.Println("[go1]get result:", res)
	}()

	ic <- 3
	fmt.Println("[Main]After send value to channel")
}

func main() {
	//testBufChannel()

	//testSelect()
	testTicker()
}

func testTicker() {
	var c1 = make(chan int, 1)
	var c2 = make(chan int, 1)

	/*
		r2: 2
		r1: 1
		r1: 11
		r1: 12
		tick...
		tick...
		tick...
		tick...

	*/
	go func() {
		var ticker = time.NewTicker(time.Second).C
		for {
			select {
			case r1 := <-c1:
				fmt.Println("r1:", r1)
			case r2 := <-c2:
				fmt.Println("r2:", r2)
			case <-ticker:
				fmt.Println("tick...")
			}
		}
	}()

	c2 <- 2
	c1 <- 1
	c1 <- 11
	c1 <- 12

	time.Sleep(5 * time.Second)
}

func testSelect() {
	var c1 = make(chan int, 1)
	var c2 = make(chan int, 1)

	go func() {
		for {
			select {
			case r1 := <-c1:
				fmt.Println("r1:", r1)
			case r2 := <-c2:
				fmt.Println("r2:", r2)
			default:
				fmt.Println("no data from all channel. return directly.")
				return
			}
		}
	}()

	c2 <- 2
	c1 <- 1
	c1 <- 11
	c1 <- 12

	time.Sleep(3 * time.Second)
}

func testBufChannel() {
	var c1 = make(chan int, 1)
	fmt.Println(c1)

	c1 <- 4
	res := <-c1

	c1 <- 3
	res2 := <-c1
	fmt.Println(res)
	fmt.Println(res2)
}
