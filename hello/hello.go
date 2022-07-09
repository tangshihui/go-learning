package main

import (
	"fmt"
	"time"
)

const CITY = "NJ"

func Hello(name string) string {
	return "Hello," + name + "!"
}

/*
 add 2 numbers
 a: number1
 b: number2
 return : a + b
*/
func Add(a int, b int) int {
	return a + b
}

func Subtract(a int, b int) int {
	return a - b
}

func Multiply(a int, b int) int {
	return a * b
}

func Divide(a int, b int) int {
	return a / b
}

func main() {
	name := "Jack"
	fmt.Println(Hello(name))

	//variables
	//type1
	var m int = 3
	var n int = 4
	fmt.Println(Add(m, n))

	//variables
	//type2
	a, b := 2, 3

	fmt.Println(Add(a, b))
	fmt.Println(Subtract(a, b))
	fmt.Println(Multiply(a, b))
	fmt.Println(Divide(a, b))

	//const
	const BIGNUMBER = 20
	fmt.Println(Add(2, BIGNUMBER))

	/////////////
	//time format
	var yyyyMMDD_HHMMSS = "2006-01-02 15:04:05T"
	var now = time.Now()

	fmt.Println(now.Format(yyyyMMDD_HHMMSS))
	fmt.Println(time.Local)
	fmt.Println(now.Format(time.RFC3339))

	sh, _ := time.LoadLocation("Asia/Shanghai")
	tokyo, _ := time.LoadLocation("Asia/Tokyo")

	fmt.Println(now.In(sh).Format(time.RFC3339))
	fmt.Println(now.In(tokyo).Format(time.RFC3339))
	fmt.Println(now.In(time.UTC).Format(time.RFC3339))

}
