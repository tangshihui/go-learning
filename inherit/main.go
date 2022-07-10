package main

import (
	"example/go-learning/inherit/service"
	"fmt"
)

func main() {

	users := service.GetAllUsers()

	fmt.Println(users)
}
