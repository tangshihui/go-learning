package main

import (
	"fmt"
	"math"
	"unsafe"
)

func Integer() {
	var n1 int8 = math.MaxInt8
	var n2 int16 = math.MaxInt16
	var n3 int32 = math.MaxInt32
	var n4 int64 = math.MaxInt64

	//as for int:
	//size is 8 when OS is 64byte
	//size is 4 when OS is 32byte
	var n5 int = math.MaxInt

	fmt.Printf("n1 type:%T, size:%d, value: %d\n", n1, unsafe.Sizeof(n1), n1)
	fmt.Printf("n2 type:%T, size:%d, value: %d\n", n2, unsafe.Sizeof(n2), n2)
	fmt.Printf("n3 type:%T, size:%d, value: %d\n", n3, unsafe.Sizeof(n3), n3)
	fmt.Printf("n4 type:%T, size:%d, value: %d\n", n4, unsafe.Sizeof(n4), n4)
	fmt.Printf("n5 type:%T, size:%d, value: %d\n", n5, unsafe.Sizeof(n5), n5)
}

func Float() {
	var f1 float32 = math.MaxFloat32
	var f2 float64 = math.MaxFloat64

	fmt.Printf("f1 type:%T, size:%d, value: %f\n", f1, unsafe.Sizeof(f1), f1)
	fmt.Printf("f2 type:%T, size:%d, value: %f\n", f2, unsafe.Sizeof(f2), f2)
}

func main() {
	Integer()
	fmt.Println("------------------------------------")
	Float()
}
