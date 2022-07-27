package main

import (
	"fmt"
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/emirpasic/gods/utils"
)

type User struct {
	Id      int
	Name    string
	Address string
}

func main() {
	list := arraylist.New()
	list.Add("a")                         // ["a"]
	list.Add("c", "b")                    // ["a","c","b"]
	list.Sort(utils.StringComparator)     // ["a","b","c"]
	_, _ = list.Get(0)                    // "a",true
	_, _ = list.Get(100)                  // nil,false
	_ = list.Contains("a", "b", "c")      // true
	_ = list.Contains("a", "b", "c", "d") // false
	list.Swap(0, 1)                       // ["b","a",c"]
	list.Remove(2)                        // ["b","a"]
	list.Remove(1)                        // ["b"]
	list.Remove(0)                        // []
	list.Remove(0)                        // [] (ignored)
	_ = list.Empty()                      // true
	_ = list.Size()                       // 0
	list.Add("a")                         // ["a"]
	list.Clear()                          // []
	list.Add(1)                           // ["b"]
	list.Add(2)
	list.Add(3)
	list.Add(3)
	list.Add(87)
	list.Add(99)
	list.Add(99)
	list.Add(99)

	list.Select(func(index int, value interface{}) bool {
		return value.(int) > 3
	}).Each(func(index int, value interface{}) {
		fmt.Printf("%d:%v\n", index, value)
	})

	list.Clear()

	fmt.Printf("============================\n")

	list.Add(User{
		Id: 1, Name: "Tom", Address: "Beijing",
	})

	list.Add(User{
		Id: 2, Name: "Jack", Address: "Beijing",
	})

	list.Add(User{
		Id: 3, Name: "Lucy", Address: "Shanghai",
	})

	list.Each(func(index int, value interface{}) {
		fmt.Printf("%v\n", value)
	})

	list.Select(func(index int, value interface{}) bool {
		return value.(User).Id > 2
	}).Each(func(index int, value interface{}) {
		fmt.Printf("%v\n", value)
	})

	fmt.Printf("===================\n")

	set := hashset.New()   // empty
	set.Add(1)             // 1
	set.Add(2, 2, 3, 4, 5) // 3, 1, 2, 4, 5 (random order, duplicates ignored)
	set.Remove(4)          // 5, 3, 2, 1 (random order)
	set.Remove(2, 3)       // 1, 5 (random order)
	set.Contains(1)        // true
	set.Contains(1, 5)     // true
	set.Contains(1, 6)     // false
	_ = set.Values()       // []int{5,1} (random order)
	set.Clear()            // empty
	set.Empty()            // true
	set.Size()

	values := list.Map(func(index int, value interface{}) interface{} {
		return value.(User).Address
	}).Values()

	adds := hashset.New(values...)
	fmt.Printf("%v\n", adds)
	fmt.Printf("size:%d,values:%v\n", adds.Size(), adds)

	s := hashset.New(1, 2, 3, 4, 4, 5, 5)
	fmt.Printf("%v\n", s)
	fmt.Printf("size:%d,values:%v\n", s.Size(), s)

	v := hashset.New("a", "v", "b", "bb", "bb", "a", "a")
	fmt.Printf("%v\n", v)
	fmt.Printf("size:%d,values:%v\n", v.Size(), v)
}
