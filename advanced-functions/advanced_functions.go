// Package advancedfunctions
package advancedfunctions

import "fmt"

func add(x, y int) int {
	return x + y
}

func mul(x, y int) int {
	return x * y
}

// aggregate applies the given math function to the first 3 inputs
func aggregate(a, b, c int, arithmetic func(int, int) int) int {
	return arithmetic(arithmetic(a, b), c)
}

func Run() {
	fmt.Println(aggregate(2, 3, 3, add))
	fmt.Println(aggregate(2, 3, 3, mul))
}

