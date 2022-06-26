package main

import "fmt"

func TestArray() {
	var nums [5]int
	fmt.Println(nums) //[0 0 0 0 0]

	var nums2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(nums2) //[1 2 3 4 5]

	var num3 = [5]int{1, 2}
	fmt.Println(num3) //[1 2 0 0 0]

	var num4 = [5]int{1: 2, 3: 4}
	fmt.Println(num4) //[0 2 0 4 0]

	nums5 := [...]int{2, 3, 4}
	fmt.Println(nums5) // [2 3 4]

	fmt.Printf("length of nums5 is: %d ", len(nums5))
}

func Test2DArray() {
	nums := [3][2]int{
		{2, 3},
		{3, 4},
		{4, 5},
	}

	fmt.Println(nums) //[[2 3] [3 4] [4 5]]
}

func TestArrayType() {
	nums1 := [3]int{1, 2, 3}
	nums2 := [4]int{2, 3, 4, 5}

	//length is also the part of array type

	fmt.Printf("nums1 type:%T\n", nums1) //nums1 type:[3]int
	fmt.Printf("nums2 type:%T\n", nums2) //nums2 type:[4]int
}

func TestArrayValueType() {
	arr1 := [3]int{1, 2, 3}
	arr2 := arr1

	arr2[1] = 5
	fmt.Println(arr1) //[1 2 3]
	fmt.Println(arr2) //[1 5 3]
}

func TestItemsInArray() {
	arr := [4]int{12, 23, 45, 66}

	for index, value := range arr {
		fmt.Printf("index:%d,value:%d\n", index, value)
	}

	fmt.Println("----------------------------")

	for _, value := range arr {
		fmt.Printf("value:%d\n", value)
	}
}

//slice

func createSlice() {
	// slice: []type, slice is reference type
	// array: [n]type, array is value type

	var slice1 []int
	fmt.Println(slice1)

	slice2 := []int{}
	fmt.Println(slice2)

	slice3 := make([]int, 3, 5)
	fmt.Println(slice3)

	////////////
	//modify slice
	arr1 := [5]int{1, 2, 3, 4, 5}
	slice4 := arr1[:]
	fmt.Println(arr1)   //[1 2 3 4 5]
	fmt.Println(slice4) //[1 2 3 4 5]

	slice4[2] = 0
	fmt.Println(arr1)   //[1 2 0 4 5]
	fmt.Println(slice4) //[1 2 0 4 5]
}

func AppendSlice() {
	//slice scale out size = 2 * current size
	names := []string{"Tom"}
	names = append(names, "Jack")
	names = append(names, "Lucy", "Lily")
	names = append(names, []string{"Wang", "Zhang"}...)

	fmt.Println(names)
}

func testMap() {
	//map[KeyType]ValueType
	userStore1 := map[string]int{}
	fmt.Println(userStore1)

	userScore := map[string]int{
		"Tom": 90,
	}
	fmt.Printf("userScore type:%T\n", userScore)
	fmt.Println(userScore)

	//add key value
	userScore["Lucy"] = 99
	userScore["Jack"] = 88
	fmt.Println(userScore)

	//delete key
	delete(userScore, "Jack")
	fmt.Println(userScore)

	//modify key
	userScore["Lucy"] = 96
	fmt.Println(userScore)

	//search
	value, exists := userScore["Tom"]
	if exists {
		fmt.Printf("Exists Tom, value:%d\n", value)
	} else {
		fmt.Printf("Not exists Tom")
	}

	// for range
	for key, value := range userScore {
		fmt.Printf("key:%s, value:%d\n", key, value)
	}

	// length
	fmt.Println(len(userScore))

	//map is reference type
	userScoreCopy := userScore
	fmt.Println(userScore)
	fmt.Println(userScoreCopy)

	userScoreCopy["Lucy"] = 98
	fmt.Println(userScore)
	fmt.Println(userScoreCopy)
}

func main() {
	//TestArray()
	//Test2DArray()
	//TestArrayType()
	//TestArrayValueType()
	//TestItemsInArray()

	//createSlice()
	//AppendSlice()

	testMap()
}
