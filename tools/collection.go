package main

import "fmt"

type Number interface {
	int | float64
}

type Collection[T Number] struct {
	data []T
}

func NewCollection[T Number](d []T) *Collection[T] {
	return &Collection[T]{
		data: d,
	}
}

func (c *Collection[T]) Filter(fn func(e T) bool) *Collection[T] {
	res := make([]T, 0)
	for _, d := range c.data {
		if fn(d) {
			res = append(res, d)
		}
	}
	return &Collection[T]{
		data: res,
	}
}

func (c *Collection[T]) ForEach(fn func(e T)) {
	for _, d := range c.data {
		fn(d)
	}
}

func main() {
	nums := make([]int, 0)
	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3)
	nums = append(nums, 4)
	nums = append(nums, 67)
	collection := NewCollection(nums)

	collection.Filter(func(a int) bool {
		return a > 3
	}).ForEach(func(a int) {
		fmt.Printf("num: %d\n", a)
	})

}
