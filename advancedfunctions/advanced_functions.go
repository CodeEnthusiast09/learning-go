// Package advancedfunctions
package advancedfunctions

import "fmt"

// ─────────────────────────────────────────────
// 1. FUNCTIONS AS VALUES
// ─────────────────────────────────────────────
// In Go, a function has a TYPE just like int or string.
// The type of a function is described by its signature.
// For example: func(int, int) int
// means "a function that takes two ints and returns an int"

func add(x, y int) int { return x + y }
func mul(x, y int) int { return x * y }
func sub(x, y int) int { return x - y }

// ─────────────────────────────────────────────
// 2. FUNCTIONS AS PARAMETERS (Higher-order functions)
// ─────────────────────────────────────────────
// aggregate accepts a function as its 4th argument.
// Whatever function you pass in, it uses it to combine a, b, and c.
// This is called a "higher-order function" — a function that takes another function.
func aggregate(a, b, c int, operation func(int, int) int) int {
	return operation(operation(a, b), c)
}

// ─────────────────────────────────────────────
// 3. FUNCTIONS AS RETURN VALUES
// ─────────────────────────────────────────────
// makeMultiplier returns a NEW function.
// You give it a factor (e.g. 3), and it hands back a function
// that multiplies any number by that factor.
//
// Return type here is: func(int) int
// meaning: "a function that takes an int and returns an int"
func makeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// ─────────────────────────────────────────────
// 4. CLOSURES
// ─────────────────────────────────────────────
// A closure is a function that "closes over" (captures/remembers)
// a variable from its surrounding scope.
//
// makeCounter returns a function. Each time you CALL that returned
// function, it increments its own internal count.
// The count variable lives INSIDE the closure — nobody outside can touch it.
func makeCounter() func() int {
	count := 0 // this variable is "captured" by the inner function
	return func() int {
		count++ // it remembers and modifies `count` every time it's called
		return count
	}
}

// ─────────────────────────────────────────────
// 5. VARIADIC FUNCTIONS
// ─────────────────────────────────────────────
// The `...int` means "accept zero or more ints".
// Inside the function, `nums` is just a []int slice.
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// ─────────────────────────────────────────────
// 6. ANONYMOUS FUNCTIONS (inline / immediately invoked)
// ─────────────────────────────────────────────
// An anonymous function has no name. You can:
//   a) Assign it to a variable
//   b) Pass it directly as an argument
//   c) Call it immediately (IIFE — Immediately Invoked Function Expression)

func Run() {
	fmt.Println("─── 1. Functions as values ───")

	// Storing a function in a variable — just like storing an int
	operation := add

	fmt.Println("add stored in variable:", operation(4, 5)) // 9

	fmt.Println("\n─── 2. Functions as parameters ───")
	fmt.Println("aggregate with add:", aggregate(2, 3, 3, add))  // (2+3)+3 = 8
	fmt.Println("aggregate with mul:", aggregate(2, 3, 3, mul))  // (2*3)*3 = 18
	fmt.Println("aggregate with sub:", aggregate(10, 3, 2, sub)) // (10-3)-2 = 5

	fmt.Println("\n─── 3. Functions as return values ───")
	double := makeMultiplier(2)
	triple := makeMultiplier(3)
	fmt.Println("double(7):", double(7)) // 14
	fmt.Println("triple(7):", triple(7)) // 21

	fmt.Println("\n─── 4. Closures ───")
	counter := makeCounter()
	fmt.Println("counter():", counter()) // 1
	fmt.Println("counter():", counter()) // 2
	fmt.Println("counter():", counter()) // 3
	// Each call remembers the previous count — that's the closure at work

	// Two separate counters are completely independent
	counterA := makeCounter()
	counterB := makeCounter()
	counterA()
	counterA()
	counterB()
	fmt.Println("counterA:", counterA()) // 3  (called 3 times)
	fmt.Println("counterB:", counterB()) // 2  (called 2 times)

	fmt.Println("\n─── 5. Variadic functions ───")
	fmt.Println("sum(1,2,3):", sum(1, 2, 3)) // 6
	fmt.Println("sum(10,20):", sum(10, 20))  // 30
	// You can also spread a slice into a variadic function using `...`
	numbers := []int{5, 10, 15, 20}
	fmt.Println("sum(slice...):", sum(numbers...)) // 50

	fmt.Println("\n─── 6. Anonymous functions ───")
	// a) Assigned to a variable
	square := func(n int) int {
		return n * n
	}
	fmt.Println("square(9):", square(9)) // 81

	// b) Passed directly as an argument (no need to define it separately)
	fmt.Println("aggregate with inline func:", aggregate(2, 3, 4, func(a, b int) int {
		return a + b + 1 // some custom operation defined right here
	}))

	// c) Immediately invoked — defined and called on the same line
	result := func(a, b int) int {
		return a * b
	}(6, 7) // <-- called immediately with (6, 7)
	fmt.Println("immediately invoked func result:", result) // 42
}
