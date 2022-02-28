package main

import (
	"constraints"
	"fmt"
)

func Sum[T constraints.Integer | constraints.Float](items []T) T {
	var s T
	for _, item := range items {
		s += item
	}
	return s
}

type Map interface {
}

// func SumMap[T Map[T comparable, V constraints.Integer]](m T)  V {
// 	var s V
// 	for k, v := range m {
// 		s += v
// 	}
// 	return s
// }

func main() {
	var items = []int{1, 2, 3, 4, 5}
	var items2 = []int64{1, 2, 3, 4, 5}
	var items3 = []float32{1, 2, 2.5, 4.5, 5}
	var items4 = []float64{1, 2, 3, 4, 5}
	fmt.Println(Sum(items),
		Sum(items2),
		Sum(items3),
		Sum(items4))

	// m := map[string]int64 {
	// 	"x": 1,
	// 	"y": 2
	// }
	// SumMap(m)
}
