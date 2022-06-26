package main

import "fmt"

func changeByReference(score *int) {
	*score = 99
}

func main() {
	//pointer is the memory address of variable

	var pstr *string
	fmt.Println(pstr) // <nil>

	name := "Tom"
	pstr = &name       //set the memory address of variable name to pointer
	fmt.Println(pstr)  //0xc000010250
	fmt.Println(*pstr) //Tom

	//set new value
	*pstr = "Jack"
	fmt.Println(name)  //Jack
	fmt.Println(*pstr) //Jack

	scoreOfTome := 90
	fmt.Println(scoreOfTome) //90
	changeByReference(&scoreOfTome)
	fmt.Println(scoreOfTome) //99
}
