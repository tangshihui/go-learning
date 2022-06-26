package main

import "fmt"

type Address struct {
	city   string
	street string
}

type Student struct {
	name    string
	age     int
	class   string
	address Address
}

func (a *Address) showInfo() {
	fmt.Printf("city:%s, street:%s\n", a.city, a.street)
}

func (s *Student) showInfo() {
	fmt.Printf("name:%s,age:%d\n", s.name, s.age)
}

func (s *Student) showAddress() {
	s.address.showInfo()
}

func (s *Student) incrAge() {
	s.age = s.age + 1
}

func main() {
	tom := Student{
		name:  "Tom",
		age:   20,
		class: "CN0",
		address: Address{
			city:   "WST",
			street: "Wolf",
		},
	}

	fmt.Println(tom)

	tom.showInfo()
	tom.showAddress()

	tom.incrAge()
	tom.showInfo()
}
