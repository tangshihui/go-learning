package main

import (
	"fmt"
	"os"
)

type IOError struct {
	path string
}

func (e *IOError) Error() string {
	return "file not found:" + e.path
}

func testOpenFile() (string, error) {
	path := "/a.txt"
	v, err := os.Open(path)
	if err != nil {
		fmt.Println("Error:", err.Error())

		return "", &IOError{path: path}
	}
	return v.Name(), nil
}

func main() {
	f, err := testOpenFile()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(f)
}
